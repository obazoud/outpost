<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"openapi"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Health.Check(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->