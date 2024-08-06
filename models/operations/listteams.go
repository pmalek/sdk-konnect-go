// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListTeamsServerList = []string{
	"https://global.api.konghq.com/",
}

// ListTeamsQueryParamFilter - Filter teams returned in the response.
type ListTeamsQueryParamFilter struct {
	// Filter a string value field either by exact match or partial contains.
	Name *components.StringFieldFilter `queryParam:"name=name"`
}

func (o *ListTeamsQueryParamFilter) GetName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Name
}

type ListTeamsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter teams returned in the response.
	Filter *ListTeamsQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListTeamsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListTeamsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListTeamsRequest) GetFilter() *ListTeamsQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListTeamsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of users.
	TeamCollection *components.TeamCollection
}

func (o *ListTeamsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTeamsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTeamsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTeamsResponse) GetTeamCollection() *components.TeamCollection {
	if o == nil {
		return nil
	}
	return o.TeamCollection
}
