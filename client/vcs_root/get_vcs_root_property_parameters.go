// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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

// NewGetVcsRootPropertyParams creates a new GetVcsRootPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootPropertyParams() *GetVcsRootPropertyParams {
	return &GetVcsRootPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootPropertyParamsWithTimeout creates a new GetVcsRootPropertyParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootPropertyParamsWithTimeout(timeout time.Duration) *GetVcsRootPropertyParams {
	return &GetVcsRootPropertyParams{
		timeout: timeout,
	}
}

// NewGetVcsRootPropertyParamsWithContext creates a new GetVcsRootPropertyParams object
// with the ability to set a context for a request.
func NewGetVcsRootPropertyParamsWithContext(ctx context.Context) *GetVcsRootPropertyParams {
	return &GetVcsRootPropertyParams{
		Context: ctx,
	}
}

// NewGetVcsRootPropertyParamsWithHTTPClient creates a new GetVcsRootPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootPropertyParamsWithHTTPClient(client *http.Client) *GetVcsRootPropertyParams {
	return &GetVcsRootPropertyParams{
		HTTPClient: client,
	}
}

/* GetVcsRootPropertyParams contains all the parameters to send to the API endpoint
   for the get vcs root property operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootPropertyParams struct {

	// Name.
	Name string

	// VcsRootLocator.
	//
	// Format: VcsRootLocator
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs root property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootPropertyParams) WithDefaults() *GetVcsRootPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root property params
func (o *GetVcsRootPropertyParams) WithTimeout(timeout time.Duration) *GetVcsRootPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root property params
func (o *GetVcsRootPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root property params
func (o *GetVcsRootPropertyParams) WithContext(ctx context.Context) *GetVcsRootPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root property params
func (o *GetVcsRootPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root property params
func (o *GetVcsRootPropertyParams) WithHTTPClient(client *http.Client) *GetVcsRootPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root property params
func (o *GetVcsRootPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get vcs root property params
func (o *GetVcsRootPropertyParams) WithName(name string) *GetVcsRootPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get vcs root property params
func (o *GetVcsRootPropertyParams) SetName(name string) {
	o.Name = name
}

// WithVcsRootLocator adds the vcsRootLocator to the get vcs root property params
func (o *GetVcsRootPropertyParams) WithVcsRootLocator(vcsRootLocator string) *GetVcsRootPropertyParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the get vcs root property params
func (o *GetVcsRootPropertyParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
