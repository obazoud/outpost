# Outpost Documentation

Outpost is the first open source and self-hostable implementation of “Event Destinations” that enables delivery of your platform events directly to your user’s preferred event destinations. It supports destinations such Webhooks, Hookdeck, AWS SQS, RabbitMQ, Kafka, GCP Pub/Sub, AWS EventBridge, and Kafka.

Outpost is built and maintained by [Hookdeck](https://hookdeck.com?ref=github-outpost). It’s written in Golang and distributed as a binary and Docker container under the Apache-2.0 license.

Outpost has minimal dependencies (Redis, Clickhouse and one of the support message bus), is 100% backward compatible with your existing webhooks implementation and is highly optimized for high throughput low-cost operation.

## Get Started

Follow the quickstart and run Outpost with either RabbitMQ or SQS via LocalStack:

- [Try the Outpost quickstart &rarr;](1-get-started/)

## Overview

Explore the Outpost architecture, core concepts, and see the benchmark results.

[Learn more &rarr;](2-overview/README.md)

## Features

Outpost supports features required to provide a best-in-class event destinations developer experience for event producers and consumers:

- Large choice of event destination types
- Fanning out events to multiple destinations
- Sending events to specific destinations
- Destination connection instructions & authentication flows
- Event topics and topics-based subscriptions
- Automatic retries with exponential back-off
- Manual retries via Portal or API
- Multi-tenant support
- Alerts for failing destinations with debouncing
- Auto-disabling of failed destinations after too many failures
- Ability to view and list events, including request and responses
- User portal to configure destinations & inspect events
- Opt-out webhook best practices, such as headers for idempotency, timestamp and signature
- Webhook signature secret rotation
- Webhook signature format compatibility and “bring your own secrets”
- OpenTelemetry for essential performance metric observability
- Event cross-system referencing for supported destinations to display status, metadata and deep linking

[Learn more &rarr;](3-features/README.md)

## Guides

[Explore all the guides &rarr;](4-guides/README.md)

## References

- [API Reference](5-references/1-api.md)
- [Configuration Reference](5-references/2-configuration.md)
