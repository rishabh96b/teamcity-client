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

// NewDeleteBuildParametersOfBuildTypeParams creates a new DeleteBuildParametersOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildParametersOfBuildTypeParams() *DeleteBuildParametersOfBuildTypeParams {
	return &DeleteBuildParametersOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildParametersOfBuildTypeParamsWithTimeout creates a new DeleteBuildParametersOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildParametersOfBuildTypeParamsWithTimeout(timeout time.Duration) *DeleteBuildParametersOfBuildTypeParams {
	return &DeleteBuildParametersOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewDeleteBuildParametersOfBuildTypeParamsWithContext creates a new DeleteBuildParametersOfBuildTypeParams object
// with the ability to set a context for a request.
func NewDeleteBuildParametersOfBuildTypeParamsWithContext(ctx context.Context) *DeleteBuildParametersOfBuildTypeParams {
	return &DeleteBuildParametersOfBuildTypeParams{
		Context: ctx,
	}
}

// NewDeleteBuildParametersOfBuildTypeParamsWithHTTPClient creates a new DeleteBuildParametersOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildParametersOfBuildTypeParamsWithHTTPClient(client *http.Client) *DeleteBuildParametersOfBuildTypeParams {
	return &DeleteBuildParametersOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* DeleteBuildParametersOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the delete build parameters of build type operation.

   Typically these are written to a http.Request.
*/
type DeleteBuildParametersOfBuildTypeParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete build parameters of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParametersOfBuildTypeParams) WithDefaults() *DeleteBuildParametersOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build parameters of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParametersOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) WithTimeout(timeout time.Duration) *DeleteBuildParametersOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) WithContext(ctx context.Context) *DeleteBuildParametersOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) WithHTTPClient(client *http.Client) *DeleteBuildParametersOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) WithBtLocator(btLocator string) *DeleteBuildParametersOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete build parameters of build type params
func (o *DeleteBuildParametersOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildParametersOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
