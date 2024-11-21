# Outpost

TODO: What is outpost?

## Get Started

Start the Outpost dependencies and services:

```sh
make up
```

Create a tenant:

```sh
curl --location --request PUT 'localhost:4000/api/v1/<TENANT_ID>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <API_KEY>'
```

Run a local server or use a hosted service such as the [Hookdeck Console](https://console.hookdeck.com?ref=github-outpost) to capture webhook events.

Create a webhook destination where events will be delivered to:

```sh
curl --location 'localhost:4000/api/v1/<TENANT_ID>/destinations' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <API_KEY>' \
--data '{
    "type": "webhooks",
    "topics": ["*"],
    "config": {
        "url": "<URL>"
    },
    "credentials": null
}'
```

curl --location 'localhost:4000/api/v1/hookdeck/destinations' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer apikey' \
--data '{
    "type": "webhooks",
    "topics": ["*"],
    "config": {
        "url": "https://hkdk.events/p2mjk25ge583vv"
    },
    "credentials": null
}'



## Contributing

- [Getting Started](getting-started.md)
- [Test](test.md)
- [Destinations](destinations.md)
