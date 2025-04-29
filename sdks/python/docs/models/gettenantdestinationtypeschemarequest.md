# GetTenantDestinationTypeSchemaRequest


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `tenant_id`                                                                                  | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | The ID of the tenant. Required when using AdminApiKey authentication.                        |
| `type`                                                                                       | [models.GetTenantDestinationTypeSchemaType](../models/gettenantdestinationtypeschematype.md) | :heavy_check_mark:                                                                           | The type of the destination.                                                                 |