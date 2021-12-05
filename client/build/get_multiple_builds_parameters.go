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

// NewGetMultipleBuildsParams creates a new GetMultipleBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMultipleBuildsParams() *GetMultipleBuildsParams {
	return &GetMultipleBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMultipleBuildsParamsWithTimeout creates a new GetMultipleBuildsParams object
// with the ability to set a timeout on a request.
func NewGetMultipleBuildsParamsWithTimeout(timeout time.Duration) *GetMultipleBuildsParams {
	return &GetMultipleBuildsParams{
		timeout: timeout,
	}
}

// NewGetMultipleBuildsParamsWithContext creates a new GetMultipleBuildsParams object
// with the ability to set a context for a request.
func NewGetMultipleBuildsParamsWithContext(ctx context.Context) *GetMultipleBuildsParams {
	return &GetMultipleBuildsParams{
		Context: ctx,
	}
}

// NewGetMultipleBuildsParamsWithHTTPClient creates a new GetMultipleBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMultipleBuildsParamsWithHTTPClient(client *http.Client) *GetMultipleBuildsParams {
	return &GetMultipleBuildsParams{
		HTTPClient: client,
	}
}

/* GetMultipleBuildsParams contains all the parameters to send to the API endpoint
   for the get multiple builds operation.

   Typically these are written to a http.Request.
*/
type GetMultipleBuildsParams struct {

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

// WithDefaults hydrates default values in the get multiple builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMultipleBuildsParams) WithDefaults() *GetMultipleBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get multiple builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMultipleBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get multiple builds params
func (o *GetMultipleBuildsParams) WithTimeout(timeout time.Duration) *GetMultipleBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get multiple builds params
func (o *GetMultipleBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get multiple builds params
func (o *GetMultipleBuildsParams) WithContext(ctx context.Context) *GetMultipleBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get multiple builds params
func (o *GetMultipleBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get multiple builds params
func (o *GetMultipleBuildsParams) WithHTTPClient(client *http.Client) *GetMultipleBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get multiple builds params
func (o *GetMultipleBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get multiple builds params
func (o *GetMultipleBuildsParams) WithBuildLocator(buildLocator string) *GetMultipleBuildsParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get multiple builds params
func (o *GetMultipleBuildsParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get multiple builds params
func (o *GetMultipleBuildsParams) WithFields(fields *string) *GetMultipleBuildsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get multiple builds params
func (o *GetMultipleBuildsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetMultipleBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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