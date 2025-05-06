# Schemas
(*schemas*)

## Overview

Operations for retrieving destination type schemas.

### Available Operations

* [list_tenant_destination_types](#list_tenant_destination_types) - List Destination Type Schemas (for Tenant)
* [get](#get) - Get Destination Type Schema (for Tenant)
* [list_destination_types_jwt](#list_destination_types_jwt) - List Destination Type Schemas (JWT Auth)
* [get_destination_type_jwt](#get_destination_type_jwt) - Get Destination Type Schema (JWT Auth)

## list_tenant_destination_types

Returns a list of JSON-based input schemas for each available destination type. Requires Admin API Key or Tenant JWT.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.schemas.list_tenant_destination_types(security=models.ListTenantDestinationTypeSchemasSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `security`                                                                                                  | [models.ListTenantDestinationTypeSchemasSecurity](../../models/listtenantdestinationtypeschemassecurity.md) | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `tenant_id`                                                                                                 | *Optional[str]*                                                                                             | :heavy_minus_sign:                                                                                          | The ID of the tenant. Required when using AdminApiKey authentication.                                       |
| `retries`                                                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                            | :heavy_minus_sign:                                                                                          | Configuration to override the default retry behavior of the client.                                         |

### Response

**[List[models.DestinationTypeSchema]](../../models/.md)**

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

Returns the input schema for a specific destination type. Requires Admin API Key or Tenant JWT.

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.schemas.get(security=models.GetTenantDestinationTypeSchemaSecurity(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ), type_=models.GetTenantDestinationTypeSchemaType.RABBITMQ, tenant_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `security`                                                                                              | [models.GetTenantDestinationTypeSchemaSecurity](../../models/gettenantdestinationtypeschemasecurity.md) | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `type`                                                                                                  | [models.GetTenantDestinationTypeSchemaType](../../models/gettenantdestinationtypeschematype.md)         | :heavy_check_mark:                                                                                      | The type of the destination.                                                                            |
| `tenant_id`                                                                                             | *Optional[str]*                                                                                         | :heavy_minus_sign:                                                                                      | The ID of the tenant. Required when using AdminApiKey authentication.                                   |
| `retries`                                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                        | :heavy_minus_sign:                                                                                      | Configuration to override the default retry behavior of the client.                                     |

### Response

**[models.DestinationTypeSchema](../../models/destinationtypeschema.md)**

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

## list_destination_types_jwt

Returns a list of JSON-based input schemas for each available destination type (infers tenant from JWT).

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.schemas.list_destination_types_jwt(security=models.ListDestinationTypeSchemasJwtSecurity(
        tenant_jwt="<YOUR_BEARER_TOKEN_HERE>",
    ))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `security`                                                                                     | [models.ListDestinationTypeSchemasJwtSecurity](../../listdestinationtypeschemasjwtsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `retries`                                                                                      | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                               | :heavy_minus_sign:                                                                             | Configuration to override the default retry behavior of the client.                            |

### Response

**[List[models.DestinationTypeSchema]](../../models/.md)**

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

## get_destination_type_jwt

Returns the input schema for a specific destination type (infers tenant from JWT).

### Example Usage

```python
from openapi import SDK, models


with SDK() as sdk:

    res = sdk.schemas.get_destination_type_jwt(security=models.GetDestinationTypeSchemaJwtSecurity(
        tenant_jwt="<YOUR_BEARER_TOKEN_HERE>",
    ), type_=models.GetDestinationTypeSchemaJwtType.RABBITMQ)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `security`                                                                                        | [models.GetDestinationTypeSchemaJwtSecurity](../../models/getdestinationtypeschemajwtsecurity.md) | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `type`                                                                                            | [models.GetDestinationTypeSchemaJwtType](../../models/getdestinationtypeschemajwttype.md)         | :heavy_check_mark:                                                                                | The type of the destination.                                                                      |
| `retries`                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                  | :heavy_minus_sign:                                                                                | Configuration to override the default retry behavior of the client.                               |

### Response

**[models.DestinationTypeSchema](../../models/destinationtypeschema.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
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