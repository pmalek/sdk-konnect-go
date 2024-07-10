// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type CreateRouteDestinations struct {
	IP   *string `json:"ip,omitempty"`
	Port *int64  `json:"port,omitempty"`
}

func (o *CreateRouteDestinations) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *CreateRouteDestinations) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// CreateRouteHTTPSRedirectStatusCode - The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol.
type CreateRouteHTTPSRedirectStatusCode int64

const (
	CreateRouteHTTPSRedirectStatusCodeFourHundredAndTwentySix CreateRouteHTTPSRedirectStatusCode = 426
	CreateRouteHTTPSRedirectStatusCodeThreeHundredAndOne      CreateRouteHTTPSRedirectStatusCode = 301
	CreateRouteHTTPSRedirectStatusCodeThreeHundredAndTwo      CreateRouteHTTPSRedirectStatusCode = 302
	CreateRouteHTTPSRedirectStatusCodeThreeHundredAndSeven    CreateRouteHTTPSRedirectStatusCode = 307
	CreateRouteHTTPSRedirectStatusCodeThreeHundredAndEight    CreateRouteHTTPSRedirectStatusCode = 308
)

func (e CreateRouteHTTPSRedirectStatusCode) ToPointer() *CreateRouteHTTPSRedirectStatusCode {
	return &e
}
func (e *CreateRouteHTTPSRedirectStatusCode) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 426:
		fallthrough
	case 301:
		fallthrough
	case 302:
		fallthrough
	case 307:
		fallthrough
	case 308:
		*e = CreateRouteHTTPSRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRouteHTTPSRedirectStatusCode: %v", v)
	}
}

// CreateRoutePathHandling - Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.
type CreateRoutePathHandling string

const (
	CreateRoutePathHandlingV0 CreateRoutePathHandling = "v0"
	CreateRoutePathHandlingV1 CreateRoutePathHandling = "v1"
)

func (e CreateRoutePathHandling) ToPointer() *CreateRoutePathHandling {
	return &e
}
func (e *CreateRoutePathHandling) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v0":
		fallthrough
	case "v1":
		*e = CreateRoutePathHandling(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRoutePathHandling: %v", v)
	}
}

type CreateRouteProtocols string

const (
	CreateRouteProtocolsGrpc           CreateRouteProtocols = "grpc"
	CreateRouteProtocolsGrpcs          CreateRouteProtocols = "grpcs"
	CreateRouteProtocolsHTTP           CreateRouteProtocols = "http"
	CreateRouteProtocolsHTTPS          CreateRouteProtocols = "https"
	CreateRouteProtocolsTCP            CreateRouteProtocols = "tcp"
	CreateRouteProtocolsTLS            CreateRouteProtocols = "tls"
	CreateRouteProtocolsTLSPassthrough CreateRouteProtocols = "tls_passthrough"
	CreateRouteProtocolsUDP            CreateRouteProtocols = "udp"
	CreateRouteProtocolsWs             CreateRouteProtocols = "ws"
	CreateRouteProtocolsWss            CreateRouteProtocols = "wss"
)

func (e CreateRouteProtocols) ToPointer() *CreateRouteProtocols {
	return &e
}
func (e *CreateRouteProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateRouteProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRouteProtocols: %v", v)
	}
}

type CreateRouteSources struct {
	IP   *string `json:"ip,omitempty"`
	Port *int64  `json:"port,omitempty"`
}

func (o *CreateRouteSources) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *CreateRouteSources) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// CreateRouteService - The Service this Route is associated to. This is where the Route proxies traffic to.
type CreateRouteService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRouteService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRoute - Route entities define rules to match client requests. Each Route is associated with a Service, and a Service may have multiple Routes associated to it. Every request matching a given Route will be proxied to its associated Service. The combination of Routes and Services (and the separation of concerns between them) offers a powerful routing mechanism with which it is possible to define fine-grained entry-points in Kong leading to different upstream services of your infrastructure. You need at least one matching rule that applies to the protocol being matched by the Route.
type CreateRoute struct {
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations []CreateRouteDestinations `json:"destinations,omitempty"`
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers map[string]string `json:"headers,omitempty"`
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts []string `json:"hosts,omitempty"`
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol.
	HTTPSRedirectStatusCode *CreateRouteHTTPSRedirectStatusCode `json:"https_redirect_status_code,omitempty"`
	// A list of HTTP methods that match this Route.
	Methods []string `json:"methods,omitempty"`
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name *string `json:"name,omitempty"`
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.
	PathHandling *CreateRoutePathHandling `json:"path_handling,omitempty"`
	// A list of paths that match this Route.
	Paths []string `json:"paths,omitempty"`
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost *bool `json:"preserve_host,omitempty"`
	// An array of the protocols this Route should allow. See the [Route Object](#route-object) section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols []CreateRouteProtocols `json:"protocols,omitempty"`
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regex_priority`, the older one (lowest `created_at`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority *int64 `json:"regex_priority,omitempty"`
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering *bool `json:"request_buffering,omitempty"`
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering *bool `json:"response_buffering,omitempty"`
	// A list of SNIs that match this Route when using stream routing.
	Snis []string `json:"snis,omitempty"`
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources []CreateRouteSources `json:"sources,omitempty"`
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath *bool `json:"strip_path,omitempty"`
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service *CreateRouteService `json:"service,omitempty"`
}

func (o *CreateRoute) GetDestinations() []CreateRouteDestinations {
	if o == nil {
		return nil
	}
	return o.Destinations
}

func (o *CreateRoute) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRoute) GetHosts() []string {
	if o == nil {
		return nil
	}
	return o.Hosts
}

func (o *CreateRoute) GetHTTPSRedirectStatusCode() *CreateRouteHTTPSRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.HTTPSRedirectStatusCode
}

func (o *CreateRoute) GetMethods() []string {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *CreateRoute) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateRoute) GetPathHandling() *CreateRoutePathHandling {
	if o == nil {
		return nil
	}
	return o.PathHandling
}

func (o *CreateRoute) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}

func (o *CreateRoute) GetPreserveHost() *bool {
	if o == nil {
		return nil
	}
	return o.PreserveHost
}

func (o *CreateRoute) GetProtocols() []CreateRouteProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRoute) GetRegexPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.RegexPriority
}

func (o *CreateRoute) GetRequestBuffering() *bool {
	if o == nil {
		return nil
	}
	return o.RequestBuffering
}

func (o *CreateRoute) GetResponseBuffering() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseBuffering
}

func (o *CreateRoute) GetSnis() []string {
	if o == nil {
		return nil
	}
	return o.Snis
}

func (o *CreateRoute) GetSources() []CreateRouteSources {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *CreateRoute) GetStripPath() *bool {
	if o == nil {
		return nil
	}
	return o.StripPath
}

func (o *CreateRoute) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRoute) GetService() *CreateRouteService {
	if o == nil {
		return nil
	}
	return o.Service
}
