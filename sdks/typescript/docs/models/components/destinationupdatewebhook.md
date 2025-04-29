# DestinationUpdateWebhook

## Example Usage

```typescript
import { DestinationUpdateWebhook } from "openapi/models/components";

let value: DestinationUpdateWebhook = {
  topics: "*",
  config: {
    url: "https://example.com/webhooks/user",
  },
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `topics`                                                                                   | *components.Topics*                                                                        | :heavy_minus_sign:                                                                         | "*" or an array of enabled topics.                                                         | *                                                                                          |
| `config`                                                                                   | [components.WebhookConfig](../../models/components/webhookconfig.md)                       | :heavy_minus_sign:                                                                         | N/A                                                                                        |                                                                                            |
| `credentials`                                                                              | [components.WebhookCredentialsUpdate](../../models/components/webhookcredentialsupdate.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |                                                                                            |