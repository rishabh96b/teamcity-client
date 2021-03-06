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

// NewGetBuildParameterTypeOfBuildTypeParams creates a new GetBuildParameterTypeOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildParameterTypeOfBuildTypeParams() *GetBuildParameterTypeOfBuildTypeParams {
	return &GetBuildParameterTypeOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildParameterTypeOfBuildTypeParamsWithTimeout creates a new GetBuildParameterTypeOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewGetBuildParameterTypeOfBuildTypeParamsWithTimeout(timeout time.Duration) *GetBuildParameterTypeOfBuildTypeParams {
	return &GetBuildParameterTypeOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewGetBuildParameterTypeOfBuildTypeParamsWithContext creates a new GetBuildParameterTypeOfBuildTypeParams object
// with the ability to set a context for a request.
func NewGetBuildParameterTypeOfBuildTypeParamsWithContext(ctx context.Context) *GetBuildParameterTypeOfBuildTypeParams {
	return &GetBuildParameterTypeOfBuildTypeParams{
		Context: ctx,
	}
}

// NewGetBuildParameterTypeOfBuildTypeParamsWithHTTPClient creates a new GetBuildParameterTypeOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildParameterTypeOfBuildTypeParamsWithHTTPClient(client *http.Client) *GetBuildParameterTypeOfBuildTypeParams {
	return &GetBuildParameterTypeOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* GetBuildParameterTypeOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the get build parameter type of build type operation.

   Typically these are written to a http.Request.
*/
type GetBuildParameterTypeOfBuildTypeParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build parameter type of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildParameterTypeOfBuildTypeParams) WithDefaults() *GetBuildParameterTypeOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build parameter type of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildParameterTypeOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) WithTimeout(timeout time.Duration) *GetBuildParameterTypeOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) WithContext(ctx context.Context) *GetBuildParameterTypeOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) WithHTTPClient(client *http.Client) *GetBuildParameterTypeOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) WithBtLocator(btLocator string) *GetBuildParameterTypeOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithName adds the name to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) WithName(name string) *GetBuildParameterTypeOfBuildTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get build parameter type of build type params
func (o *GetBuildParameterTypeOfBuildTypeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildParameterTypeOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
