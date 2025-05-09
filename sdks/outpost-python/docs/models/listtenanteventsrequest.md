# ListTenantEventsRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `tenant_id`                                                                    | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | The ID of the tenant. Required when using AdminApiKey authentication.          |
| `destination_id`                                                               | [Optional[models.DestinationID]](../models/destinationid.md)                   | :heavy_minus_sign:                                                             | Filter events by destination ID(s).                                            |
| `status`                                                                       | [Optional[models.ListTenantEventsStatus]](../models/listtenanteventsstatus.md) | :heavy_minus_sign:                                                             | Filter events by delivery status.                                              |