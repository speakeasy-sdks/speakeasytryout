# Pet

A single Pet object


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `ID`                                       | *int64*                                    | :heavy_check_mark:                         | a unique identifier for a Pet              | 1234                                       |
| `Name`                                     | *string*                                   | :heavy_check_mark:                         | the name lovingly given to the pet         | Fido                                       |
| `Type`                                     | [*PetType](../../models/shared/pettype.md) | :heavy_minus_sign:                         | the type of pets allowed                   |                                            |