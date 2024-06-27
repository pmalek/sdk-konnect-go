// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/Kong/sdk-konnect-go/models/components"
)

// BadRequestError - standard error
type BadRequestError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
	// invalid parameters
	InvalidParameters []components.InvalidParameters `json:"invalid_parameters"`
}

var _ error = &BadRequestError{}

func (e *BadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
