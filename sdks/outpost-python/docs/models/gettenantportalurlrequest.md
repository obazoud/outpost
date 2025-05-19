# GetTenantPortalURLRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tenant_id`                                                           | *Optional[str]*                                                       | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `theme`                                                               | [Optional[models.Theme]](../models/theme.md)                          | :heavy_minus_sign:                                                    | Optional theme preference for the portal.                             |