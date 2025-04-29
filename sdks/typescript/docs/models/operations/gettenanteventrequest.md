# GetTenantEventRequest

## Example Usage

```typescript
import { GetTenantEventRequest } from "openapi/models/operations";

let value: GetTenantEventRequest = {
  tenantId: "<id>",
  eventId: "<id>",
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenantId`                                                            | *string*                                                              | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `eventId`                                                             | *string*                                                              | :heavy_check_mark:                                                    | The ID of the event.                                                  |