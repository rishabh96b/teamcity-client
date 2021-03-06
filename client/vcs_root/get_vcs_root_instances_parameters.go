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

// NewGetVcsRootInstancesParams creates a new GetVcsRootInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootInstancesParams() *GetVcsRootInstancesParams {
	return &GetVcsRootInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootInstancesParamsWithTimeout creates a new GetVcsRootInstancesParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootInstancesParamsWithTimeout(timeout time.Duration) *GetVcsRootInstancesParams {
	return &GetVcsRootInstancesParams{
		timeout: timeout,
	}
}

// NewGetVcsRootInstancesParamsWithContext creates a new GetVcsRootInstancesParams object
// with the ability to set a context for a request.
func NewGetVcsRootInstancesParamsWithContext(ctx context.Context) *GetVcsRootInstancesParams {
	return &GetVcsRootInstancesParams{
		Context: ctx,
	}
}

// NewGetVcsRootInstancesParamsWithHTTPClient creates a new GetVcsRootInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootInstancesParamsWithHTTPClient(client *http.Client) *GetVcsRootInstancesParams {
	return &GetVcsRootInstancesParams{
		HTTPClient: client,
	}
}

/* GetVcsRootInstancesParams contains all the parameters to send to the API endpoint
   for the get vcs root instances operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootInstancesParams struct {

	// Fields.
	Fields *string

	// VcsRootLocator.
	//
	// Format: VcsRootLocator
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs root instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstancesParams) WithDefaults() *GetVcsRootInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root instances params
func (o *GetVcsRootInstancesParams) WithTimeout(timeout time.Duration) *GetVcsRootInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root instances params
func (o *GetVcsRootInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root instances params
func (o *GetVcsRootInstancesParams) WithContext(ctx context.Context) *GetVcsRootInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root instances params
func (o *GetVcsRootInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root instances params
func (o *GetVcsRootInstancesParams) WithHTTPClient(client *http.Client) *GetVcsRootInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root instances params
func (o *GetVcsRootInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get vcs root instances params
func (o *GetVcsRootInstancesParams) WithFields(fields *string) *GetVcsRootInstancesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root instances params
func (o *GetVcsRootInstancesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootLocator adds the vcsRootLocator to the get vcs root instances params
func (o *GetVcsRootInstancesParams) WithVcsRootLocator(vcsRootLocator string) *GetVcsRootInstancesParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the get vcs root instances params
func (o *GetVcsRootInstancesParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
