// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewGetAgentFieldParams creates a new GetAgentFieldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentFieldParams() *GetAgentFieldParams {
	return &GetAgentFieldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentFieldParamsWithTimeout creates a new GetAgentFieldParams object
// with the ability to set a timeout on a request.
func NewGetAgentFieldParamsWithTimeout(timeout time.Duration) *GetAgentFieldParams {
	return &GetAgentFieldParams{
		timeout: timeout,
	}
}

// NewGetAgentFieldParamsWithContext creates a new GetAgentFieldParams object
// with the ability to set a context for a request.
func NewGetAgentFieldParamsWithContext(ctx context.Context) *GetAgentFieldParams {
	return &GetAgentFieldParams{
		Context: ctx,
	}
}

// NewGetAgentFieldParamsWithHTTPClient creates a new GetAgentFieldParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentFieldParamsWithHTTPClient(client *http.Client) *GetAgentFieldParams {
	return &GetAgentFieldParams{
		HTTPClient: client,
	}
}

/* GetAgentFieldParams contains all the parameters to send to the API endpoint
   for the get agent field operation.

   Typically these are written to a http.Request.
*/
type GetAgentFieldParams struct {

	// AgentLocator.
	//
	// Format: AgentLocator
	AgentLocator string

	// Field.
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agent field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentFieldParams) WithDefaults() *GetAgentFieldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentFieldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent field params
func (o *GetAgentFieldParams) WithTimeout(timeout time.Duration) *GetAgentFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent field params
func (o *GetAgentFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent field params
func (o *GetAgentFieldParams) WithContext(ctx context.Context) *GetAgentFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent field params
func (o *GetAgentFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent field params
func (o *GetAgentFieldParams) WithHTTPClient(client *http.Client) *GetAgentFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent field params
func (o *GetAgentFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the get agent field params
func (o *GetAgentFieldParams) WithAgentLocator(agentLocator string) *GetAgentFieldParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the get agent field params
func (o *GetAgentFieldParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithField adds the field to the get agent field params
func (o *GetAgentFieldParams) WithField(field string) *GetAgentFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the get agent field params
func (o *GetAgentFieldParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
