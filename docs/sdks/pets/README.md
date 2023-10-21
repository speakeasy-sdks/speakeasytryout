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
	"context"
	"log"
	"github.com/speakeasy-sdks/speakeasytryout"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
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
| `request`                                             | [shared.Pet](../../models/shared/pet.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreatePetsResponse](../../models/operations/createpetsresponse.md), error**


## ListPets

List all pets

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/speakeasytryout"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/operations"
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListPetsRequest](../../models/operations/listpetsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.ListPetsResponse](../../models/operations/listpetsresponse.md), error**


## ShowPetByID

Info for a specific pet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/speakeasytryout"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/shared"
	"github.com/speakeasy-sdks/speakeasytryout/pkg/models/operations"
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ShowPetByIDRequest](../../models/operations/showpetbyidrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ShowPetByIDResponse](../../models/operations/showpetbyidresponse.md), error**

