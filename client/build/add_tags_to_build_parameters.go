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

	"github.com/rishabh96b/teamcity-client/models"
)

// NewAddTagsToBuildParams creates a new AddTagsToBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddTagsToBuildParams() *AddTagsToBuildParams {
	return &AddTagsToBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddTagsToBuildParamsWithTimeout creates a new AddTagsToBuildParams object
// with the ability to set a timeout on a request.
func NewAddTagsToBuildParamsWithTimeout(timeout time.Duration) *AddTagsToBuildParams {
	return &AddTagsToBuildParams{
		timeout: timeout,
	}
}

// NewAddTagsToBuildParamsWithContext creates a new AddTagsToBuildParams object
// with the ability to set a context for a request.
func NewAddTagsToBuildParamsWithContext(ctx context.Context) *AddTagsToBuildParams {
	return &AddTagsToBuildParams{
		Context: ctx,
	}
}

// NewAddTagsToBuildParamsWithHTTPClient creates a new AddTagsToBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddTagsToBuildParamsWithHTTPClient(client *http.Client) *AddTagsToBuildParams {
	return &AddTagsToBuildParams{
		HTTPClient: client,
	}
}

/* AddTagsToBuildParams contains all the parameters to send to the API endpoint
   for the add tags to build operation.

   Typically these are written to a http.Request.
*/
type AddTagsToBuildParams struct {

	// Body.
	Body *models.Tags

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

// WithDefaults hydrates default values in the add tags to build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTagsToBuildParams) WithDefaults() *AddTagsToBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add tags to build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTagsToBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add tags to build params
func (o *AddTagsToBuildParams) WithTimeout(timeout time.Duration) *AddTagsToBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add tags to build params
func (o *AddTagsToBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add tags to build params
func (o *AddTagsToBuildParams) WithContext(ctx context.Context) *AddTagsToBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add tags to build params
func (o *AddTagsToBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add tags to build params
func (o *AddTagsToBuildParams) WithHTTPClient(client *http.Client) *AddTagsToBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add tags to build params
func (o *AddTagsToBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add tags to build params
func (o *AddTagsToBuildParams) WithBody(body *models.Tags) *AddTagsToBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add tags to build params
func (o *AddTagsToBuildParams) SetBody(body *models.Tags) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the add tags to build params
func (o *AddTagsToBuildParams) WithBuildLocator(buildLocator string) *AddTagsToBuildParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the add tags to build params
func (o *AddTagsToBuildParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the add tags to build params
func (o *AddTagsToBuildParams) WithFields(fields *string) *AddTagsToBuildParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add tags to build params
func (o *AddTagsToBuildParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddTagsToBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}