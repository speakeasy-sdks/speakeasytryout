# Pets
(*Pets*)

## Overview

the pet grouping

### Available Operations

* [CreatePets](#createpets) - createPets
* [ListPets](#listpets) - listPets
* [ShowPetByID](#showpetbyid) - showPetById

## CreatePets

Create a pet and key characteristics

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"github.com/speakeasy-sdks/speakeasytryout"
	"context"
	"log"
	"net/http"
)

func main() {
    s := speakeasytryout.New(
        speakeasytryout.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreatePetsResponse](../../pkg/models/operations/createpetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListPets

List all pets

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"github.com/speakeasy-sdks/speakeasytryout"
	"context"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasytryout.New(
        speakeasytryout.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pets.ListPets(ctx, operations.ListPetsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPetsResponse200JSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListPetsRequest](../../pkg/models/operations/listpetsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListPetsResponse](../../pkg/models/operations/listpetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ShowPetByID

Info for a specific pet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"github.com/speakeasy-sdks/speakeasytryout"
	"context"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasytryout.New(
        speakeasytryout.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pets.ShowPetByID(ctx, operations.ShowPetByIDRequest{
        PetID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ShowPetByIDRequest](../../pkg/models/operations/showpetbyidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ShowPetByIDResponse](../../pkg/models/operations/showpetbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
