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

// NewRemoveGroupPropertyParams creates a new RemoveGroupPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveGroupPropertyParams() *RemoveGroupPropertyParams {
	return &RemoveGroupPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveGroupPropertyParamsWithTimeout creates a new RemoveGroupPropertyParams object
// with the ability to set a timeout on a request.
func NewRemoveGroupPropertyParamsWithTimeout(timeout time.Duration) *RemoveGroupPropertyParams {
	return &RemoveGroupPropertyParams{
		timeout: timeout,
	}
}

// NewRemoveGroupPropertyParamsWithContext creates a new RemoveGroupPropertyParams object
// with the ability to set a context for a request.
func NewRemoveGroupPropertyParamsWithContext(ctx context.Context) *RemoveGroupPropertyParams {
	return &RemoveGroupPropertyParams{
		Context: ctx,
	}
}

// NewRemoveGroupPropertyParamsWithHTTPClient creates a new RemoveGroupPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveGroupPropertyParamsWithHTTPClient(client *http.Client) *RemoveGroupPropertyParams {
	return &RemoveGroupPropertyParams{
		HTTPClient: client,
	}
}

/* RemoveGroupPropertyParams contains all the parameters to send to the API endpoint
   for the remove group property operation.

   Typically these are written to a http.Request.
*/
type RemoveGroupPropertyParams struct {

	// GroupLocator.
	GroupLocator string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove group property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveGroupPropertyParams) WithDefaults() *RemoveGroupPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove group property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveGroupPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove group property params
func (o *RemoveGroupPropertyParams) WithTimeout(timeout time.Duration) *RemoveGroupPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove group property params
func (o *RemoveGroupPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove group property params
func (o *RemoveGroupPropertyParams) WithContext(ctx context.Context) *RemoveGroupPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove group property params
func (o *RemoveGroupPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove group property params
func (o *RemoveGroupPropertyParams) WithHTTPClient(client *http.Client) *RemoveGroupPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove group property params
func (o *RemoveGroupPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupLocator adds the groupLocator to the remove group property params
func (o *RemoveGroupPropertyParams) WithGroupLocator(groupLocator string) *RemoveGroupPropertyParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the remove group property params
func (o *RemoveGroupPropertyParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WithName adds the name to the remove group property params
func (o *RemoveGroupPropertyParams) WithName(name string) *RemoveGroupPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove group property params
func (o *RemoveGroupPropertyParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveGroupPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
