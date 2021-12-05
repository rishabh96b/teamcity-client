// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAgentPoolParams creates a new DeleteAgentPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAgentPoolParams() *DeleteAgentPoolParams {
	return &DeleteAgentPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentPoolParamsWithTimeout creates a new DeleteAgentPoolParams object
// with the ability to set a timeout on a request.
func NewDeleteAgentPoolParamsWithTimeout(timeout time.Duration) *DeleteAgentPoolParams {
	return &DeleteAgentPoolParams{
		timeout: timeout,
	}
}

// NewDeleteAgentPoolParamsWithContext creates a new DeleteAgentPoolParams object
// with the ability to set a context for a request.
func NewDeleteAgentPoolParamsWithContext(ctx context.Context) *DeleteAgentPoolParams {
	return &DeleteAgentPoolParams{
		Context: ctx,
	}
}

// NewDeleteAgentPoolParamsWithHTTPClient creates a new DeleteAgentPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAgentPoolParamsWithHTTPClient(client *http.Client) *DeleteAgentPoolParams {
	return &DeleteAgentPoolParams{
		HTTPClient: client,
	}
}

/* DeleteAgentPoolParams contains all the parameters to send to the API endpoint
   for the delete agent pool operation.

   Typically these are written to a http.Request.
*/
type DeleteAgentPoolParams struct {

	// AgentPoolLocator.
	//
	// Format: AgentPoolLocator
	AgentPoolLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentPoolParams) WithDefaults() *DeleteAgentPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete agent pool params
func (o *DeleteAgentPoolParams) WithTimeout(timeout time.Duration) *DeleteAgentPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent pool params
func (o *DeleteAgentPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent pool params
func (o *DeleteAgentPoolParams) WithContext(ctx context.Context) *DeleteAgentPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent pool params
func (o *DeleteAgentPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent pool params
func (o *DeleteAgentPoolParams) WithHTTPClient(client *http.Client) *DeleteAgentPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent pool params
func (o *DeleteAgentPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the delete agent pool params
func (o *DeleteAgentPoolParams) WithAgentPoolLocator(agentPoolLocator string) *DeleteAgentPoolParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the delete agent pool params
func (o *DeleteAgentPoolParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}