# PublishRequest

## Example Usage

```typescript
import { PublishRequest } from "@hookdeck/outpost-sdk/models/components";

let value: PublishRequest = {
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
};
```

## Fields

| Field                                            | Type                                             | Required                                         | Description                                      | Example                                          |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `tenantId`                                       | *string*                                         | :heavy_check_mark:                               | The ID of the tenant to publish for.             | <TENANT_ID>                                      |
| `destinationId`                                  | *string*                                         | :heavy_minus_sign:                               | Optional. Route event to a specific destination. | <DESTINATION_ID>                                 |
| `topic`                                          | *string*                                         | :heavy_check_mark:                               | Topic name for the event.                        | topic.name                                       |
| `eligibleForRetry`                               | *boolean*                                        | :heavy_check_mark:                               | Should event delivery be retried on failure.     |                                                  |
| `metadata`                                       | Record<string, *string*>                         | :heavy_minus_sign:                               | Any key-value string pairs for metadata.         | {<br/>"source": "crm"<br/>}                      |
| `data`                                           | Record<string, *any*>                            | :heavy_check_mark:                               | Any JSON payload for the event data.             | {<br/>"user_id": "userid",<br/>"status": "active"<br/>} |