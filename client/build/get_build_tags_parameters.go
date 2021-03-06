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
)

// NewGetBuildTagsParams creates a new GetBuildTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildTagsParams() *GetBuildTagsParams {
	return &GetBuildTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildTagsParamsWithTimeout creates a new GetBuildTagsParams object
// with the ability to set a timeout on a request.
func NewGetBuildTagsParamsWithTimeout(timeout time.Duration) *GetBuildTagsParams {
	return &GetBuildTagsParams{
		timeout: timeout,
	}
}

// NewGetBuildTagsParamsWithContext creates a new GetBuildTagsParams object
// with the ability to set a context for a request.
func NewGetBuildTagsParamsWithContext(ctx context.Context) *GetBuildTagsParams {
	return &GetBuildTagsParams{
		Context: ctx,
	}
}

// NewGetBuildTagsParamsWithHTTPClient creates a new GetBuildTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildTagsParamsWithHTTPClient(client *http.Client) *GetBuildTagsParams {
	return &GetBuildTagsParams{
		HTTPClient: client,
	}
}

/* GetBuildTagsParams contains all the parameters to send to the API endpoint
   for the get build tags operation.

   Typically these are written to a http.Request.
*/
type GetBuildTagsParams struct {

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	// Fields.
	Fields *string

	// Locator.
	//
	// Format: TagLocator
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get build tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTagsParams) WithDefaults() *GetBuildTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build tags params
func (o *GetBuildTagsParams) WithTimeout(timeout time.Duration) *GetBuildTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build tags params
func (o *GetBuildTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build tags params
func (o *GetBuildTagsParams) WithContext(ctx context.Context) *GetBuildTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build tags params
func (o *GetBuildTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build tags params
func (o *GetBuildTagsParams) WithHTTPClient(client *http.Client) *GetBuildTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build tags params
func (o *GetBuildTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get build tags params
func (o *GetBuildTagsParams) WithBuildLocator(buildLocator string) *GetBuildTagsParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get build tags params
func (o *GetBuildTagsParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get build tags params
func (o *GetBuildTagsParams) WithFields(fields *string) *GetBuildTagsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build tags params
func (o *GetBuildTagsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get build tags params
func (o *GetBuildTagsParams) WithLocator(locator *string) *GetBuildTagsParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get build tags params
func (o *GetBuildTagsParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
