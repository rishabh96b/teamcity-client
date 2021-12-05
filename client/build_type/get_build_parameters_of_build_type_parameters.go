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

// NewGetBuildParametersOfBuildTypeParams creates a new GetBuildParametersOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildParametersOfBuildTypeParams() *GetBuildParametersOfBuildTypeParams {
	return &GetBuildParametersOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildParametersOfBuildTypeParamsWithTimeout creates a new GetBuildParametersOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewGetBuildParametersOfBuildTypeParamsWithTimeout(timeout time.Duration) *GetBuildParametersOfBuildTypeParams {
	return &GetBuildParametersOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewGetBuildParametersOfBuildTypeParamsWithContext creates a new GetBuildParametersOfBuildTypeParams object
// with the ability to set a context for a request.
func NewGetBuildParametersOfBuildTypeParamsWithContext(ctx context.Context) *GetBuildParametersOfBuildTypeParams {
	return &GetBuildParametersOfBuildTypeParams{
		Context: ctx,
	}
}

// NewGetBuildParametersOfBuildTypeParamsWithHTTPClient creates a new GetBuildParametersOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildParametersOfBuildTypeParamsWithHTTPClient(client *http.Client) *GetBuildParametersOfBuildTypeParams {
	return &GetBuildParametersOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* GetBuildParametersOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the get build parameters of build type operation.

   Typically these are written to a http.Request.
*/
type GetBuildParametersOfBuildTypeParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	// Locator.
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build parameters of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildParametersOfBuildTypeParams) WithDefaults() *GetBuildParametersOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build parameters of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildParametersOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithTimeout(timeout time.Duration) *GetBuildParametersOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithContext(ctx context.Context) *GetBuildParametersOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithHTTPClient(client *http.Client) *GetBuildParametersOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithBtLocator(btLocator string) *GetBuildParametersOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithFields(fields *string) *GetBuildParametersOfBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) WithLocator(locator *string) *GetBuildParametersOfBuildTypeParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get build parameters of build type params
func (o *GetBuildParametersOfBuildTypeParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildParametersOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string

		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {

			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}