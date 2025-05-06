# UpdateTenantDestinationRequest

## Example Usage

```typescript
import { UpdateTenantDestinationRequest } from "openapi/models/operations";

let value: UpdateTenantDestinationRequest = {
  tenantId: "<id>",
  destinationId: "<id>",
  destinationUpdate: {
    topics: "*",
    config: {
      url: "https://example.com/webhooks/user",
    },
  },
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenantId`                                                            | *string*                                                              | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `destinationId`                                                       | *string*                                                              | :heavy_check_mark:                                                    | The ID of the destination.                                            |
| `destinationUpdate`                                                   | *components.DestinationUpdate*                                        | :heavy_check_mark:                                                    | N/A                                                                   |