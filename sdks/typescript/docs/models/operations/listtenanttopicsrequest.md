# ListTenantTopicsRequest

## Example Usage

```typescript
import { ListTenantTopicsRequest } from "openapi/models/operations";

let value: ListTenantTopicsRequest = {
  tenantId: "<id>",
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenantId`                                                            | *string*                                                              | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |