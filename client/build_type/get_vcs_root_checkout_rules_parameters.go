// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewGetVcsRootCheckoutRulesParams creates a new GetVcsRootCheckoutRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootCheckoutRulesParams() *GetVcsRootCheckoutRulesParams {
	return &GetVcsRootCheckoutRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootCheckoutRulesParamsWithTimeout creates a new GetVcsRootCheckoutRulesParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootCheckoutRulesParamsWithTimeout(timeout time.Duration) *GetVcsRootCheckoutRulesParams {
	return &GetVcsRootCheckoutRulesParams{
		timeout: timeout,
	}
}

// NewGetVcsRootCheckoutRulesParamsWithContext creates a new GetVcsRootCheckoutRulesParams object
// with the ability to set a context for a request.
func NewGetVcsRootCheckoutRulesParamsWithContext(ctx context.Context) *GetVcsRootCheckoutRulesParams {
	return &GetVcsRootCheckoutRulesParams{
		Context: ctx,
	}
}

// NewGetVcsRootCheckoutRulesParamsWithHTTPClient creates a new GetVcsRootCheckoutRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootCheckoutRulesParamsWithHTTPClient(client *http.Client) *GetVcsRootCheckoutRulesParams {
	return &GetVcsRootCheckoutRulesParams{
		HTTPClient: client,
	}
}

/* GetVcsRootCheckoutRulesParams contains all the parameters to send to the API endpoint
   for the get vcs root checkout rules operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootCheckoutRulesParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// VcsRootLocator.
	//
	// Format: VcsRootLocator
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs root checkout rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootCheckoutRulesParams) WithDefaults() *GetVcsRootCheckoutRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root checkout rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootCheckoutRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) WithTimeout(timeout time.Duration) *GetVcsRootCheckoutRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) WithContext(ctx context.Context) *GetVcsRootCheckoutRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) WithHTTPClient(client *http.Client) *GetVcsRootCheckoutRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) WithBtLocator(btLocator string) *GetVcsRootCheckoutRulesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithVcsRootLocator adds the vcsRootLocator to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) WithVcsRootLocator(vcsRootLocator string) *GetVcsRootCheckoutRulesParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the get vcs root checkout rules params
func (o *GetVcsRootCheckoutRulesParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootCheckoutRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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
