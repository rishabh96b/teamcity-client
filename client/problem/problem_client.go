// Code generated by go-swagger; DO NOT EDIT.

package problem

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new problem API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for problem API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllBuildProblems(params *GetAllBuildProblemsParams, opts ...ClientOption) (*GetAllBuildProblemsOK, error)

	GetBuildProblem(params *GetBuildProblemParams, opts ...ClientOption) (*GetBuildProblemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllBuildProblems gets all build problems
*/
func (a *Client) GetAllBuildProblems(params *GetAllBuildProblemsParams, opts ...ClientOption) (*GetAllBuildProblemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllBuildProblemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllBuildProblems",
		Method:             "GET",
		PathPattern:        "/app/rest/problems",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllBuildProblemsReader{formats: a.formats},
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
	success, ok := result.(*GetAllBuildProblemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllBuildProblems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBuildProblem gets a matching build problem
*/
func (a *Client) GetBuildProblem(params *GetBuildProblemParams, opts ...ClientOption) (*GetBuildProblemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildProblemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBuildProblem",
		Method:             "GET",
		PathPattern:        "/app/rest/problems/{problemLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildProblemReader{formats: a.formats},
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
	success, ok := result.(*GetBuildProblemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBuildProblem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
