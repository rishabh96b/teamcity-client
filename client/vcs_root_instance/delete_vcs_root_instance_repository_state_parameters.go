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

// NewDeleteVcsRootInstanceRepositoryStateParams creates a new DeleteVcsRootInstanceRepositoryStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVcsRootInstanceRepositoryStateParams() *DeleteVcsRootInstanceRepositoryStateParams {
	return &DeleteVcsRootInstanceRepositoryStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVcsRootInstanceRepositoryStateParamsWithTimeout creates a new DeleteVcsRootInstanceRepositoryStateParams object
// with the ability to set a timeout on a request.
func NewDeleteVcsRootInstanceRepositoryStateParamsWithTimeout(timeout time.Duration) *DeleteVcsRootInstanceRepositoryStateParams {
	return &DeleteVcsRootInstanceRepositoryStateParams{
		timeout: timeout,
	}
}

// NewDeleteVcsRootInstanceRepositoryStateParamsWithContext creates a new DeleteVcsRootInstanceRepositoryStateParams object
// with the ability to set a context for a request.
func NewDeleteVcsRootInstanceRepositoryStateParamsWithContext(ctx context.Context) *DeleteVcsRootInstanceRepositoryStateParams {
	return &DeleteVcsRootInstanceRepositoryStateParams{
		Context: ctx,
	}
}

// NewDeleteVcsRootInstanceRepositoryStateParamsWithHTTPClient creates a new DeleteVcsRootInstanceRepositoryStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVcsRootInstanceRepositoryStateParamsWithHTTPClient(client *http.Client) *DeleteVcsRootInstanceRepositoryStateParams {
	return &DeleteVcsRootInstanceRepositoryStateParams{
		HTTPClient: client,
	}
}

/* DeleteVcsRootInstanceRepositoryStateParams contains all the parameters to send to the API endpoint
   for the delete vcs root instance repository state operation.

   Typically these are written to a http.Request.
*/
type DeleteVcsRootInstanceRepositoryStateParams struct {

	// VcsRootInstanceLocator.
	//
	// Format: VcsRootInstanceLocator
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vcs root instance repository state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcsRootInstanceRepositoryStateParams) WithDefaults() *DeleteVcsRootInstanceRepositoryStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vcs root instance repository state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcsRootInstanceRepositoryStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) WithTimeout(timeout time.Duration) *DeleteVcsRootInstanceRepositoryStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) WithContext(ctx context.Context) *DeleteVcsRootInstanceRepositoryStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) WithHTTPClient(client *http.Client) *DeleteVcsRootInstanceRepositoryStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *DeleteVcsRootInstanceRepositoryStateParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the delete vcs root instance repository state params
func (o *DeleteVcsRootInstanceRepositoryStateParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVcsRootInstanceRepositoryStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vcsRootInstanceLocator
	if err := r.SetPathParam("vcsRootInstanceLocator", o.VcsRootInstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
