// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetUserServerList = []string{
	"https://global.api.konghq.com/",
}

type GetUserRequest struct {
	// The ID of the user being deleted.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *GetUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type GetUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A get action response of a single user.
	User *components.User
}

func (o *GetUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
