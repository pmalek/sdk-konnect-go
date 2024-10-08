// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FetchPluginSchemaRequest struct {
	// The name of the plugin
	PluginName string `pathParam:"style=simple,explode=false,name=pluginName"`
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *FetchPluginSchemaRequest) GetPluginName() string {
	if o == nil {
		return ""
	}
	return o.PluginName
}

func (o *FetchPluginSchemaRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// FetchPluginSchemaResponseBody - The schema for the plugin
type FetchPluginSchemaResponseBody struct {
	Fields []map[string]any `json:"fields,omitempty"`
}

func (o *FetchPluginSchemaResponseBody) GetFields() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Fields
}

type FetchPluginSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The schema for the plugin
	Object *FetchPluginSchemaResponseBody
}

func (o *FetchPluginSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchPluginSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchPluginSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *FetchPluginSchemaResponse) GetObject() *FetchPluginSchemaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
