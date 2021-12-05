// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewGetBuildResolvedParams creates a new GetBuildResolvedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildResolvedParams() *GetBuildResolvedParams {
	return &GetBuildResolvedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildResolvedParamsWithTimeout creates a new GetBuildResolvedParams object
// with the ability to set a timeout on a request.
func NewGetBuildResolvedParamsWithTimeout(timeout time.Duration) *GetBuildResolvedParams {
	return &GetBuildResolvedParams{
		timeout: timeout,
	}
}

// NewGetBuildResolvedParamsWithContext creates a new GetBuildResolvedParams object
// with the ability to set a context for a request.
func NewGetBuildResolvedParamsWithContext(ctx context.Context) *GetBuildResolvedParams {
	return &GetBuildResolvedParams{
		Context: ctx,
	}
}

// NewGetBuildResolvedParamsWithHTTPClient creates a new GetBuildResolvedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildResolvedParamsWithHTTPClient(client *http.Client) *GetBuildResolvedParams {
	return &GetBuildResolvedParams{
		HTTPClient: client,
	}
}

/* GetBuildResolvedParams contains all the parameters to send to the API endpoint
   for the get build resolved operation.

   Typically these are written to a http.Request.
*/
type GetBuildResolvedParams struct {

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	// Value.
	Value string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build resolved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildResolvedParams) WithDefaults() *GetBuildResolvedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build resolved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildResolvedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build resolved params
func (o *GetBuildResolvedParams) WithTimeout(timeout time.Duration) *GetBuildResolvedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build resolved params
func (o *GetBuildResolvedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build resolved params
func (o *GetBuildResolvedParams) WithContext(ctx context.Context) *GetBuildResolvedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build resolved params
func (o *GetBuildResolvedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build resolved params
func (o *GetBuildResolvedParams) WithHTTPClient(client *http.Client) *GetBuildResolvedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build resolved params
func (o *GetBuildResolvedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get build resolved params
func (o *GetBuildResolvedParams) WithBuildLocator(buildLocator string) *GetBuildResolvedParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get build resolved params
func (o *GetBuildResolvedParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithValue adds the value to the get build resolved params
func (o *GetBuildResolvedParams) WithValue(value string) *GetBuildResolvedParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get build resolved params
func (o *GetBuildResolvedParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildResolvedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param value
	if err := r.SetPathParam("value", o.Value); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
