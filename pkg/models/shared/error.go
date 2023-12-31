// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Error struct {
	// the code associated with the error returned
	Code int64 `json:"code"`
	// detailed message about the error returned
	Message string `json:"message"`
}

func (o *Error) GetCode() int64 {
	if o == nil {
		return 0
	}
	return o.Code
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
