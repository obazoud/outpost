# TopicsSDK
(*topics*)

## Overview

Operations for retrieving available event topics.

### Available Operations

* [list](#list) - List Available Topics (for Tenant)
* [list_jwt](#list_jwt) - List Available Topics (JWT Auth)

## list

Returns a list of available event topics configured in the Outpost instance. Requires Admin API Key or Tenant JWT.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.topics.list(security=models.ListTenantTopicsSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `security`                                                                  | [models.ListTenantTopicsSecurity](../../models/listtenanttopicssecurity.md) | :heavy_check_mark:                                                          | N/A                                                                         |
| `tenant_id`                                                                 | *Optional[str]*                                                             | :heavy_minus_sign:                                                          | The ID of the tenant. Required when using AdminApiKey authentication.       |
| `retries`                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)            | :heavy_minus_sign:                                                          | Configuration to override the default retry behavior of the client.         |

### Response

**[List[str]](../../models/.md)**

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

## list_jwt

Returns a list of available event topics configured in the Outpost instance (infers tenant from JWT).

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.topics.list_jwt(security=models.ListTopicsJwtSecurity(
        tenant_jwt="<YOUR_BEARER_TOKEN_HERE>",
    ))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `security`                                                          | [models.ListTopicsJwtSecurity](../../listtopicsjwtsecurity.md)      | :heavy_check_mark:                                                  | The security requirements to use for the request.                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[str]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.NotFoundError         | 404                          | application/json             |
| errors.UnauthorizedError     | 403, 407                     | application/json             |
| errors.TimeoutErrorT         | 408                          | application/json             |
| errors.RateLimitedError      | 429                          | application/json             |
| errors.BadRequestError       | 400, 413, 414, 415, 422, 431 | application/json             |
| errors.TimeoutErrorT         | 504                          | application/json             |
| errors.NotFoundError         | 501, 505                     | application/json             |
| errors.InternalServerError   | 500, 502, 503, 506, 507, 508 | application/json             |
| errors.BadRequestError       | 510                          | application/json             |
| errors.UnauthorizedError     | 511                          | application/json             |
| errors.APIError              | 4XX, 5XX                     | \*/\*                        |