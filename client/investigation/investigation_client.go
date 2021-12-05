// Code generated by go-swagger; DO NOT EDIT.

package investigation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new investigation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for investigation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddInvestigation(params *AddInvestigationParams, opts ...ClientOption) (*AddInvestigationOK, error)

	AddMultipleInvestigations(params *AddMultipleInvestigationsParams, opts ...ClientOption) (*AddMultipleInvestigationsOK, error)

	DeleteInvestigation(params *DeleteInvestigationParams, opts ...ClientOption) error

	GetAllInvestigations(params *GetAllInvestigationsParams, opts ...ClientOption) (*GetAllInvestigationsOK, error)

	GetInvestigation(params *GetInvestigationParams, opts ...ClientOption) (*GetInvestigationOK, error)

	ReplaceInvestigation(params *ReplaceInvestigationParams, opts ...ClientOption) (*ReplaceInvestigationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddInvestigation creates a new investigation
*/
func (a *Client) AddInvestigation(params *AddInvestigationParams, opts ...ClientOption) (*AddInvestigationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddInvestigationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addInvestigation",
		Method:             "POST",
		PathPattern:        "/app/rest/investigations",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddInvestigationReader{formats: a.formats},
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
	success, ok := result.(*AddInvestigationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addInvestigation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddMultipleInvestigations creates multiple new investigations
*/
func (a *Client) AddMultipleInvestigations(params *AddMultipleInvestigationsParams, opts ...ClientOption) (*AddMultipleInvestigationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMultipleInvestigationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addMultipleInvestigations",
		Method:             "POST",
		PathPattern:        "/app/rest/investigations/multiple",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddMultipleInvestigationsReader{formats: a.formats},
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
	success, ok := result.(*AddMultipleInvestigationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addMultipleInvestigations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteInvestigation deletes investigation matching the locator
*/
func (a *Client) DeleteInvestigation(params *DeleteInvestigationParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvestigationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvestigation",
		Method:             "DELETE",
		PathPattern:        "/app/rest/investigations/{investigationLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInvestigationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  GetAllInvestigations gets all investigations
*/
func (a *Client) GetAllInvestigations(params *GetAllInvestigationsParams, opts ...ClientOption) (*GetAllInvestigationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllInvestigationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllInvestigations",
		Method:             "GET",
		PathPattern:        "/app/rest/investigations",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllInvestigationsReader{formats: a.formats},
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
	success, ok := result.(*GetAllInvestigationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllInvestigations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInvestigation gets investigation matching the locator
*/
func (a *Client) GetInvestigation(params *GetInvestigationParams, opts ...ClientOption) (*GetInvestigationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvestigationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvestigation",
		Method:             "GET",
		PathPattern:        "/app/rest/investigations/{investigationLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvestigationReader{formats: a.formats},
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
	success, ok := result.(*GetInvestigationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvestigation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceInvestigation updates investigation matching the locator
*/
func (a *Client) ReplaceInvestigation(params *ReplaceInvestigationParams, opts ...ClientOption) (*ReplaceInvestigationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceInvestigationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceInvestigation",
		Method:             "PUT",
		PathPattern:        "/app/rest/investigations/{investigationLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceInvestigationReader{formats: a.formats},
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
	success, ok := result.(*ReplaceInvestigationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for replaceInvestigation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
