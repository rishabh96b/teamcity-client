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
	"github.com/go-openapi/swag"
)

// NewDownloadFileOfBuildTypeParams creates a new DownloadFileOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadFileOfBuildTypeParams() *DownloadFileOfBuildTypeParams {
	return &DownloadFileOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadFileOfBuildTypeParamsWithTimeout creates a new DownloadFileOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewDownloadFileOfBuildTypeParamsWithTimeout(timeout time.Duration) *DownloadFileOfBuildTypeParams {
	return &DownloadFileOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewDownloadFileOfBuildTypeParamsWithContext creates a new DownloadFileOfBuildTypeParams object
// with the ability to set a context for a request.
func NewDownloadFileOfBuildTypeParamsWithContext(ctx context.Context) *DownloadFileOfBuildTypeParams {
	return &DownloadFileOfBuildTypeParams{
		Context: ctx,
	}
}

// NewDownloadFileOfBuildTypeParamsWithHTTPClient creates a new DownloadFileOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadFileOfBuildTypeParamsWithHTTPClient(client *http.Client) *DownloadFileOfBuildTypeParams {
	return &DownloadFileOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* DownloadFileOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the download file of build type operation.

   Typically these are written to a http.Request.
*/
type DownloadFileOfBuildTypeParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Path.
	Path string

	// ResolveParameters.
	ResolveParameters *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download file of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileOfBuildTypeParams) WithDefaults() *DownloadFileOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download file of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithTimeout(timeout time.Duration) *DownloadFileOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithContext(ctx context.Context) *DownloadFileOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithHTTPClient(client *http.Client) *DownloadFileOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithBtLocator(btLocator string) *DownloadFileOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithPath adds the path to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithPath(path string) *DownloadFileOfBuildTypeParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetPath(path string) {
	o.Path = path
}

// WithResolveParameters adds the resolveParameters to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) WithResolveParameters(resolveParameters *bool) *DownloadFileOfBuildTypeParams {
	o.SetResolveParameters(resolveParameters)
	return o
}

// SetResolveParameters adds the resolveParameters to the download file of build type params
func (o *DownloadFileOfBuildTypeParams) SetResolveParameters(resolveParameters *bool) {
	o.ResolveParameters = resolveParameters
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadFileOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.ResolveParameters != nil {

		// query param resolveParameters
		var qrResolveParameters bool

		if o.ResolveParameters != nil {
			qrResolveParameters = *o.ResolveParameters
		}
		qResolveParameters := swag.FormatBool(qrResolveParameters)
		if qResolveParameters != "" {

			if err := r.SetQueryParam("resolveParameters", qResolveParameters); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
