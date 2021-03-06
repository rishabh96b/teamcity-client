// Code generated by go-swagger; DO NOT EDIT.

package mute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllMutedTests(params *GetAllMutedTestsParams, opts ...ClientOption) (*GetAllMutedTestsOK, error)

	GetMutedTest(params *GetMutedTestParams, opts ...ClientOption) (*GetMutedTestOK, error)

	MuteMultipleTests(params *MuteMultipleTestsParams, opts ...ClientOption) (*MuteMultipleTestsOK, error)

	MuteTest(params *MuteTestParams, opts ...ClientOption) (*MuteTestOK, error)

	UnmuteTest(params *UnmuteTestParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllMutedTests gets all muted tests
*/
func (a *Client) GetAllMutedTests(params *GetAllMutedTestsParams, opts ...ClientOption) (*GetAllMutedTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllMutedTestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllMutedTests",
		Method:             "GET",
		PathPattern:        "/app/rest/mutes",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllMutedTestsReader{formats: a.formats},
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
	success, ok := result.(*GetAllMutedTestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllMutedTests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMutedTest gets a muted test
*/
func (a *Client) GetMutedTest(params *GetMutedTestParams, opts ...ClientOption) (*GetMutedTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMutedTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMutedTest",
		Method:             "GET",
		PathPattern:        "/app/rest/mutes/{muteLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMutedTestReader{formats: a.formats},
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
	success, ok := result.(*GetMutedTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMutedTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MuteMultipleTests mutes multiple tests
*/
func (a *Client) MuteMultipleTests(params *MuteMultipleTestsParams, opts ...ClientOption) (*MuteMultipleTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMuteMultipleTestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "muteMultipleTests",
		Method:             "POST",
		PathPattern:        "/app/rest/mutes/multiple",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MuteMultipleTestsReader{formats: a.formats},
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
	success, ok := result.(*MuteMultipleTestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for muteMultipleTests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MuteTest mutes a test
*/
func (a *Client) MuteTest(params *MuteTestParams, opts ...ClientOption) (*MuteTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMuteTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "muteTest",
		Method:             "POST",
		PathPattern:        "/app/rest/mutes",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MuteTestReader{formats: a.formats},
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
	success, ok := result.(*MuteTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for muteTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnmuteTest unmutes the matching test
*/
func (a *Client) UnmuteTest(params *UnmuteTestParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnmuteTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unmuteTest",
		Method:             "DELETE",
		PathPattern:        "/app/rest/mutes/{muteLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnmuteTestReader{formats: a.formats},
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
