# GetTenantPortalUrlRequest

## Example Usage

```typescript
import { GetTenantPortalUrlRequest } from "openapi/models/operations";

let value: GetTenantPortalUrlRequest = {
  tenantId: "<id>",
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `tenantId`                                                                               | *string*                                                                                 | :heavy_minus_sign:                                                                       | The ID of the tenant. Required when using AdminApiKey authentication.                    |
| `theme`                                                                                  | [operations.GetTenantPortalUrlTheme](../../models/operations/gettenantportalurltheme.md) | :heavy_minus_sign:                                                                       | Optional theme preference for the portal.                                                |