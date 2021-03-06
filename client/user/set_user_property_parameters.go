// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewSetUserPropertyParams creates a new SetUserPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetUserPropertyParams() *SetUserPropertyParams {
	return &SetUserPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetUserPropertyParamsWithTimeout creates a new SetUserPropertyParams object
// with the ability to set a timeout on a request.
func NewSetUserPropertyParamsWithTimeout(timeout time.Duration) *SetUserPropertyParams {
	return &SetUserPropertyParams{
		timeout: timeout,
	}
}

// NewSetUserPropertyParamsWithContext creates a new SetUserPropertyParams object
// with the ability to set a context for a request.
func NewSetUserPropertyParamsWithContext(ctx context.Context) *SetUserPropertyParams {
	return &SetUserPropertyParams{
		Context: ctx,
	}
}

// NewSetUserPropertyParamsWithHTTPClient creates a new SetUserPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetUserPropertyParamsWithHTTPClient(client *http.Client) *SetUserPropertyParams {
	return &SetUserPropertyParams{
		HTTPClient: client,
	}
}

/* SetUserPropertyParams contains all the parameters to send to the API endpoint
   for the set user property operation.

   Typically these are written to a http.Request.
*/
type SetUserPropertyParams struct {

	// Body.
	Body string

	// Name.
	Name string

	// UserLocator.
	//
	// Format: UserLocator
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set user property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserPropertyParams) WithDefaults() *SetUserPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set user property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set user property params
func (o *SetUserPropertyParams) WithTimeout(timeout time.Duration) *SetUserPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set user property params
func (o *SetUserPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set user property params
func (o *SetUserPropertyParams) WithContext(ctx context.Context) *SetUserPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set user property params
func (o *SetUserPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set user property params
func (o *SetUserPropertyParams) WithHTTPClient(client *http.Client) *SetUserPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set user property params
func (o *SetUserPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set user property params
func (o *SetUserPropertyParams) WithBody(body string) *SetUserPropertyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set user property params
func (o *SetUserPropertyParams) SetBody(body string) {
	o.Body = body
}

// WithName adds the name to the set user property params
func (o *SetUserPropertyParams) WithName(name string) *SetUserPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set user property params
func (o *SetUserPropertyParams) SetName(name string) {
	o.Name = name
}

// WithUserLocator adds the userLocator to the set user property params
func (o *SetUserPropertyParams) WithUserLocator(userLocator string) *SetUserPropertyParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the set user property params
func (o *SetUserPropertyParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetUserPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
