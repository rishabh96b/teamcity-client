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

// NewGetAgentPoolOfAgentPoolParams creates a new GetAgentPoolOfAgentPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentPoolOfAgentPoolParams() *GetAgentPoolOfAgentPoolParams {
	return &GetAgentPoolOfAgentPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentPoolOfAgentPoolParamsWithTimeout creates a new GetAgentPoolOfAgentPoolParams object
// with the ability to set a timeout on a request.
func NewGetAgentPoolOfAgentPoolParamsWithTimeout(timeout time.Duration) *GetAgentPoolOfAgentPoolParams {
	return &GetAgentPoolOfAgentPoolParams{
		timeout: timeout,
	}
}

// NewGetAgentPoolOfAgentPoolParamsWithContext creates a new GetAgentPoolOfAgentPoolParams object
// with the ability to set a context for a request.
func NewGetAgentPoolOfAgentPoolParamsWithContext(ctx context.Context) *GetAgentPoolOfAgentPoolParams {
	return &GetAgentPoolOfAgentPoolParams{
		Context: ctx,
	}
}

// NewGetAgentPoolOfAgentPoolParamsWithHTTPClient creates a new GetAgentPoolOfAgentPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentPoolOfAgentPoolParamsWithHTTPClient(client *http.Client) *GetAgentPoolOfAgentPoolParams {
	return &GetAgentPoolOfAgentPoolParams{
		HTTPClient: client,
	}
}

/* GetAgentPoolOfAgentPoolParams contains all the parameters to send to the API endpoint
   for the get agent pool of agent pool operation.

   Typically these are written to a http.Request.
*/
type GetAgentPoolOfAgentPoolParams struct {

	// AgentPoolLocator.
	//
	// Format: AgentPoolLocator
	AgentPoolLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agent pool of agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentPoolOfAgentPoolParams) WithDefaults() *GetAgentPoolOfAgentPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent pool of agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentPoolOfAgentPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) WithTimeout(timeout time.Duration) *GetAgentPoolOfAgentPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) WithContext(ctx context.Context) *GetAgentPoolOfAgentPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) WithHTTPClient(client *http.Client) *GetAgentPoolOfAgentPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) WithAgentPoolLocator(agentPoolLocator string) *GetAgentPoolOfAgentPoolParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WithFields adds the fields to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) WithFields(fields *string) *GetAgentPoolOfAgentPoolParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get agent pool of agent pool params
func (o *GetAgentPoolOfAgentPoolParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentPoolOfAgentPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
