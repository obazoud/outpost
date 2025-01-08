# Outpost Documentation

Outpost is the first open source and self-hostable implementation of “Event Destinations” that enables delivery of your platform events directly to your user’s preferred event destinations. It supports destinations such Webhooks, Hookdeck, AWS SQS, RabbitMQ, Kafka, GCP Pub/Sub, AWS EventBridge, and Kafka.

Outpost is built and maintained by the Hookdeck. It’s written in Golang and distributed as a binary and Docker container under the Apache-2.0 licence.

Outpost has minimal dependencies (Redis, Clickhouse and one of the support message bus), is 100% backward compatible with your existing webhooks implementation and is highly optimized for high throughput low cost operation.

## Get Started

- [Outpost with RabbitMQ quickstart](1-get-started/1-rabbitmq.md)
- [Outpost with Localstack (for AWS) quickstart](1-get-started/2-localstack.md)

## Learn More

Explore the Outpost architecture, core concepts, and see the benchmark results.

[Learn more &rarr;](2-overview/README.md)

## Features

Outpost features include multi-tenant support, event topics and topics-based subscriptions, publishing events, at-least-once event delivery to a several event destination types, user alerts, and an in-built tenant user portal.

[Learn more &rarr;](3-features/README.md)

## Guides

[Explore all the guides &rarr;](4-guides/README.md)

## References

- [API Reference](5-references/1-api.md)
- [Configuration Reference](5-references/2-configuration.md)
