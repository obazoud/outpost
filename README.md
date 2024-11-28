<br>

<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="images/outpost-white.svg">
    <img alt="Outpost logo" src="images/outpost-black.svg" width="40%">
  </picture>
</div>

<br>

<div align="center">

[![License](https://img.shields.io/badge/License-Apache--2.0-blue)](#license)
[![Go Report Card](https://goreportcard.com/badge/github.com/hookdeck/outpost)](https://goreportcard.com/report/github.com/hookdeck/outpost)
[![Issues - Outpost](https://img.shields.io/github/issues/hookdeck/outposta)](https://github.com/hookdeck/outpost/issues)
![GitHub Release](https://img.shields.io/github/v/release/hookdeck/outpost)

</div>

# Open Source Event Destinations Infrastructure

Outpost is a self-hosted and open-source product that enables event producers to add Event Destinations to their platform with support for destination types such as Webhooks, Hookdeck Event Gateway, AWS SQS, Amazon EventBridge, GCP Pub/Sub, MQTT, and Azure EventBus.

Learn more about Event Destinations in the [Event Destinations Manifesto](https://event-destinations-website.vercel.app/).

## Quickstart

Clone the repo:

```sh
git clone https://github.com/hookdeck/outpost.git
```

Navigate to `outpost/examples/docker-compose/`:

```sh
cd outpost/examples/docker-compose/
```

Create a `.env` file from the example:

```sh
cp .env.example .env
```

The `TOPICS` defined in the `.env` determine:

- Which topics that destinations can subscribe to
- The topics that can be published to

Start the Outpost dependencies and services:

```sh
docker-compose -f compose.yml -f compose-rabbitmq.yml up
```

Check the services are running:

```sh
curl localhost:3333/api/v1/healthz
```

Wait until you get a `OK%` response.

Create a tenant:

```sh
curl --location --request PUT 'localhost:3333/api/v1/<TENANT_ID>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <API_KEY>'
```

Run a local server exposed via a localtunnel or use a hosted service such as the [Hookdeck Console](https://console.hookdeck.com?ref=github-outpost) to capture webhook events.

Create a webhook destination where events will be delivered to:

```sh
curl --location 'localhost:3333/api/v1/<TENANT_ID>/destinations' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <API_KEY>' \
--data '{
    "type": "webhook",
    "topics": ["*"],
    "config": {
        "url": "<URL>"
    },
    "credentials": {}
}'
```

Publish an event:

```sh
curl --location 'localhost:3333/api/v1/publish' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <API_KEY>' \
--data '{
    "tenant_id": "<TENANT_ID>",
    "topic": "user.created",
    "eligible_for_retry": true,
    "metadata": {
        "meta": "data"
    },
    "data": {
        "user_id": "userid"
    }
}'
```

Check the logs on your server or your webhook capture tool for the delivered event.

## Documentation

TODO

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This repository contains Outpost, covered under the [Apache License 2.0](LICENSE), except where noted (any Outpost logos or trademarks are not covered under the Apache License, and should be explicitly noted by a LICENSE file.)