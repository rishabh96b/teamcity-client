// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewGetAgentRequirementParameterParams creates a new GetAgentRequirementParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentRequirementParameterParams() *GetAgentRequirementParameterParams {
	return &GetAgentRequirementParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentRequirementParameterParamsWithTimeout creates a new GetAgentRequirementParameterParams object
// with the ability to set a timeout on a request.
func NewGetAgentRequirementParameterParamsWithTimeout(timeout time.Duration) *GetAgentRequirementParameterParams {
	return &GetAgentRequirementParameterParams{
		timeout: timeout,
	}
}

// NewGetAgentRequirementParameterParamsWithContext creates a new GetAgentRequirementParameterParams object
// with the ability to set a context for a request.
func NewGetAgentRequirementParameterParamsWithContext(ctx context.Context) *GetAgentRequirementParameterParams {
	return &GetAgentRequirementParameterParams{
		Context: ctx,
	}
}

// NewGetAgentRequirementParameterParamsWithHTTPClient creates a new GetAgentRequirementParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentRequirementParameterParamsWithHTTPClient(client *http.Client) *GetAgentRequirementParameterParams {
	return &GetAgentRequirementParameterParams{
		HTTPClient: client,
	}
}

/* GetAgentRequirementParameterParams contains all the parameters to send to the API endpoint
   for the get agent requirement parameter operation.

   Typically these are written to a http.Request.
*/
type GetAgentRequirementParameterParams struct {

	// AgentRequirementLocator.
	AgentRequirementLocator string

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// FieldName.
	FieldName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agent requirement parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentRequirementParameterParams) WithDefaults() *GetAgentRequirementParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent requirement parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentRequirementParameterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithTimeout(timeout time.Duration) *GetAgentRequirementParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithContext(ctx context.Context) *GetAgentRequirementParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithHTTPClient(client *http.Client) *GetAgentRequirementParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentRequirementLocator adds the agentRequirementLocator to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithAgentRequirementLocator(agentRequirementLocator string) *GetAgentRequirementParameterParams {
	o.SetAgentRequirementLocator(agentRequirementLocator)
	return o
}

// SetAgentRequirementLocator adds the agentRequirementLocator to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetAgentRequirementLocator(agentRequirementLocator string) {
	o.AgentRequirementLocator = agentRequirementLocator
}

// WithBtLocator adds the btLocator to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithBtLocator(btLocator string) *GetAgentRequirementParameterParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFieldName adds the fieldName to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) WithFieldName(fieldName string) *GetAgentRequirementParameterParams {
	o.SetFieldName(fieldName)
	return o
}

// SetFieldName adds the fieldName to the get agent requirement parameter params
func (o *GetAgentRequirementParameterParams) SetFieldName(fieldName string) {
	o.FieldName = fieldName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentRequirementParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentRequirementLocator
	if err := r.SetPathParam("agentRequirementLocator", o.AgentRequirementLocator); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param fieldName
	if err := r.SetPathParam("fieldName", o.FieldName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}