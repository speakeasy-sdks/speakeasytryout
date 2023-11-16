// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Pet - A single Pet object
type Pet struct {
	// a unique identifier for a Pet
	ID int64 `json:"id"`
	// the name lovingly given to the pet
	Name string `json:"name"`
	// the type of pets allowed
	Type *PetType `json:"type,omitempty"`
}

func (o *Pet) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Pet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Pet) GetType() *PetType {
	if o == nil {
		return nil
	}
	return o.Type
}
