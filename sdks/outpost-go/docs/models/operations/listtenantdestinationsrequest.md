# ListTenantDestinationsRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `TenantID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `Type`                                                                | [*operations.Type](../../models/operations/type.md)                   | :heavy_minus_sign:                                                    | Filter destinations by type(s).                                       |
| `Topics`                                                              | [*operations.Topics](../../models/operations/topics.md)               | :heavy_minus_sign:                                                    | Filter destinations by supported topic(s).                            |