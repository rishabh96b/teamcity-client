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

// NewGetBuildTypeBuildsParams creates a new GetBuildTypeBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildTypeBuildsParams() *GetBuildTypeBuildsParams {
	return &GetBuildTypeBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildTypeBuildsParamsWithTimeout creates a new GetBuildTypeBuildsParams object
// with the ability to set a timeout on a request.
func NewGetBuildTypeBuildsParamsWithTimeout(timeout time.Duration) *GetBuildTypeBuildsParams {
	return &GetBuildTypeBuildsParams{
		timeout: timeout,
	}
}

// NewGetBuildTypeBuildsParamsWithContext creates a new GetBuildTypeBuildsParams object
// with the ability to set a context for a request.
func NewGetBuildTypeBuildsParamsWithContext(ctx context.Context) *GetBuildTypeBuildsParams {
	return &GetBuildTypeBuildsParams{
		Context: ctx,
	}
}

// NewGetBuildTypeBuildsParamsWithHTTPClient creates a new GetBuildTypeBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildTypeBuildsParamsWithHTTPClient(client *http.Client) *GetBuildTypeBuildsParams {
	return &GetBuildTypeBuildsParams{
		HTTPClient: client,
	}
}

/* GetBuildTypeBuildsParams contains all the parameters to send to the API endpoint
   for the get build type builds operation.

   Typically these are written to a http.Request.
*/
type GetBuildTypeBuildsParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build type builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTypeBuildsParams) WithDefaults() *GetBuildTypeBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build type builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTypeBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build type builds params
func (o *GetBuildTypeBuildsParams) WithTimeout(timeout time.Duration) *GetBuildTypeBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build type builds params
func (o *GetBuildTypeBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build type builds params
func (o *GetBuildTypeBuildsParams) WithContext(ctx context.Context) *GetBuildTypeBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build type builds params
func (o *GetBuildTypeBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build type builds params
func (o *GetBuildTypeBuildsParams) WithHTTPClient(client *http.Client) *GetBuildTypeBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build type builds params
func (o *GetBuildTypeBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get build type builds params
func (o *GetBuildTypeBuildsParams) WithBtLocator(btLocator string) *GetBuildTypeBuildsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get build type builds params
func (o *GetBuildTypeBuildsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get build type builds params
func (o *GetBuildTypeBuildsParams) WithFields(fields *string) *GetBuildTypeBuildsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build type builds params
func (o *GetBuildTypeBuildsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildTypeBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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
