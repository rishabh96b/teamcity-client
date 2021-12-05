// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

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

// NewStopInstanceParams creates a new StopInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopInstanceParams() *StopInstanceParams {
	return &StopInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopInstanceParamsWithTimeout creates a new StopInstanceParams object
// with the ability to set a timeout on a request.
func NewStopInstanceParamsWithTimeout(timeout time.Duration) *StopInstanceParams {
	return &StopInstanceParams{
		timeout: timeout,
	}
}

// NewStopInstanceParamsWithContext creates a new StopInstanceParams object
// with the ability to set a context for a request.
func NewStopInstanceParamsWithContext(ctx context.Context) *StopInstanceParams {
	return &StopInstanceParams{
		Context: ctx,
	}
}

// NewStopInstanceParamsWithHTTPClient creates a new StopInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopInstanceParamsWithHTTPClient(client *http.Client) *StopInstanceParams {
	return &StopInstanceParams{
		HTTPClient: client,
	}
}

/* StopInstanceParams contains all the parameters to send to the API endpoint
   for the stop instance operation.

   Typically these are written to a http.Request.
*/
type StopInstanceParams struct {

	// InstanceLocator.
	InstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopInstanceParams) WithDefaults() *StopInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop instance params
func (o *StopInstanceParams) WithTimeout(timeout time.Duration) *StopInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop instance params
func (o *StopInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop instance params
func (o *StopInstanceParams) WithContext(ctx context.Context) *StopInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop instance params
func (o *StopInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop instance params
func (o *StopInstanceParams) WithHTTPClient(client *http.Client) *StopInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop instance params
func (o *StopInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceLocator adds the instanceLocator to the stop instance params
func (o *StopInstanceParams) WithInstanceLocator(instanceLocator string) *StopInstanceParams {
	o.SetInstanceLocator(instanceLocator)
	return o
}

// SetInstanceLocator adds the instanceLocator to the stop instance params
func (o *StopInstanceParams) SetInstanceLocator(instanceLocator string) {
	o.InstanceLocator = instanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *StopInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceLocator
	if err := r.SetPathParam("instanceLocator", o.InstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
