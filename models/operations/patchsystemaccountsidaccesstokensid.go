// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var PatchSystemAccountsIDAccessTokensIDServerList = []string{
	"https://global.api.konghq.com/",
}

type PatchSystemAccountsIDAccessTokensIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// ID of the system account access token.
	TokenID                        string                                     `pathParam:"style=simple,explode=false,name=tokenId"`
	UpdateSystemAccountAccessToken *components.UpdateSystemAccountAccessToken `request:"mediaType=application/json"`
}

func (o *PatchSystemAccountsIDAccessTokensIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PatchSystemAccountsIDAccessTokensIDRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

func (o *PatchSystemAccountsIDAccessTokensIDRequest) GetUpdateSystemAccountAccessToken() *components.UpdateSystemAccountAccessToken {
	if o == nil {
		return nil
	}
	return o.UpdateSystemAccountAccessToken
}

type PatchSystemAccountsIDAccessTokensIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response including a single system account access token.
	SystemAccountAccessToken *components.SystemAccountAccessToken
}

func (o *PatchSystemAccountsIDAccessTokensIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchSystemAccountsIDAccessTokensIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchSystemAccountsIDAccessTokensIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchSystemAccountsIDAccessTokensIDResponse) GetSystemAccountAccessToken() *components.SystemAccountAccessToken {
	if o == nil {
		return nil
	}
	return o.SystemAccountAccessToken
}
