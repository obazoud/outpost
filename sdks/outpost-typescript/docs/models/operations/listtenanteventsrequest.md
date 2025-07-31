# ListTenantEventsRequest

## Example Usage

```typescript
import { ListTenantEventsRequest } from "@hookdeck/outpost-sdk/models/operations";

let value: ListTenantEventsRequest = {};
```

## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `tenantId`                                                                             | *string*                                                                               | :heavy_minus_sign:                                                                     | The ID of the tenant. Required when using AdminApiKey authentication.                  |
| `destinationId`                                                                        | *operations.DestinationId*                                                             | :heavy_minus_sign:                                                                     | Filter events by destination ID(s).                                                    |
| `status`                                                                               | [operations.ListTenantEventsStatus](../../models/operations/listtenanteventsstatus.md) | :heavy_minus_sign:                                                                     | Filter events by delivery status.                                                      |