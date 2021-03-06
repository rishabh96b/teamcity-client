// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

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

// NewGetVcsRootInstanceParams creates a new GetVcsRootInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootInstanceParams() *GetVcsRootInstanceParams {
	return &GetVcsRootInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootInstanceParamsWithTimeout creates a new GetVcsRootInstanceParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootInstanceParamsWithTimeout(timeout time.Duration) *GetVcsRootInstanceParams {
	return &GetVcsRootInstanceParams{
		timeout: timeout,
	}
}

// NewGetVcsRootInstanceParamsWithContext creates a new GetVcsRootInstanceParams object
// with the ability to set a context for a request.
func NewGetVcsRootInstanceParamsWithContext(ctx context.Context) *GetVcsRootInstanceParams {
	return &GetVcsRootInstanceParams{
		Context: ctx,
	}
}

// NewGetVcsRootInstanceParamsWithHTTPClient creates a new GetVcsRootInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootInstanceParamsWithHTTPClient(client *http.Client) *GetVcsRootInstanceParams {
	return &GetVcsRootInstanceParams{
		HTTPClient: client,
	}
}

/* GetVcsRootInstanceParams contains all the parameters to send to the API endpoint
   for the get vcs root instance operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootInstanceParams struct {

	// Fields.
	Fields *string

	// VcsRootInstanceLocator.
	//
	// Format: VcsRootInstanceLocator
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs root instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstanceParams) WithDefaults() *GetVcsRootInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root instance params
func (o *GetVcsRootInstanceParams) WithTimeout(timeout time.Duration) *GetVcsRootInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root instance params
func (o *GetVcsRootInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root instance params
func (o *GetVcsRootInstanceParams) WithContext(ctx context.Context) *GetVcsRootInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root instance params
func (o *GetVcsRootInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root instance params
func (o *GetVcsRootInstanceParams) WithHTTPClient(client *http.Client) *GetVcsRootInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root instance params
func (o *GetVcsRootInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get vcs root instance params
func (o *GetVcsRootInstanceParams) WithFields(fields *string) *GetVcsRootInstanceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root instance params
func (o *GetVcsRootInstanceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get vcs root instance params
func (o *GetVcsRootInstanceParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *GetVcsRootInstanceParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get vcs root instance params
func (o *GetVcsRootInstanceParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param vcsRootInstanceLocator
	if err := r.SetPathParam("vcsRootInstanceLocator", o.VcsRootInstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
