// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteBasicAuthWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the Basic-auth credential to lookup
	BasicAuthID string `pathParam:"style=simple,explode=false,name=BasicAuthId"`
}

func (o *DeleteBasicAuthWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *DeleteBasicAuthWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *DeleteBasicAuthWithConsumerRequest) GetBasicAuthID() string {
	if o == nil {
		return ""
	}
	return o.BasicAuthID
}

type DeleteBasicAuthWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteBasicAuthWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteBasicAuthWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteBasicAuthWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
