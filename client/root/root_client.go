// Code generated by go-swagger; DO NOT EDIT.

package root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new root API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for root API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIVersion(params *GetAPIVersionParams, opts ...ClientOption) (*GetAPIVersionOK, error)

	GetPluginInfo(params *GetPluginInfoParams, opts ...ClientOption) (*GetPluginInfoOK, error)

	GetRootEndpointsOfRoot(params *GetRootEndpointsOfRootParams, opts ...ClientOption) (*GetRootEndpointsOfRootOK, error)

	GetVersion(params *GetVersionParams, opts ...ClientOption) (*GetVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPIVersion gets the API version
*/
func (a *Client) GetAPIVersion(params *GetAPIVersionParams, opts ...ClientOption) (*GetAPIVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApiVersion",
		Method:             "GET",
		PathPattern:        "/app/rest/apiVersion",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApiVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPluginInfo gets the plugin info
*/
func (a *Client) GetPluginInfo(params *GetPluginInfoParams, opts ...ClientOption) (*GetPluginInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPluginInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPluginInfo",
		Method:             "GET",
		PathPattern:        "/app/rest/info",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPluginInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPluginInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPluginInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRootEndpointsOfRoot gets root endpoints
*/
func (a *Client) GetRootEndpointsOfRoot(params *GetRootEndpointsOfRootParams, opts ...ClientOption) (*GetRootEndpointsOfRootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRootEndpointsOfRootParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRootEndpointsOfRoot",
		Method:             "GET",
		PathPattern:        "/app/rest",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRootEndpointsOfRootReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRootEndpointsOfRootOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRootEndpointsOfRoot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVersion gets the team city server version
*/
func (a *Client) GetVersion(params *GetVersionParams, opts ...ClientOption) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVersion",
		Method:             "GET",
		PathPattern:        "/app/rest/version",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}