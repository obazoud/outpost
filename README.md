# Outpost

TODO: What is outpost?

## Quickstart

Create a `.env` file from the example:

```sh
cp .env.example .env
```

The `TOPICS` defined in the `.env` determine:

- Which topics that destinations can subscribe to
- The topics that can be published to

Start the Outpost dependencies and services:

```sh
make up
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
    "type": "webhooks",
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


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).
