# RetryTenantEventRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `TenantID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `DestinationID`                                                       | *string*                                                              | :heavy_check_mark:                                                    | The ID of the destination.                                            |
| `EventID`                                                             | *string*                                                              | :heavy_check_mark:                                                    | The ID of the event to retry.                                         |