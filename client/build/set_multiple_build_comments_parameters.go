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

// NewSetMultipleBuildCommentsParams creates a new SetMultipleBuildCommentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetMultipleBuildCommentsParams() *SetMultipleBuildCommentsParams {
	return &SetMultipleBuildCommentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetMultipleBuildCommentsParamsWithTimeout creates a new SetMultipleBuildCommentsParams object
// with the ability to set a timeout on a request.
func NewSetMultipleBuildCommentsParamsWithTimeout(timeout time.Duration) *SetMultipleBuildCommentsParams {
	return &SetMultipleBuildCommentsParams{
		timeout: timeout,
	}
}

// NewSetMultipleBuildCommentsParamsWithContext creates a new SetMultipleBuildCommentsParams object
// with the ability to set a context for a request.
func NewSetMultipleBuildCommentsParamsWithContext(ctx context.Context) *SetMultipleBuildCommentsParams {
	return &SetMultipleBuildCommentsParams{
		Context: ctx,
	}
}

// NewSetMultipleBuildCommentsParamsWithHTTPClient creates a new SetMultipleBuildCommentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetMultipleBuildCommentsParamsWithHTTPClient(client *http.Client) *SetMultipleBuildCommentsParams {
	return &SetMultipleBuildCommentsParams{
		HTTPClient: client,
	}
}

/* SetMultipleBuildCommentsParams contains all the parameters to send to the API endpoint
   for the set multiple build comments operation.

   Typically these are written to a http.Request.
*/
type SetMultipleBuildCommentsParams struct {

	// Body.
	Body string

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set multiple build comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetMultipleBuildCommentsParams) WithDefaults() *SetMultipleBuildCommentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set multiple build comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetMultipleBuildCommentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithTimeout(timeout time.Duration) *SetMultipleBuildCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithContext(ctx context.Context) *SetMultipleBuildCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithHTTPClient(client *http.Client) *SetMultipleBuildCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithBody(body string) *SetMultipleBuildCommentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetBody(body string) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithBuildLocator(buildLocator string) *SetMultipleBuildCommentsParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) WithFields(fields *string) *SetMultipleBuildCommentsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the set multiple build comments params
func (o *SetMultipleBuildCommentsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *SetMultipleBuildCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
