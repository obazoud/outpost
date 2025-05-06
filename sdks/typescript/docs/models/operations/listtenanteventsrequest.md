# ListTenantEventsRequest

## Example Usage

```typescript
import { ListTenantEventsRequest } from "openapi/models/operations";

let value: ListTenantEventsRequest = {
  tenantId: "<id>",
};
```

## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `tenantId`                                                                             | *string*                                                                               | :heavy_minus_sign:                                                                     | The ID of the tenant. Required when using AdminApiKey authentication.                  |
| `destinationId`                                                                        | *operations.DestinationId*                                                             | :heavy_minus_sign:                                                                     | Filter events by destination ID(s).                                                    |
| `status`                                                                               | [operations.ListTenantEventsStatus](../../models/operations/listtenanteventsstatus.md) | :heavy_minus_sign:                                                                     | Filter events by delivery status.                                                      |