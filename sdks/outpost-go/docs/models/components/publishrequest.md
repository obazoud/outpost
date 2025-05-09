# PublishRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      | Example                                          |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `TenantID`                                       | *string*                                         | :heavy_check_mark:                               | The ID of the tenant to publish for.             | <TENANT_ID>                                      |
| `DestinationID`                                  | **string*                                        | :heavy_minus_sign:                               | Optional. Route event to a specific destination. | <DESTINATION_ID>                                 |
| `Topic`                                          | *string*                                         | :heavy_check_mark:                               | Topic name for the event.                        | topic.name                                       |
| `EligibleForRetry`                               | *bool*                                           | :heavy_check_mark:                               | Should event delivery be retried on failure.     |                                                  |
| `Metadata`                                       | map[string]*string*                              | :heavy_minus_sign:                               | Any key-value string pairs for metadata.         | {<br/>"source": "crm"<br/>}                      |
| `Data`                                           | map[string]*any*                                 | :heavy_check_mark:                               | Any JSON payload for the event data.             | {<br/>"user_id": "userid",<br/>"status": "active"<br/>} |