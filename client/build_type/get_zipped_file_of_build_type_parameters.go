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

// NewGetZippedFileOfBuildTypeParams creates a new GetZippedFileOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZippedFileOfBuildTypeParams() *GetZippedFileOfBuildTypeParams {
	return &GetZippedFileOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZippedFileOfBuildTypeParamsWithTimeout creates a new GetZippedFileOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewGetZippedFileOfBuildTypeParamsWithTimeout(timeout time.Duration) *GetZippedFileOfBuildTypeParams {
	return &GetZippedFileOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewGetZippedFileOfBuildTypeParamsWithContext creates a new GetZippedFileOfBuildTypeParams object
// with the ability to set a context for a request.
func NewGetZippedFileOfBuildTypeParamsWithContext(ctx context.Context) *GetZippedFileOfBuildTypeParams {
	return &GetZippedFileOfBuildTypeParams{
		Context: ctx,
	}
}

// NewGetZippedFileOfBuildTypeParamsWithHTTPClient creates a new GetZippedFileOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZippedFileOfBuildTypeParamsWithHTTPClient(client *http.Client) *GetZippedFileOfBuildTypeParams {
	return &GetZippedFileOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* GetZippedFileOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the get zipped file of build type operation.

   Typically these are written to a http.Request.
*/
type GetZippedFileOfBuildTypeParams struct {

	// BasePath.
	BasePath *string

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Locator.
	Locator *string

	// Name.
	Name *string

	// Path.
	Path string

	// ResolveParameters.
	ResolveParameters *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zipped file of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZippedFileOfBuildTypeParams) WithDefaults() *GetZippedFileOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zipped file of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZippedFileOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithTimeout(timeout time.Duration) *GetZippedFileOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithContext(ctx context.Context) *GetZippedFileOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithHTTPClient(client *http.Client) *GetZippedFileOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasePath adds the basePath to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithBasePath(basePath *string) *GetZippedFileOfBuildTypeParams {
	o.SetBasePath(basePath)
	return o
}

// SetBasePath adds the basePath to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetBasePath(basePath *string) {
	o.BasePath = basePath
}

// WithBtLocator adds the btLocator to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithBtLocator(btLocator string) *GetZippedFileOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithLocator adds the locator to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithLocator(locator *string) *GetZippedFileOfBuildTypeParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithName adds the name to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithName(name *string) *GetZippedFileOfBuildTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetName(name *string) {
	o.Name = name
}

// WithPath adds the path to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithPath(path string) *GetZippedFileOfBuildTypeParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetPath(path string) {
	o.Path = path
}

// WithResolveParameters adds the resolveParameters to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) WithResolveParameters(resolveParameters *bool) *GetZippedFileOfBuildTypeParams {
	o.SetResolveParameters(resolveParameters)
	return o
}

// SetResolveParameters adds the resolveParameters to the get zipped file of build type params
func (o *GetZippedFileOfBuildTypeParams) SetResolveParameters(resolveParameters *bool) {
	o.ResolveParameters = resolveParameters
}

// WriteToRequest writes these params to a swagger request
func (o *GetZippedFileOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BasePath != nil {

		// query param basePath
		var qrBasePath string

		if o.BasePath != nil {
			qrBasePath = *o.BasePath
		}
		qBasePath := qrBasePath
		if qBasePath != "" {

			if err := r.SetQueryParam("basePath", qBasePath); err != nil {
				return err
			}
		}
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if o.Locator != nil {

		// query param locator
		var qrLocator string

		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {

			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
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
