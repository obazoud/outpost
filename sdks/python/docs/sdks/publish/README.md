# Publish
(*publish*)

## Overview

Operations for publishing events.

### Available Operations

* [event](#event) - Publish Event

## event

Publishes an event to the specified topic, potentially routed to a specific destination. Requires Admin API Key.

### Example Usage

```python
from openapi import SDK


with SDK(
    admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
) as sdk:

    sdk.publish.event(tenant_id="<TENANT_ID>", topic="topic.name", eligible_for_retry=False, data={
        "user_id": "userid",
        "status": "active",
    }, destination_id="<DESTINATION_ID>", metadata={
        "source": "crm",
    })

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | The ID of the tenant to publish for.                                | <TENANT_ID>                                                         |
| `topic`                                                             | *str*                                                               | :heavy_check_mark:                                                  | Topic name for the event.                                           | topic.name                                                          |
| `eligible_for_retry`                                                | *bool*                                                              | :heavy_check_mark:                                                  | Should event delivery be retried on failure.                        |                                                                     |
| `data`                                                              | Dict[str, *Any*]                                                    | :heavy_check_mark:                                                  | Any JSON payload for the event data.                                | {<br/>"user_id": "userid",<br/>"status": "active"<br/>}             |
| `destination_id`                                                    | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Optional. Route event to a specific destination.                    | <DESTINATION_ID>                                                    |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | Any key-value string pairs for metadata.                            | {<br/>"source": "crm"<br/>}                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.NotFoundError         | 404                          | application/json             |
| errors.UnauthorizedError     | 403, 407                     | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 413, 414, 415, 422, 431      | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |