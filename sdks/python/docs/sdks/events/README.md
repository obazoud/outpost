# Events
(*events*)

## Overview

Operations related to event history and deliveries.

### Available Operations

* [list](#list) - List Events
* [get](#get) - Get Event
* [list_deliveries](#list_deliveries) - List Event Delivery Attempts
* [list_by_destination](#list_by_destination) - List Events by Destination
* [get_by_destination](#get_by_destination) - Get Event by Destination
* [retry](#retry) - Retry Event Delivery

## list

Retrieves a list of events for the tenant, supporting cursor navigation (details TBD) and filtering.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.events.list(security=models.ListTenantEventsSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `security`                                                                        | [models.ListTenantEventsSecurity](../../models/listtenanteventssecurity.md)       | :heavy_check_mark:                                                                | N/A                                                                               |
| `tenant_id`                                                                       | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | The ID of the tenant. Required when using AdminApiKey authentication.             |
| `destination_id`                                                                  | [Optional[models.DestinationID]](../../models/destinationid.md)                   | :heavy_minus_sign:                                                                | Filter events by destination ID(s).                                               |
| `status`                                                                          | [Optional[models.ListTenantEventsStatus]](../../models/listtenanteventsstatus.md) | :heavy_minus_sign:                                                                | Filter events by delivery status.                                                 |
| `retries`                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                  | :heavy_minus_sign:                                                                | Configuration to override the default retry behavior of the client.               |

### Response

**[List[models.Event]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |

## get

Retrieves details for a specific event.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.events.get(security=models.GetTenantEventSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), event_id="<id>", tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `security`                                                              | [models.GetTenantEventSecurity](../../models/gettenanteventsecurity.md) | :heavy_check_mark:                                                      | N/A                                                                     |
| `event_id`                                                              | *str*                                                                   | :heavy_check_mark:                                                      | The ID of the event.                                                    |
| `tenant_id`                                                             | *Optional[str]*                                                         | :heavy_minus_sign:                                                      | The ID of the tenant. Required when using AdminApiKey authentication.   |
| `retries`                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)        | :heavy_minus_sign:                                                      | Configuration to override the default retry behavior of the client.     |

### Response

**[models.Event](../../models/event.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |

## list_deliveries

Retrieves a list of delivery attempts for a specific event, including response details.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.events.list_deliveries(security=models.ListTenantEventDeliveriesSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), event_id="<id>", tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `security`                                                                                    | [models.ListTenantEventDeliveriesSecurity](../../models/listtenanteventdeliveriessecurity.md) | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `event_id`                                                                                    | *str*                                                                                         | :heavy_check_mark:                                                                            | The ID of the event.                                                                          |
| `tenant_id`                                                                                   | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | The ID of the tenant. Required when using AdminApiKey authentication.                         |
| `retries`                                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                              | :heavy_minus_sign:                                                                            | Configuration to override the default retry behavior of the client.                           |

### Response

**[List[models.DeliveryAttempt]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |

## list_by_destination

Retrieves events associated with a specific destination for the tenant.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.events.list_by_destination(security=models.ListTenantEventsByDestinationSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), destination_id="<id>", tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `security`                                                                                                  | [models.ListTenantEventsByDestinationSecurity](../../models/listtenanteventsbydestinationsecurity.md)       | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `destination_id`                                                                                            | *str*                                                                                                       | :heavy_check_mark:                                                                                          | The ID of the destination.                                                                                  |
| `tenant_id`                                                                                                 | *Optional[str]*                                                                                             | :heavy_minus_sign:                                                                                          | The ID of the tenant. Required when using AdminApiKey authentication.                                       |
| `status`                                                                                                    | [Optional[models.ListTenantEventsByDestinationStatus]](../../models/listtenanteventsbydestinationstatus.md) | :heavy_minus_sign:                                                                                          | Filter events by delivery status.                                                                           |
| `retries`                                                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                            | :heavy_minus_sign:                                                                                          | Configuration to override the default retry behavior of the client.                                         |

### Response

**[List[models.Event]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |

## get_by_destination

Retrieves a specific event associated with a specific destination for the tenant.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.events.get_by_destination(security=models.GetTenantEventByDestinationSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), destination_id="<id>", event_id="<id>", tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `security`                                                                                        | [models.GetTenantEventByDestinationSecurity](../../models/gettenanteventbydestinationsecurity.md) | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `destination_id`                                                                                  | *str*                                                                                             | :heavy_check_mark:                                                                                | The ID of the destination.                                                                        |
| `event_id`                                                                                        | *str*                                                                                             | :heavy_check_mark:                                                                                | The ID of the event.                                                                              |
| `tenant_id`                                                                                       | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | The ID of the tenant. Required when using AdminApiKey authentication.                             |
| `retries`                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                  | :heavy_minus_sign:                                                                                | Configuration to override the default retry behavior of the client.                               |

### Response

**[models.Event](../../models/event.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |

## retry

Triggers a retry for a failed event delivery.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    sdk.events.retry(security=models.RetryTenantEventSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), destination_id="<id>", event_id="<id>", tenant_id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `security`                                                                  | [models.RetryTenantEventSecurity](../../models/retrytenanteventsecurity.md) | :heavy_check_mark:                                                          | N/A                                                                         |
| `destination_id`                                                            | *str*                                                                       | :heavy_check_mark:                                                          | The ID of the destination.                                                  |
| `event_id`                                                                  | *str*                                                                       | :heavy_check_mark:                                                          | The ID of the event to retry.                                               |
| `tenant_id`                                                                 | *Optional[str]*                                                             | :heavy_minus_sign:                                                          | The ID of the tenant. Required when using AdminApiKey authentication.       |
| `retries`                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)            | :heavy_minus_sign:                                                          | Configuration to override the default retry behavior of the client.         |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.UnauthorizedError     | 401, 403, 407                | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |