// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewGetFilesListForSubpathOfBuildParams creates a new GetFilesListForSubpathOfBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFilesListForSubpathOfBuildParams() *GetFilesListForSubpathOfBuildParams {
	return &GetFilesListForSubpathOfBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilesListForSubpathOfBuildParamsWithTimeout creates a new GetFilesListForSubpathOfBuildParams object
// with the ability to set a timeout on a request.
func NewGetFilesListForSubpathOfBuildParamsWithTimeout(timeout time.Duration) *GetFilesListForSubpathOfBuildParams {
	return &GetFilesListForSubpathOfBuildParams{
		timeout: timeout,
	}
}

// NewGetFilesListForSubpathOfBuildParamsWithContext creates a new GetFilesListForSubpathOfBuildParams object
// with the ability to set a context for a request.
func NewGetFilesListForSubpathOfBuildParamsWithContext(ctx context.Context) *GetFilesListForSubpathOfBuildParams {
	return &GetFilesListForSubpathOfBuildParams{
		Context: ctx,
	}
}

// NewGetFilesListForSubpathOfBuildParamsWithHTTPClient creates a new GetFilesListForSubpathOfBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFilesListForSubpathOfBuildParamsWithHTTPClient(client *http.Client) *GetFilesListForSubpathOfBuildParams {
	return &GetFilesListForSubpathOfBuildParams{
		HTTPClient: client,
	}
}

/* GetFilesListForSubpathOfBuildParams contains all the parameters to send to the API endpoint
   for the get files list for subpath of build operation.

   Typically these are written to a http.Request.
*/
type GetFilesListForSubpathOfBuildParams struct {

	// BasePath.
	BasePath *string

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	// Fields.
	Fields *string

	// Locator.
	Locator *string

	// LogBuildUsage.
	LogBuildUsage *bool

	// Path.
	Path string

	// ResolveParameters.
	ResolveParameters *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get files list for subpath of build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilesListForSubpathOfBuildParams) WithDefaults() *GetFilesListForSubpathOfBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get files list for subpath of build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilesListForSubpathOfBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithTimeout(timeout time.Duration) *GetFilesListForSubpathOfBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithContext(ctx context.Context) *GetFilesListForSubpathOfBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithHTTPClient(client *http.Client) *GetFilesListForSubpathOfBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasePath adds the basePath to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithBasePath(basePath *string) *GetFilesListForSubpathOfBuildParams {
	o.SetBasePath(basePath)
	return o
}

// SetBasePath adds the basePath to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetBasePath(basePath *string) {
	o.BasePath = basePath
}

// WithBuildLocator adds the buildLocator to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithBuildLocator(buildLocator string) *GetFilesListForSubpathOfBuildParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithFields(fields *string) *GetFilesListForSubpathOfBuildParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithLocator(locator *string) *GetFilesListForSubpathOfBuildParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithLogBuildUsage adds the logBuildUsage to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithLogBuildUsage(logBuildUsage *bool) *GetFilesListForSubpathOfBuildParams {
	o.SetLogBuildUsage(logBuildUsage)
	return o
}

// SetLogBuildUsage adds the logBuildUsage to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetLogBuildUsage(logBuildUsage *bool) {
	o.LogBuildUsage = logBuildUsage
}

// WithPath adds the path to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithPath(path string) *GetFilesListForSubpathOfBuildParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetPath(path string) {
	o.Path = path
}

// WithResolveParameters adds the resolveParameters to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) WithResolveParameters(resolveParameters *bool) *GetFilesListForSubpathOfBuildParams {
	o.SetResolveParameters(resolveParameters)
	return o
}

// SetResolveParameters adds the resolveParameters to the get files list for subpath of build params
func (o *GetFilesListForSubpathOfBuildParams) SetResolveParameters(resolveParameters *bool) {
	o.ResolveParameters = resolveParameters
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilesListForSubpathOfBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

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

	if o.LogBuildUsage != nil {

		// query param logBuildUsage
		var qrLogBuildUsage bool

		if o.LogBuildUsage != nil {
			qrLogBuildUsage = *o.LogBuildUsage
		}
		qLogBuildUsage := swag.FormatBool(qrLogBuildUsage)
		if qLogBuildUsage != "" {

			if err := r.SetQueryParam("logBuildUsage", qLogBuildUsage); err != nil {
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
