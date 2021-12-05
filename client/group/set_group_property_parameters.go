// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewSetGroupPropertyParams creates a new SetGroupPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetGroupPropertyParams() *SetGroupPropertyParams {
	return &SetGroupPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetGroupPropertyParamsWithTimeout creates a new SetGroupPropertyParams object
// with the ability to set a timeout on a request.
func NewSetGroupPropertyParamsWithTimeout(timeout time.Duration) *SetGroupPropertyParams {
	return &SetGroupPropertyParams{
		timeout: timeout,
	}
}

// NewSetGroupPropertyParamsWithContext creates a new SetGroupPropertyParams object
// with the ability to set a context for a request.
func NewSetGroupPropertyParamsWithContext(ctx context.Context) *SetGroupPropertyParams {
	return &SetGroupPropertyParams{
		Context: ctx,
	}
}

// NewSetGroupPropertyParamsWithHTTPClient creates a new SetGroupPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetGroupPropertyParamsWithHTTPClient(client *http.Client) *SetGroupPropertyParams {
	return &SetGroupPropertyParams{
		HTTPClient: client,
	}
}

/* SetGroupPropertyParams contains all the parameters to send to the API endpoint
   for the set group property operation.

   Typically these are written to a http.Request.
*/
type SetGroupPropertyParams struct {

	// Body.
	Body string

	// GroupLocator.
	GroupLocator string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set group property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetGroupPropertyParams) WithDefaults() *SetGroupPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set group property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetGroupPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set group property params
func (o *SetGroupPropertyParams) WithTimeout(timeout time.Duration) *SetGroupPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set group property params
func (o *SetGroupPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set group property params
func (o *SetGroupPropertyParams) WithContext(ctx context.Context) *SetGroupPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set group property params
func (o *SetGroupPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set group property params
func (o *SetGroupPropertyParams) WithHTTPClient(client *http.Client) *SetGroupPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set group property params
func (o *SetGroupPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set group property params
func (o *SetGroupPropertyParams) WithBody(body string) *SetGroupPropertyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set group property params
func (o *SetGroupPropertyParams) SetBody(body string) {
	o.Body = body
}

// WithGroupLocator adds the groupLocator to the set group property params
func (o *SetGroupPropertyParams) WithGroupLocator(groupLocator string) *SetGroupPropertyParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the set group property params
func (o *SetGroupPropertyParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WithName adds the name to the set group property params
func (o *SetGroupPropertyParams) WithName(name string) *SetGroupPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set group property params
func (o *SetGroupPropertyParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SetGroupPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param groupLocator
	if err := r.SetPathParam("groupLocator", o.GroupLocator); err != nil {
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
