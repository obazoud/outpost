# DestinationCreateWebhook

## Example Usage

```typescript
import { DestinationCreateWebhook } from "@hookdeck/outpost-sdk/models/components";

let value: DestinationCreateWebhook = {
  id: "user-provided-id",
  type: "webhook",
  topics: "*",
  config: {
    url: "https://example.com/webhooks/user",
  },
  credentials: {
    secret: "whsec_abc123",
    previousSecret: "whsec_xyz789",
    previousSecretInvalidAt: new Date("2024-01-02T00:00:00Z"),
  },
};
```

## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `id`                                                                                               | *string*                                                                                           | :heavy_minus_sign:                                                                                 | Optional user-provided ID. A UUID will be generated if empty.                                      | user-provided-id                                                                                   |
| `type`                                                                                             | [components.DestinationCreateWebhookType](../../models/components/destinationcreatewebhooktype.md) | :heavy_check_mark:                                                                                 | Type of the destination. Must be 'webhook'.                                                        |                                                                                                    |
| `topics`                                                                                           | *components.Topics*                                                                                | :heavy_check_mark:                                                                                 | "*" or an array of enabled topics.                                                                 | *                                                                                                  |
| `config`                                                                                           | [components.WebhookConfig](../../models/components/webhookconfig.md)                               | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `credentials`                                                                                      | [components.WebhookCredentials](../../models/components/webhookcredentials.md)                     | :heavy_minus_sign:                                                                                 | N/A                                                                                                |                                                                                                    |