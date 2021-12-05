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

// NewDeleteVcsRootParams creates a new DeleteVcsRootParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVcsRootParams() *DeleteVcsRootParams {
	return &DeleteVcsRootParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVcsRootParamsWithTimeout creates a new DeleteVcsRootParams object
// with the ability to set a timeout on a request.
func NewDeleteVcsRootParamsWithTimeout(timeout time.Duration) *DeleteVcsRootParams {
	return &DeleteVcsRootParams{
		timeout: timeout,
	}
}

// NewDeleteVcsRootParamsWithContext creates a new DeleteVcsRootParams object
// with the ability to set a context for a request.
func NewDeleteVcsRootParamsWithContext(ctx context.Context) *DeleteVcsRootParams {
	return &DeleteVcsRootParams{
		Context: ctx,
	}
}

// NewDeleteVcsRootParamsWithHTTPClient creates a new DeleteVcsRootParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVcsRootParamsWithHTTPClient(client *http.Client) *DeleteVcsRootParams {
	return &DeleteVcsRootParams{
		HTTPClient: client,
	}
}

/* DeleteVcsRootParams contains all the parameters to send to the API endpoint
   for the delete vcs root operation.

   Typically these are written to a http.Request.
*/
type DeleteVcsRootParams struct {

	// VcsRootLocator.
	//
	// Format: VcsRootLocator
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcsRootParams) WithDefaults() *DeleteVcsRootParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcsRootParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vcs root params
func (o *DeleteVcsRootParams) WithTimeout(timeout time.Duration) *DeleteVcsRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vcs root params
func (o *DeleteVcsRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vcs root params
func (o *DeleteVcsRootParams) WithContext(ctx context.Context) *DeleteVcsRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vcs root params
func (o *DeleteVcsRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vcs root params
func (o *DeleteVcsRootParams) WithHTTPClient(client *http.Client) *DeleteVcsRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vcs root params
func (o *DeleteVcsRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVcsRootLocator adds the vcsRootLocator to the delete vcs root params
func (o *DeleteVcsRootParams) WithVcsRootLocator(vcsRootLocator string) *DeleteVcsRootParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the delete vcs root params
func (o *DeleteVcsRootParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVcsRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}