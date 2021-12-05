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

// NewGetCompatibleBuildTypesParams creates a new GetCompatibleBuildTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompatibleBuildTypesParams() *GetCompatibleBuildTypesParams {
	return &GetCompatibleBuildTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompatibleBuildTypesParamsWithTimeout creates a new GetCompatibleBuildTypesParams object
// with the ability to set a timeout on a request.
func NewGetCompatibleBuildTypesParamsWithTimeout(timeout time.Duration) *GetCompatibleBuildTypesParams {
	return &GetCompatibleBuildTypesParams{
		timeout: timeout,
	}
}

// NewGetCompatibleBuildTypesParamsWithContext creates a new GetCompatibleBuildTypesParams object
// with the ability to set a context for a request.
func NewGetCompatibleBuildTypesParamsWithContext(ctx context.Context) *GetCompatibleBuildTypesParams {
	return &GetCompatibleBuildTypesParams{
		Context: ctx,
	}
}

// NewGetCompatibleBuildTypesParamsWithHTTPClient creates a new GetCompatibleBuildTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompatibleBuildTypesParamsWithHTTPClient(client *http.Client) *GetCompatibleBuildTypesParams {
	return &GetCompatibleBuildTypesParams{
		HTTPClient: client,
	}
}

/* GetCompatibleBuildTypesParams contains all the parameters to send to the API endpoint
   for the get compatible build types operation.

   Typically these are written to a http.Request.
*/
type GetCompatibleBuildTypesParams struct {

	// AgentLocator.
	//
	// Format: AgentLocator
	AgentLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compatible build types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompatibleBuildTypesParams) WithDefaults() *GetCompatibleBuildTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compatible build types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompatibleBuildTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) WithTimeout(timeout time.Duration) *GetCompatibleBuildTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) WithContext(ctx context.Context) *GetCompatibleBuildTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) WithHTTPClient(client *http.Client) *GetCompatibleBuildTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) WithAgentLocator(agentLocator string) *GetCompatibleBuildTypesParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithFields adds the fields to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) WithFields(fields *string) *GetCompatibleBuildTypesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get compatible build types params
func (o *GetCompatibleBuildTypesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompatibleBuildTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
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