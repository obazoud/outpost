# Outpost Overview

- [What is Outpost](#what-is-outpost)
- [Concepts](#concepts)
    - Tenants
    - Destinations
    - Events
    - Topics
- [Architecture](#architecture)
- [Runtime Requirements](#runtime-requirements)
- [Benchmarks](#benchmarks)

## What is Outpost?

TODO


## Concepts

TODO

## Architecture

Outpost is deployed with three services:

- **API Service**: captures events and exposes the necessary APIs to configure tenants and destinations.
- **Delivery Service**: is responsible for delivering events to tenants' destinations and contains adapters for each destination type. It must be configured to operate over one of the supported message queues, such as SQS and Pub/Sub.
- **Log Service**: enables storing and retrieving events, their status, and their responses.

![Outpost Architecture](images/architecture.png)

## Runtime Requirements

### API Service & Delivery Service

- Redis 6.0+ or wire-compatible alternative (RBD or AOF strongly recommended)
- One of the supported control plane message queues:
    - AWS SQS
    - GCP Pubsub
    - Azure Service Bus
    - RabbitMQ

### Log Service

- Clickhouse
    
<details>
<summary>Future Roadmap</summary>
- Postgres (simpler alternative to CH)
</details>

## Benchmarks

Coming soon...