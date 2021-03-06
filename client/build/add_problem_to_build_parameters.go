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

// NewAddProblemToBuildParams creates a new AddProblemToBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddProblemToBuildParams() *AddProblemToBuildParams {
	return &AddProblemToBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddProblemToBuildParamsWithTimeout creates a new AddProblemToBuildParams object
// with the ability to set a timeout on a request.
func NewAddProblemToBuildParamsWithTimeout(timeout time.Duration) *AddProblemToBuildParams {
	return &AddProblemToBuildParams{
		timeout: timeout,
	}
}

// NewAddProblemToBuildParamsWithContext creates a new AddProblemToBuildParams object
// with the ability to set a context for a request.
func NewAddProblemToBuildParamsWithContext(ctx context.Context) *AddProblemToBuildParams {
	return &AddProblemToBuildParams{
		Context: ctx,
	}
}

// NewAddProblemToBuildParamsWithHTTPClient creates a new AddProblemToBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddProblemToBuildParamsWithHTTPClient(client *http.Client) *AddProblemToBuildParams {
	return &AddProblemToBuildParams{
		HTTPClient: client,
	}
}

/* AddProblemToBuildParams contains all the parameters to send to the API endpoint
   for the add problem to build operation.

   Typically these are written to a http.Request.
*/
type AddProblemToBuildParams struct {

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

// WithDefaults hydrates default values in the add problem to build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProblemToBuildParams) WithDefaults() *AddProblemToBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add problem to build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProblemToBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add problem to build params
func (o *AddProblemToBuildParams) WithTimeout(timeout time.Duration) *AddProblemToBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add problem to build params
func (o *AddProblemToBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add problem to build params
func (o *AddProblemToBuildParams) WithContext(ctx context.Context) *AddProblemToBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add problem to build params
func (o *AddProblemToBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add problem to build params
func (o *AddProblemToBuildParams) WithHTTPClient(client *http.Client) *AddProblemToBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add problem to build params
func (o *AddProblemToBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add problem to build params
func (o *AddProblemToBuildParams) WithBody(body string) *AddProblemToBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add problem to build params
func (o *AddProblemToBuildParams) SetBody(body string) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the add problem to build params
func (o *AddProblemToBuildParams) WithBuildLocator(buildLocator string) *AddProblemToBuildParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the add problem to build params
func (o *AddProblemToBuildParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the add problem to build params
func (o *AddProblemToBuildParams) WithFields(fields *string) *AddProblemToBuildParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add problem to build params
func (o *AddProblemToBuildParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddProblemToBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
