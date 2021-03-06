// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewSetAgentPoolsProjectParams creates a new SetAgentPoolsProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAgentPoolsProjectParams() *SetAgentPoolsProjectParams {
	return &SetAgentPoolsProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAgentPoolsProjectParamsWithTimeout creates a new SetAgentPoolsProjectParams object
// with the ability to set a timeout on a request.
func NewSetAgentPoolsProjectParamsWithTimeout(timeout time.Duration) *SetAgentPoolsProjectParams {
	return &SetAgentPoolsProjectParams{
		timeout: timeout,
	}
}

// NewSetAgentPoolsProjectParamsWithContext creates a new SetAgentPoolsProjectParams object
// with the ability to set a context for a request.
func NewSetAgentPoolsProjectParamsWithContext(ctx context.Context) *SetAgentPoolsProjectParams {
	return &SetAgentPoolsProjectParams{
		Context: ctx,
	}
}

// NewSetAgentPoolsProjectParamsWithHTTPClient creates a new SetAgentPoolsProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAgentPoolsProjectParamsWithHTTPClient(client *http.Client) *SetAgentPoolsProjectParams {
	return &SetAgentPoolsProjectParams{
		HTTPClient: client,
	}
}

/* SetAgentPoolsProjectParams contains all the parameters to send to the API endpoint
   for the set agent pools project operation.

   Typically these are written to a http.Request.
*/
type SetAgentPoolsProjectParams struct {

	// Body.
	Body *models.AgentPools

	// Fields.
	Fields *string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set agent pools project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAgentPoolsProjectParams) WithDefaults() *SetAgentPoolsProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set agent pools project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAgentPoolsProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithTimeout(timeout time.Duration) *SetAgentPoolsProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithContext(ctx context.Context) *SetAgentPoolsProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithHTTPClient(client *http.Client) *SetAgentPoolsProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithBody(body *models.AgentPools) *SetAgentPoolsProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetBody(body *models.AgentPools) {
	o.Body = body
}

// WithFields adds the fields to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithFields(fields *string) *SetAgentPoolsProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the set agent pools project params
func (o *SetAgentPoolsProjectParams) WithProjectLocator(projectLocator string) *SetAgentPoolsProjectParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set agent pools project params
func (o *SetAgentPoolsProjectParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetAgentPoolsProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
