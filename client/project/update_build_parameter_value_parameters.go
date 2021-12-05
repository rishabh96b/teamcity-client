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
)

// NewUpdateBuildParameterValueParams creates a new UpdateBuildParameterValueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBuildParameterValueParams() *UpdateBuildParameterValueParams {
	return &UpdateBuildParameterValueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBuildParameterValueParamsWithTimeout creates a new UpdateBuildParameterValueParams object
// with the ability to set a timeout on a request.
func NewUpdateBuildParameterValueParamsWithTimeout(timeout time.Duration) *UpdateBuildParameterValueParams {
	return &UpdateBuildParameterValueParams{
		timeout: timeout,
	}
}

// NewUpdateBuildParameterValueParamsWithContext creates a new UpdateBuildParameterValueParams object
// with the ability to set a context for a request.
func NewUpdateBuildParameterValueParamsWithContext(ctx context.Context) *UpdateBuildParameterValueParams {
	return &UpdateBuildParameterValueParams{
		Context: ctx,
	}
}

// NewUpdateBuildParameterValueParamsWithHTTPClient creates a new UpdateBuildParameterValueParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBuildParameterValueParamsWithHTTPClient(client *http.Client) *UpdateBuildParameterValueParams {
	return &UpdateBuildParameterValueParams{
		HTTPClient: client,
	}
}

/* UpdateBuildParameterValueParams contains all the parameters to send to the API endpoint
   for the update build parameter value operation.

   Typically these are written to a http.Request.
*/
type UpdateBuildParameterValueParams struct {

	// Body.
	Body string

	// Name.
	Name string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update build parameter value params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBuildParameterValueParams) WithDefaults() *UpdateBuildParameterValueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update build parameter value params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBuildParameterValueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithTimeout(timeout time.Duration) *UpdateBuildParameterValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithContext(ctx context.Context) *UpdateBuildParameterValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithHTTPClient(client *http.Client) *UpdateBuildParameterValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithBody(body string) *UpdateBuildParameterValueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetBody(body string) {
	o.Body = body
}

// WithName adds the name to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithName(name string) *UpdateBuildParameterValueParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the update build parameter value params
func (o *UpdateBuildParameterValueParams) WithProjectLocator(projectLocator string) *UpdateBuildParameterValueParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the update build parameter value params
func (o *UpdateBuildParameterValueParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBuildParameterValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
