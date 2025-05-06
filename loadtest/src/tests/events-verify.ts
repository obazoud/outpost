import http from "k6/http";
import { check, sleep } from "k6";
import { Rate, Trend } from "k6/metrics";
import redis from "k6/experimental/redis";
import { loadEventsConfig } from "../lib/config.ts";

const ENVIRONMENT = __ENV.ENVIRONMENT || "local";
const SCENARIO = __ENV.SCENARIO || "basic";

// Only load config if we need the environment settings
const config = await loadEventsConfig(
  `../../config/environments/${ENVIRONMENT}.json`,
  `../../config/scenarios/events-verify/${SCENARIO}.json`
);

const API_KEY = __ENV.API_KEY;
if (!API_KEY) {
  throw new Error("API_KEY environment variable is required");
}

const TESTID = __ENV.TESTID;
if (!TESTID) {
  throw new Error("TESTID environment variable is required");
}

// Get max events to verify
const MAX_ITERATIONS = parseInt(__ENV.MAX_ITERATIONS || "1000", 10);

// Redis client for retrieving event IDs
// @ts-ignore - Redis client types don't match k6's implementation
const redisClient = new redis.Client(config.env.redis);

// The Redis set key containing all the events
const eventsSetKey = `events:${TESTID}`;

// Custom metrics
const verificationRate = new Rate("event_verification_rate");
const verificationTime = new Trend("event_verification_time");
const eventLatency = new Trend("event_latency");

// Default options for the test - these are our actual settings
export const options = {
  thresholds: {
    checks: ["rate>=1.0"], // 100% of checks must pass
    event_verification_rate: ["rate>=1.0"], // 100% of events must be verified
    event_latency: ["p(95)<1000"], // 95% of events should be processed in under 1 second
  },
  scenarios: {
    verify: {
      executor: "shared-iterations",
      iterations: MAX_ITERATIONS,
      vus: 10,
      maxDuration: "5m",
    },
  },
};

// Simple verification function - each VU pops one event from the set and verifies it
export default async function () {
  try {
    // Atomically pop one event ID from the set
    const eventId = await redisClient.spop(eventsSetKey);

    // If no event ID was popped, we're done
    if (!eventId) {
      return;
    }

    const verificationUrl = `${config.env.mockWebhook.url}/events/${eventId}`;

    // Parse verification poll timeout to seconds
    const timeoutStr = config.env.mockWebhook.verificationPollTimeout;
    let verificationTimeoutSec = 5; // Default 5 seconds

    if (timeoutStr.endsWith("s")) {
      verificationTimeoutSec = parseInt(timeoutStr.slice(0, -1), 10);
    } else if (timeoutStr.endsWith("ms")) {
      verificationTimeoutSec = parseInt(timeoutStr.slice(0, -2), 10) / 1000;
    }

    // Poll until timeout
    let verified = false;
    let receivedTimestamp = null;
    let successfulResponse = null; // Capture the successful verification response
    const pollInterval = 1; // 1 second between checks
    const startTime = new Date().getTime();

    for (
      let elapsed = 0;
      elapsed < verificationTimeoutSec;
      elapsed += pollInterval
    ) {
      const verificationResponse = http.get(verificationUrl);

      if (verificationResponse.status === 200) {
        verified = true;
        successfulResponse = verificationResponse; // Store the successful response
        // Get received timestamp from response
        try {
          const responseBody = verificationResponse.body;
          if (responseBody) {
            const responseData = JSON.parse(responseBody.toString());
            // Extract received_at timestamp if available
            if (responseData && responseData.received_at) {
              // Handle ISO timestamp format (e.g. "2025-04-28T13:35:17.115820003Z")
              try {
                const receivedDate = new Date(responseData.received_at);
                // Convert to milliseconds timestamp for comparison
                if (!isNaN(receivedDate.getTime())) {
                  receivedTimestamp = receivedDate.getTime();
                } else {
                  console.error(
                    `Invalid received_at timestamp format: ${responseData.received_at}`
                  );
                }
              } catch (parseError) {
                console.error(
                  `Failed to parse received_at timestamp: ${parseError}`
                );
              }
            }
          }
        } catch (e) {
          console.error(`Failed to parse verification response: ${e}`);
        }
        // Calculate and record verification time
        const verificationDuration = new Date().getTime() - startTime;
        verificationTime.add(verificationDuration);
        break;
      }

      if (elapsed + pollInterval < verificationTimeoutSec) {
        sleep(pollInterval);
      }
    }

    // Calculate event latency if verification was successful
    if (verified) {
      try {
        // Try to get sent timestamp from Redis first
        // @ts-ignore: Redis types don't handle null returns properly
        const sentTimestampStr = await redisClient.get(
          `event:${TESTID}:${eventId}:sent_at`
        );

        // Parse sent timestamp if available from Redis
        let sentTimestamp = null;
        if (sentTimestampStr !== null) {
          const parsedTimestamp = parseInt(sentTimestampStr.toString(), 10);
          if (!isNaN(parsedTimestamp) && parsedTimestamp > 0) {
            sentTimestamp = parsedTimestamp;
          }
        }

        // If no sent timestamp from Redis, try to get it from the response payload
        if (!sentTimestamp && successfulResponse && successfulResponse.body) {
          try {
            const responseBody = successfulResponse.body.toString();
            const responseData = JSON.parse(responseBody);

            // Check for sent_at in the payload
            if (
              responseData &&
              responseData.payload &&
              responseData.payload.sent_at
            ) {
              const payloadSentAt = Number(responseData.payload.sent_at);
              if (!isNaN(payloadSentAt) && payloadSentAt > 0) {
                sentTimestamp = payloadSentAt;
                console.error(
                  `Using sent_at from payload for event ${eventId}`
                );
              }
            }
          } catch (parseError) {
            console.error(
              `Failed to extract sent_at from payload: ${parseError}`
            );
          }
        }

        // Only calculate latency if we have both sent and received timestamps
        if (sentTimestamp && receivedTimestamp) {
          // Calculate latency
          const latency = receivedTimestamp - sentTimestamp;

          // Only record valid latency
          if (!isNaN(latency) && latency >= 0) {
            eventLatency.add(latency);
          } else {
            console.error(`Invalid latency calculation: ${latency}`);
          }
        } else {
          // Log diagnostic information
          if (!sentTimestamp) {
            console.error(`Missing sent timestamp for event ${eventId}`);
          }
          if (!receivedTimestamp) {
            console.error(`Missing received timestamp for event ${eventId}`);
          }
        }
      } catch (latencyErr) {
        console.error(`Failed to calculate latency: ${latencyErr}`);
      }

      // Clean up Redis key for sent timestamp
      redisClient.del([`event:${TESTID}:${eventId}:sent_at`]);
    }

    // Add to verification rate metric (true if verified, false if not)
    verificationRate.add(verified);

    // Record verification result
    check(null, {
      "event verified": () => verified,
    });
  } catch (err) {
    console.error(`Error during verification: ${err}`);
    // Throw to fail the test and prevent further iterations
    throw new Error(`Verification failed: ${err}`);
  }
}
