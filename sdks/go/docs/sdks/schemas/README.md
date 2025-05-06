# Schemas
(*Schemas*)

## Overview

Operations for retrieving destination type schemas.

### Available Operations

* [ListTenantDestinationTypes](#listtenantdestinationtypes) - List Destination Type Schemas (for Tenant)
* [Get](#get) - Get Destination Type Schema (for Tenant)
* [ListDestinationTypesJwt](#listdestinationtypesjwt) - List Destination Type Schemas (JWT Auth)
* [GetDestinationTypeJwt](#getdestinationtypejwt) - Get Destination Type Schema (JWT Auth)

## ListTenantDestinationTypes

Returns a list of JSON-based input schemas for each available destination type. Requires Admin API Key or Tenant JWT.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New()

    res, err := s.Schemas.ListTenantDestinationTypes(ctx, operations.ListTenantDestinationTypeSchemasSecurity{
        AdminAPIKey: openapi.String("<YOUR_BEARER_TOKEN_HERE>"),
    }, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationTypeSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `security`                                                                                                                 | [operations.ListTenantDestinationTypeSchemasSecurity](../../models/operations/listtenantdestinationtypeschemassecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |
| `tenantID`                                                                                                                 | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | The ID of the tenant. Required when using AdminApiKey authentication.                                                      |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListTenantDestinationTypeSchemasResponse](../../models/operations/listtenantdestinationtypeschemasresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| apierrors.TimeoutError        | 408                           | application/json              |
| apierrors.RateLimitedError    | 429                           | application/json              |
| apierrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| apierrors.TimeoutError        | 504                           | application/json              |
| apierrors.NotFoundError       | 501, 505                      | application/json              |
| apierrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| apierrors.BadRequestError     | 510                           | application/json              |
| apierrors.UnauthorizedError   | 511                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Returns the input schema for a specific destination type. Requires Admin API Key or Tenant JWT.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New()

    res, err := s.Schemas.Get(ctx, operations.GetTenantDestinationTypeSchemaSecurity{
        AdminAPIKey: openapi.String("<YOUR_BEARER_TOKEN_HERE>"),
    }, operations.GetTenantDestinationTypeSchemaTypeRabbitmq, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationTypeSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `security`                                                                                                             | [operations.GetTenantDestinationTypeSchemaSecurity](../../models/operations/gettenantdestinationtypeschemasecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
| `type_`                                                                                                                | [operations.GetTenantDestinationTypeSchemaType](../../models/operations/gettenantdestinationtypeschematype.md)         | :heavy_check_mark:                                                                                                     | The type of the destination.                                                                                           |
| `tenantID`                                                                                                             | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | The ID of the tenant. Required when using AdminApiKey authentication.                                                  |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.GetTenantDestinationTypeSchemaResponse](../../models/operations/gettenantdestinationtypeschemaresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| apierrors.TimeoutError        | 408                           | application/json              |
| apierrors.RateLimitedError    | 429                           | application/json              |
| apierrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| apierrors.TimeoutError        | 504                           | application/json              |
| apierrors.NotFoundError       | 501, 505                      | application/json              |
| apierrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| apierrors.BadRequestError     | 510                           | application/json              |
| apierrors.UnauthorizedError   | 511                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListDestinationTypesJwt

Returns a list of JSON-based input schemas for each available destination type (infers tenant from JWT).

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New()

    res, err := s.Schemas.ListDestinationTypesJwt(ctx, operations.ListDestinationTypeSchemasJwtSecurity{
        TenantJwt: "<YOUR_BEARER_TOKEN_HERE>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationTypeSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `security`                                                                                                           | [operations.ListDestinationTypeSchemasJwtSecurity](../../models/operations/listdestinationtypeschemasjwtsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListDestinationTypeSchemasJwtResponse](../../models/operations/listdestinationtypeschemasjwtresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.UnauthorizedError   | 403, 407                      | application/json              |
| apierrors.TimeoutError        | 408                           | application/json              |
| apierrors.RateLimitedError    | 429                           | application/json              |
| apierrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| apierrors.TimeoutError        | 504                           | application/json              |
| apierrors.NotFoundError       | 501, 505                      | application/json              |
| apierrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| apierrors.BadRequestError     | 510                           | application/json              |
| apierrors.UnauthorizedError   | 511                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetDestinationTypeJwt

Returns the input schema for a specific destination type (infers tenant from JWT).

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New()

    res, err := s.Schemas.GetDestinationTypeJwt(ctx, operations.GetDestinationTypeSchemaJwtSecurity{
        TenantJwt: "<YOUR_BEARER_TOKEN_HERE>",
    }, operations.GetDestinationTypeSchemaJwtTypeRabbitmq)
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationTypeSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.GetDestinationTypeSchemaJwtSecurity](../../models/operations/getdestinationtypeschemajwtsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `type_`                                                                                                          | [operations.GetDestinationTypeSchemaJwtType](../../models/operations/getdestinationtypeschemajwttype.md)         | :heavy_check_mark:                                                                                               | The type of the destination.                                                                                     |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetDestinationTypeSchemaJwtResponse](../../models/operations/getdestinationtypeschemajwtresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 403, 407                      | application/json              |
| apierrors.TimeoutError        | 408                           | application/json              |
| apierrors.RateLimitedError    | 429                           | application/json              |
| apierrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| apierrors.TimeoutError        | 504                           | application/json              |
| apierrors.NotFoundError       | 501, 505                      | application/json              |
| apierrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| apierrors.BadRequestError     | 510                           | application/json              |
| apierrors.UnauthorizedError   | 511                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |