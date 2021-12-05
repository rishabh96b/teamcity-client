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

// NewGetVcsRootInstanceRepositoryStateParams creates a new GetVcsRootInstanceRepositoryStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootInstanceRepositoryStateParams() *GetVcsRootInstanceRepositoryStateParams {
	return &GetVcsRootInstanceRepositoryStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootInstanceRepositoryStateParamsWithTimeout creates a new GetVcsRootInstanceRepositoryStateParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootInstanceRepositoryStateParamsWithTimeout(timeout time.Duration) *GetVcsRootInstanceRepositoryStateParams {
	return &GetVcsRootInstanceRepositoryStateParams{
		timeout: timeout,
	}
}

// NewGetVcsRootInstanceRepositoryStateParamsWithContext creates a new GetVcsRootInstanceRepositoryStateParams object
// with the ability to set a context for a request.
func NewGetVcsRootInstanceRepositoryStateParamsWithContext(ctx context.Context) *GetVcsRootInstanceRepositoryStateParams {
	return &GetVcsRootInstanceRepositoryStateParams{
		Context: ctx,
	}
}

// NewGetVcsRootInstanceRepositoryStateParamsWithHTTPClient creates a new GetVcsRootInstanceRepositoryStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootInstanceRepositoryStateParamsWithHTTPClient(client *http.Client) *GetVcsRootInstanceRepositoryStateParams {
	return &GetVcsRootInstanceRepositoryStateParams{
		HTTPClient: client,
	}
}

/* GetVcsRootInstanceRepositoryStateParams contains all the parameters to send to the API endpoint
   for the get vcs root instance repository state operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootInstanceRepositoryStateParams struct {

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

// WithDefaults hydrates default values in the get vcs root instance repository state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstanceRepositoryStateParams) WithDefaults() *GetVcsRootInstanceRepositoryStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root instance repository state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstanceRepositoryStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) WithTimeout(timeout time.Duration) *GetVcsRootInstanceRepositoryStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) WithContext(ctx context.Context) *GetVcsRootInstanceRepositoryStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) WithHTTPClient(client *http.Client) *GetVcsRootInstanceRepositoryStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) WithFields(fields *string) *GetVcsRootInstanceRepositoryStateParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *GetVcsRootInstanceRepositoryStateParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get vcs root instance repository state params
func (o *GetVcsRootInstanceRepositoryStateParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootInstanceRepositoryStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
