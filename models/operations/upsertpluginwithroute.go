// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertPluginWithRouteRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Route to lookup
	RouteID string `pathParam:"style=simple,explode=false,name=RouteId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	PluginWithoutParents components.PluginWithoutParents `request:"mediaType=application/json"`
}

func (o *UpsertPluginWithRouteRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertPluginWithRouteRequest) GetRouteID() string {
	if o == nil {
		return ""
	}
	return o.RouteID
}

func (o *UpsertPluginWithRouteRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpsertPluginWithRouteRequest) GetPluginWithoutParents() components.PluginWithoutParents {
	if o == nil {
		return components.PluginWithoutParents{}
	}
	return o.PluginWithoutParents
}

type UpsertPluginWithRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	Plugin *components.Plugin
}

func (o *UpsertPluginWithRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertPluginWithRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertPluginWithRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertPluginWithRouteResponse) GetPlugin() *components.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}
