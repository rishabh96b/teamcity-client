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

	"github.com/rishabh96b/teamcity-client/models"
)

// NewAddAgentToAgentPoolParams creates a new AddAgentToAgentPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAgentToAgentPoolParams() *AddAgentToAgentPoolParams {
	return &AddAgentToAgentPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAgentToAgentPoolParamsWithTimeout creates a new AddAgentToAgentPoolParams object
// with the ability to set a timeout on a request.
func NewAddAgentToAgentPoolParamsWithTimeout(timeout time.Duration) *AddAgentToAgentPoolParams {
	return &AddAgentToAgentPoolParams{
		timeout: timeout,
	}
}

// NewAddAgentToAgentPoolParamsWithContext creates a new AddAgentToAgentPoolParams object
// with the ability to set a context for a request.
func NewAddAgentToAgentPoolParamsWithContext(ctx context.Context) *AddAgentToAgentPoolParams {
	return &AddAgentToAgentPoolParams{
		Context: ctx,
	}
}

// NewAddAgentToAgentPoolParamsWithHTTPClient creates a new AddAgentToAgentPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAgentToAgentPoolParamsWithHTTPClient(client *http.Client) *AddAgentToAgentPoolParams {
	return &AddAgentToAgentPoolParams{
		HTTPClient: client,
	}
}

/* AddAgentToAgentPoolParams contains all the parameters to send to the API endpoint
   for the add agent to agent pool operation.

   Typically these are written to a http.Request.
*/
type AddAgentToAgentPoolParams struct {

	// AgentPoolLocator.
	//
	// Format: AgentPoolLocator
	AgentPoolLocator string

	// Body.
	Body *models.Agent

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add agent to agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAgentToAgentPoolParams) WithDefaults() *AddAgentToAgentPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add agent to agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAgentToAgentPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithTimeout(timeout time.Duration) *AddAgentToAgentPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithContext(ctx context.Context) *AddAgentToAgentPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithHTTPClient(client *http.Client) *AddAgentToAgentPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithAgentPoolLocator(agentPoolLocator string) *AddAgentToAgentPoolParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WithBody adds the body to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithBody(body *models.Agent) *AddAgentToAgentPoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetBody(body *models.Agent) {
	o.Body = body
}

// WithFields adds the fields to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) WithFields(fields *string) *AddAgentToAgentPoolParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add agent to agent pool params
func (o *AddAgentToAgentPoolParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddAgentToAgentPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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