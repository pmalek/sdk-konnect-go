// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertSniRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the SNI to lookup
	SNIID string `pathParam:"style=simple,explode=false,name=SNIId"`
	// Description of the SNI
	Sni components.SNIInput `request:"mediaType=application/json"`
}

func (o *UpsertSniRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertSniRequest) GetSNIID() string {
	if o == nil {
		return ""
	}
	return o.SNIID
}

func (o *UpsertSniRequest) GetSni() components.SNIInput {
	if o == nil {
		return components.SNIInput{}
	}
	return o.Sni
}

type UpsertSniResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted SNI
	Sni *components.Sni
}

func (o *UpsertSniResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertSniResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertSniResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertSniResponse) GetSni() *components.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}
