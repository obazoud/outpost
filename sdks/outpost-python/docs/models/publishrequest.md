# PublishRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      | Example                                          |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `tenant_id`                                      | *str*                                            | :heavy_check_mark:                               | The ID of the tenant to publish for.             | <TENANT_ID>                                      |
| `destination_id`                                 | *Optional[str]*                                  | :heavy_minus_sign:                               | Optional. Route event to a specific destination. | <DESTINATION_ID>                                 |
| `topic`                                          | *str*                                            | :heavy_check_mark:                               | Topic name for the event.                        | topic.name                                       |
| `eligible_for_retry`                             | *bool*                                           | :heavy_check_mark:                               | Should event delivery be retried on failure.     |                                                  |
| `metadata`                                       | Dict[str, *str*]                                 | :heavy_minus_sign:                               | Any key-value string pairs for metadata.         | {<br/>"source": "crm"<br/>}                      |
| `data`                                           | Dict[str, *Any*]                                 | :heavy_check_mark:                               | Any JSON payload for the event data.             | {<br/>"user_id": "userid",<br/>"status": "active"<br/>} |