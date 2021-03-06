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

// NewAddBuildVcsLabelParams creates a new AddBuildVcsLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBuildVcsLabelParams() *AddBuildVcsLabelParams {
	return &AddBuildVcsLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBuildVcsLabelParamsWithTimeout creates a new AddBuildVcsLabelParams object
// with the ability to set a timeout on a request.
func NewAddBuildVcsLabelParamsWithTimeout(timeout time.Duration) *AddBuildVcsLabelParams {
	return &AddBuildVcsLabelParams{
		timeout: timeout,
	}
}

// NewAddBuildVcsLabelParamsWithContext creates a new AddBuildVcsLabelParams object
// with the ability to set a context for a request.
func NewAddBuildVcsLabelParamsWithContext(ctx context.Context) *AddBuildVcsLabelParams {
	return &AddBuildVcsLabelParams{
		Context: ctx,
	}
}

// NewAddBuildVcsLabelParamsWithHTTPClient creates a new AddBuildVcsLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBuildVcsLabelParamsWithHTTPClient(client *http.Client) *AddBuildVcsLabelParams {
	return &AddBuildVcsLabelParams{
		HTTPClient: client,
	}
}

/* AddBuildVcsLabelParams contains all the parameters to send to the API endpoint
   for the add build vcs label operation.

   Typically these are written to a http.Request.
*/
type AddBuildVcsLabelParams struct {

	// Body.
	Body string

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	// Fields.
	Fields *string

	// Locator.
	//
	// Format: VcsRootInstanceLocator
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add build vcs label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBuildVcsLabelParams) WithDefaults() *AddBuildVcsLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add build vcs label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBuildVcsLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithTimeout(timeout time.Duration) *AddBuildVcsLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithContext(ctx context.Context) *AddBuildVcsLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithHTTPClient(client *http.Client) *AddBuildVcsLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithBody(body string) *AddBuildVcsLabelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetBody(body string) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithBuildLocator(buildLocator string) *AddBuildVcsLabelParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithFields(fields *string) *AddBuildVcsLabelParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the add build vcs label params
func (o *AddBuildVcsLabelParams) WithLocator(locator *string) *AddBuildVcsLabelParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the add build vcs label params
func (o *AddBuildVcsLabelParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *AddBuildVcsLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
