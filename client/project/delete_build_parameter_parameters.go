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

// NewDeleteBuildParameterParams creates a new DeleteBuildParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildParameterParams() *DeleteBuildParameterParams {
	return &DeleteBuildParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildParameterParamsWithTimeout creates a new DeleteBuildParameterParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildParameterParamsWithTimeout(timeout time.Duration) *DeleteBuildParameterParams {
	return &DeleteBuildParameterParams{
		timeout: timeout,
	}
}

// NewDeleteBuildParameterParamsWithContext creates a new DeleteBuildParameterParams object
// with the ability to set a context for a request.
func NewDeleteBuildParameterParamsWithContext(ctx context.Context) *DeleteBuildParameterParams {
	return &DeleteBuildParameterParams{
		Context: ctx,
	}
}

// NewDeleteBuildParameterParamsWithHTTPClient creates a new DeleteBuildParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildParameterParamsWithHTTPClient(client *http.Client) *DeleteBuildParameterParams {
	return &DeleteBuildParameterParams{
		HTTPClient: client,
	}
}

/* DeleteBuildParameterParams contains all the parameters to send to the API endpoint
   for the delete build parameter operation.

   Typically these are written to a http.Request.
*/
type DeleteBuildParameterParams struct {

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

// WithDefaults hydrates default values in the delete build parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParameterParams) WithDefaults() *DeleteBuildParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParameterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build parameter params
func (o *DeleteBuildParameterParams) WithTimeout(timeout time.Duration) *DeleteBuildParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build parameter params
func (o *DeleteBuildParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build parameter params
func (o *DeleteBuildParameterParams) WithContext(ctx context.Context) *DeleteBuildParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build parameter params
func (o *DeleteBuildParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build parameter params
func (o *DeleteBuildParameterParams) WithHTTPClient(client *http.Client) *DeleteBuildParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build parameter params
func (o *DeleteBuildParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete build parameter params
func (o *DeleteBuildParameterParams) WithName(name string) *DeleteBuildParameterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete build parameter params
func (o *DeleteBuildParameterParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the delete build parameter params
func (o *DeleteBuildParameterParams) WithProjectLocator(projectLocator string) *DeleteBuildParameterParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the delete build parameter params
func (o *DeleteBuildParameterParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
