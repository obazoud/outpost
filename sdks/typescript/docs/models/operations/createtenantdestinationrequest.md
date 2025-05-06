# CreateTenantDestinationRequest

## Example Usage

```typescript
import { CreateTenantDestinationRequest } from "openapi/models/operations";

let value: CreateTenantDestinationRequest = {
  tenantId: "<id>",
  destinationCreate: {
    id: "user-provided-id",
    type: "rabbitmq",
    topics: "*",
    config: {
      serverUrl: "localhost:5672",
      exchange: "my-exchange",
      tls: "false",
    },
    credentials: {
      username: "guest",
      password: "guest",
    },
  },
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenantId`                                                            | *string*                                                              | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `destinationCreate`                                                   | *components.DestinationCreate*                                        | :heavy_check_mark:                                                    | N/A                                                                   |