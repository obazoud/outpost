# WebhookConfig

## Example Usage

```typescript
import { WebhookConfig } from "@hookdeck/outpost-sdk/models/components";

let value: WebhookConfig = {
  url: "https://example.com/webhooks/user",
};
```

## Fields

| Field                                  | Type                                   | Required                               | Description                            | Example                                |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `url`                                  | *string*                               | :heavy_check_mark:                     | The URL to send the webhook events to. | https://example.com/webhooks/user      |