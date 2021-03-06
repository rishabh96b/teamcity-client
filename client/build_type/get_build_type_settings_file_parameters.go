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

// NewGetBuildTypeSettingsFileParams creates a new GetBuildTypeSettingsFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildTypeSettingsFileParams() *GetBuildTypeSettingsFileParams {
	return &GetBuildTypeSettingsFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildTypeSettingsFileParamsWithTimeout creates a new GetBuildTypeSettingsFileParams object
// with the ability to set a timeout on a request.
func NewGetBuildTypeSettingsFileParamsWithTimeout(timeout time.Duration) *GetBuildTypeSettingsFileParams {
	return &GetBuildTypeSettingsFileParams{
		timeout: timeout,
	}
}

// NewGetBuildTypeSettingsFileParamsWithContext creates a new GetBuildTypeSettingsFileParams object
// with the ability to set a context for a request.
func NewGetBuildTypeSettingsFileParamsWithContext(ctx context.Context) *GetBuildTypeSettingsFileParams {
	return &GetBuildTypeSettingsFileParams{
		Context: ctx,
	}
}

// NewGetBuildTypeSettingsFileParamsWithHTTPClient creates a new GetBuildTypeSettingsFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildTypeSettingsFileParamsWithHTTPClient(client *http.Client) *GetBuildTypeSettingsFileParams {
	return &GetBuildTypeSettingsFileParams{
		HTTPClient: client,
	}
}

/* GetBuildTypeSettingsFileParams contains all the parameters to send to the API endpoint
   for the get build type settings file operation.

   Typically these are written to a http.Request.
*/
type GetBuildTypeSettingsFileParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build type settings file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTypeSettingsFileParams) WithDefaults() *GetBuildTypeSettingsFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build type settings file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTypeSettingsFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) WithTimeout(timeout time.Duration) *GetBuildTypeSettingsFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) WithContext(ctx context.Context) *GetBuildTypeSettingsFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) WithHTTPClient(client *http.Client) *GetBuildTypeSettingsFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) WithBtLocator(btLocator string) *GetBuildTypeSettingsFileParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get build type settings file params
func (o *GetBuildTypeSettingsFileParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildTypeSettingsFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
