// Code generated by go-swagger; DO NOT EDIT.

package test_occurrence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new test occurrence API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test occurrence API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllTestOccurrences(params *GetAllTestOccurrencesParams, opts ...ClientOption) (*GetAllTestOccurrencesOK, error)

	GetTestOccurrence(params *GetTestOccurrenceParams, opts ...ClientOption) (*GetTestOccurrenceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllTestOccurrences gets all test occurrences
*/
func (a *Client) GetAllTestOccurrences(params *GetAllTestOccurrencesParams, opts ...ClientOption) (*GetAllTestOccurrencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTestOccurrencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllTestOccurrences",
		Method:             "GET",
		PathPattern:        "/app/rest/testOccurrences",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllTestOccurrencesReader{formats: a.formats},
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
	success, ok := result.(*GetAllTestOccurrencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllTestOccurrences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTestOccurrence gets a matching test occurrence
*/
func (a *Client) GetTestOccurrence(params *GetTestOccurrenceParams, opts ...ClientOption) (*GetTestOccurrenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestOccurrenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTestOccurrence",
		Method:             "GET",
		PathPattern:        "/app/rest/testOccurrences/{testLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTestOccurrenceReader{formats: a.formats},
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
	success, ok := result.(*GetTestOccurrenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTestOccurrence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
