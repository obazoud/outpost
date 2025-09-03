# CreateTenantDestinationRequest

## Example Usage

```typescript
import { CreateTenantDestinationRequest } from "@hookdeck/outpost-sdk/models/operations";

let value: CreateTenantDestinationRequest = {
  destinationCreate: {
    type: "azure_servicebus",
    topics: "*",
    config: {
      name: "my-queue-or-topic",
    },
    credentials: {
      connectionString:
        "Endpoint=sb://namespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abc123",
    },
  },
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenantId`                                                            | *string*                                                              | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `destinationCreate`                                                   | *components.DestinationCreate*                                        | :heavy_check_mark:                                                    | N/A                                                                   |