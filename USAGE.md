<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/speakeasytryout"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasytryout.New(
		speakeasytryout.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pets.CreatePets(ctx, shared.Pet{
		ID:   1234,
		Name: "Fido",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->