# ListTenantEventsByDestinationRequest

## Example Usage

```typescript
import { ListTenantEventsByDestinationRequest } from "@hookdeck/outpost-sdk/models/operations";

let value: ListTenantEventsByDestinationRequest = {
  destinationId: "<id>",
};
```

## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `tenantId`                                                                                                       | *string*                                                                                                         | :heavy_minus_sign:                                                                                               | The ID of the tenant. Required when using AdminApiKey authentication.                                            |
| `destinationId`                                                                                                  | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The ID of the destination.                                                                                       |
| `status`                                                                                                         | [operations.ListTenantEventsByDestinationStatus](../../models/operations/listtenanteventsbydestinationstatus.md) | :heavy_minus_sign:                                                                                               | Filter events by delivery status.                                                                                |