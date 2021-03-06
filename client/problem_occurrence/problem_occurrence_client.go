// Code generated by go-swagger; DO NOT EDIT.

package problem_occurrence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new problem occurrence API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for problem occurrence API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllBuildProblemOccurrences(params *GetAllBuildProblemOccurrencesParams, opts ...ClientOption) (*GetAllBuildProblemOccurrencesOK, error)

	GetBuildProblemOccurrence(params *GetBuildProblemOccurrenceParams, opts ...ClientOption) (*GetBuildProblemOccurrenceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllBuildProblemOccurrences gets all build problem occurrences
*/
func (a *Client) GetAllBuildProblemOccurrences(params *GetAllBuildProblemOccurrencesParams, opts ...ClientOption) (*GetAllBuildProblemOccurrencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllBuildProblemOccurrencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllBuildProblemOccurrences",
		Method:             "GET",
		PathPattern:        "/app/rest/problemOccurrences",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllBuildProblemOccurrencesReader{formats: a.formats},
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
	success, ok := result.(*GetAllBuildProblemOccurrencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllBuildProblemOccurrences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBuildProblemOccurrence gets a matching build problem occurrence
*/
func (a *Client) GetBuildProblemOccurrence(params *GetBuildProblemOccurrenceParams, opts ...ClientOption) (*GetBuildProblemOccurrenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildProblemOccurrenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBuildProblemOccurrence",
		Method:             "GET",
		PathPattern:        "/app/rest/problemOccurrences/{problemLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildProblemOccurrenceReader{formats: a.formats},
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
	success, ok := result.(*GetBuildProblemOccurrenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBuildProblemOccurrence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
