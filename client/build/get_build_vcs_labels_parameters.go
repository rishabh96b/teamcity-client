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

// NewGetBuildVcsLabelsParams creates a new GetBuildVcsLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildVcsLabelsParams() *GetBuildVcsLabelsParams {
	return &GetBuildVcsLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildVcsLabelsParamsWithTimeout creates a new GetBuildVcsLabelsParams object
// with the ability to set a timeout on a request.
func NewGetBuildVcsLabelsParamsWithTimeout(timeout time.Duration) *GetBuildVcsLabelsParams {
	return &GetBuildVcsLabelsParams{
		timeout: timeout,
	}
}

// NewGetBuildVcsLabelsParamsWithContext creates a new GetBuildVcsLabelsParams object
// with the ability to set a context for a request.
func NewGetBuildVcsLabelsParamsWithContext(ctx context.Context) *GetBuildVcsLabelsParams {
	return &GetBuildVcsLabelsParams{
		Context: ctx,
	}
}

// NewGetBuildVcsLabelsParamsWithHTTPClient creates a new GetBuildVcsLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildVcsLabelsParamsWithHTTPClient(client *http.Client) *GetBuildVcsLabelsParams {
	return &GetBuildVcsLabelsParams{
		HTTPClient: client,
	}
}

/* GetBuildVcsLabelsParams contains all the parameters to send to the API endpoint
   for the get build vcs labels operation.

   Typically these are written to a http.Request.
*/
type GetBuildVcsLabelsParams struct {

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

// WithDefaults hydrates default values in the get build vcs labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildVcsLabelsParams) WithDefaults() *GetBuildVcsLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build vcs labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildVcsLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) WithTimeout(timeout time.Duration) *GetBuildVcsLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) WithContext(ctx context.Context) *GetBuildVcsLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) WithHTTPClient(client *http.Client) *GetBuildVcsLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) WithBuildLocator(buildLocator string) *GetBuildVcsLabelsParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) WithFields(fields *string) *GetBuildVcsLabelsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build vcs labels params
func (o *GetBuildVcsLabelsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildVcsLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
