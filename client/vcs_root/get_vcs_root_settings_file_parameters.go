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

// NewGetVcsRootSettingsFileParams creates a new GetVcsRootSettingsFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootSettingsFileParams() *GetVcsRootSettingsFileParams {
	return &GetVcsRootSettingsFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootSettingsFileParamsWithTimeout creates a new GetVcsRootSettingsFileParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootSettingsFileParamsWithTimeout(timeout time.Duration) *GetVcsRootSettingsFileParams {
	return &GetVcsRootSettingsFileParams{
		timeout: timeout,
	}
}

// NewGetVcsRootSettingsFileParamsWithContext creates a new GetVcsRootSettingsFileParams object
// with the ability to set a context for a request.
func NewGetVcsRootSettingsFileParamsWithContext(ctx context.Context) *GetVcsRootSettingsFileParams {
	return &GetVcsRootSettingsFileParams{
		Context: ctx,
	}
}

// NewGetVcsRootSettingsFileParamsWithHTTPClient creates a new GetVcsRootSettingsFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootSettingsFileParamsWithHTTPClient(client *http.Client) *GetVcsRootSettingsFileParams {
	return &GetVcsRootSettingsFileParams{
		HTTPClient: client,
	}
}

/* GetVcsRootSettingsFileParams contains all the parameters to send to the API endpoint
   for the get vcs root settings file operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootSettingsFileParams struct {

	// VcsRootLocator.
	//
	// Format: VcsRootLocator
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs root settings file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootSettingsFileParams) WithDefaults() *GetVcsRootSettingsFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root settings file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootSettingsFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) WithTimeout(timeout time.Duration) *GetVcsRootSettingsFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) WithContext(ctx context.Context) *GetVcsRootSettingsFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) WithHTTPClient(client *http.Client) *GetVcsRootSettingsFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVcsRootLocator adds the vcsRootLocator to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) WithVcsRootLocator(vcsRootLocator string) *GetVcsRootSettingsFileParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the get vcs root settings file params
func (o *GetVcsRootSettingsFileParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootSettingsFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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