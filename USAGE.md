<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/speakeasytryout"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
)

func main() {
    s := speakeasytryout.New(
        speakeasytryout.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pets.CreatePets(ctx, shared.Pet{
        ID: 1234,
        Name: "Fido",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->