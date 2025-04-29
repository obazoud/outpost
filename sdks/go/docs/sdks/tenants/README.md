# Tenants
(*Tenants*)

## Overview

The API segments resources per `tenant`. A tenant represents a user/team/organization in your product. The provided value determines the tenant's ID, which can be any string representation.

If your system is not multi-tenant, create a single tenant with a hard-code tenant ID upon initialization. If your system has a single tenant but multiple environments, create a tenant per environment, like `live` and `test`.


### Available Operations

* [Upsert](#upsert) - Create or Update Tenant
* [Get](#get) - Get Tenant
* [Delete](#delete) - Delete Tenant
* [GetPortalURL](#getportalurl) - Get Portal Redirect URL
* [GetToken](#gettoken) - Get Tenant JWT Token
* [GetPortalURLJwtContext](#getportalurljwtcontext) - Get Portal Redirect URL (JWT Auth Context)
* [GetTokenJwtContext](#gettokenjwtcontext) - Get Tenant JWT Token (JWT Auth Context)

## Upsert

Idempotently creates or updates a tenant. Required before associating destinations.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tenants.Upsert(ctx, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Tenant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `tenantID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.UpsertTenantResponse](../../models/operations/upserttenantresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotFoundError       | 404                           | application/json              |
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

Retrieves details for a specific tenant.

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

    res, err := s.Tenants.Get(ctx, operations.GetTenantSecurity{
        AdminAPIKey: openapi.String("<YOUR_BEARER_TOKEN_HERE>"),
    }, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Tenant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `security`                                                                   | [operations.GetTenantSecurity](../../models/operations/gettenantsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `tenantID`                                                                   | **string*                                                                    | :heavy_minus_sign:                                                           | The ID of the tenant. Required when using AdminApiKey authentication.        |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetTenantResponse](../../models/operations/gettenantresponse.md), error**

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

## Delete

Deletes the tenant and all associated destinations.

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

    res, err := s.Tenants.Delete(ctx, operations.DeleteTenantSecurity{
        AdminAPIKey: openapi.String("<YOUR_BEARER_TOKEN_HERE>"),
    }, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.DeleteTenantSecurity](../../models/operations/deletetenantsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `tenantID`                                                                         | **string*                                                                          | :heavy_minus_sign:                                                                 | The ID of the tenant. Required when using AdminApiKey authentication.              |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.DeleteTenantResponse](../../models/operations/deletetenantresponse.md), error**

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

## GetPortalURL

Returns a redirect URL containing a JWT to authenticate the user with the portal.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tenants.GetPortalURL(ctx, openapi.String("<id>"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalRedirect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `tenantID`                                                                                | **string*                                                                                 | :heavy_minus_sign:                                                                        | The ID of the tenant. Required when using AdminApiKey authentication.                     |
| `theme`                                                                                   | [*operations.GetTenantPortalURLTheme](../../models/operations/gettenantportalurltheme.md) | :heavy_minus_sign:                                                                        | Optional theme preference for the portal.                                                 |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.GetTenantPortalURLResponse](../../models/operations/gettenantportalurlresponse.md), error**

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

## GetToken

Returns a JWT token scoped to the tenant for safe browser API calls.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tenants.GetToken(ctx, openapi.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `tenantID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | The ID of the tenant. Required when using AdminApiKey authentication. |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.GetTenantTokenResponse](../../models/operations/gettenanttokenresponse.md), error**

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

## GetPortalURLJwtContext

Returns a redirect URL containing a JWT to authenticate the user with the portal (infers tenant from JWT). Requires Admin API Key.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tenants.GetPortalURLJwtContext(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalRedirect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `theme`                                                                                           | [*operations.GetPortalURLJwtContextTheme](../../models/operations/getportalurljwtcontexttheme.md) | :heavy_minus_sign:                                                                                | Optional theme preference for the portal.                                                         |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.GetPortalURLJwtContextResponse](../../models/operations/getportalurljwtcontextresponse.md), error**

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

## GetTokenJwtContext

Returns a JWT token scoped to the tenant (infers tenant from JWT). Requires Admin API Key.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
    ctx := context.Background()

    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tenants.GetTokenJwtContext(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTenantTokenJwtContextResponse](../../models/operations/gettenanttokenjwtcontextresponse.md), error**

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