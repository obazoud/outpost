# Publish
(*publish*)

## Overview

Operations for publishing events.

### Available Operations

* [event](#event) - Publish Event

## event

Publishes an event to the specified topic, potentially routed to a specific destination. Requires Admin API Key.

### Example Usage

```typescript
import { SDK } from "openapi";

const sdk = new SDK({
  adminApiKey: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  await sdk.publish.event({
    tenantId: "<TENANT_ID>",
    destinationId: "<DESTINATION_ID>",
    topic: "topic.name",
    eligibleForRetry: false,
    metadata: {
      "source": "crm",
    },
    data: {
      "user_id": "userid",
      "status": "active",
    },
  });


}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SDKCore } from "openapi/core.js";
import { publishEvent } from "openapi/funcs/publishEvent.js";

// Use `SDKCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const sdk = new SDKCore({
  adminApiKey: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const res = await publishEvent(sdk, {
    tenantId: "<TENANT_ID>",
    destinationId: "<DESTINATION_ID>",
    topic: "topic.name",
    eligibleForRetry: false,
    metadata: {
      "source": "crm",
    },
    data: {
      "user_id": "userid",
      "status": "active",
    },
  });

  if (!res.ok) {
    throw res.error;
  }

  const { value: result } = res;

  
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.PublishRequest](../../models/components/publishrequest.md)                                                                                                         | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<void\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.NotFoundError         | 404                          | application/json             |
| errors.UnauthorizedError     | 403, 407                     | application/json             |
| errors.TimeoutError          | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 413, 414, 415, 422, 431      | application/json             |
| errors.TimeoutError          | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |