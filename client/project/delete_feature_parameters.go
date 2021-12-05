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

// NewDeleteFeatureParams creates a new DeleteFeatureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFeatureParams() *DeleteFeatureParams {
	return &DeleteFeatureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeatureParamsWithTimeout creates a new DeleteFeatureParams object
// with the ability to set a timeout on a request.
func NewDeleteFeatureParamsWithTimeout(timeout time.Duration) *DeleteFeatureParams {
	return &DeleteFeatureParams{
		timeout: timeout,
	}
}

// NewDeleteFeatureParamsWithContext creates a new DeleteFeatureParams object
// with the ability to set a context for a request.
func NewDeleteFeatureParamsWithContext(ctx context.Context) *DeleteFeatureParams {
	return &DeleteFeatureParams{
		Context: ctx,
	}
}

// NewDeleteFeatureParamsWithHTTPClient creates a new DeleteFeatureParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFeatureParamsWithHTTPClient(client *http.Client) *DeleteFeatureParams {
	return &DeleteFeatureParams{
		HTTPClient: client,
	}
}

/* DeleteFeatureParams contains all the parameters to send to the API endpoint
   for the delete feature operation.

   Typically these are written to a http.Request.
*/
type DeleteFeatureParams struct {

	// FeatureLocator.
	FeatureLocator string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete feature params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFeatureParams) WithDefaults() *DeleteFeatureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete feature params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFeatureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete feature params
func (o *DeleteFeatureParams) WithTimeout(timeout time.Duration) *DeleteFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feature params
func (o *DeleteFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feature params
func (o *DeleteFeatureParams) WithContext(ctx context.Context) *DeleteFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feature params
func (o *DeleteFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feature params
func (o *DeleteFeatureParams) WithHTTPClient(client *http.Client) *DeleteFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feature params
func (o *DeleteFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureLocator adds the featureLocator to the delete feature params
func (o *DeleteFeatureParams) WithFeatureLocator(featureLocator string) *DeleteFeatureParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the delete feature params
func (o *DeleteFeatureParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithProjectLocator adds the projectLocator to the delete feature params
func (o *DeleteFeatureParams) WithProjectLocator(projectLocator string) *DeleteFeatureParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the delete feature params
func (o *DeleteFeatureParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureLocator
	if err := r.SetPathParam("featureLocator", o.FeatureLocator); err != nil {
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