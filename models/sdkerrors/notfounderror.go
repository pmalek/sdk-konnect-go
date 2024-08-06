// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// NotFoundError - standard error
type NotFoundError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
}

var _ error = &NotFoundError{}

func (e *NotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
