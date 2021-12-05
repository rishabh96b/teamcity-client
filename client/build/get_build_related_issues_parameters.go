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

// NewGetBuildRelatedIssuesParams creates a new GetBuildRelatedIssuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildRelatedIssuesParams() *GetBuildRelatedIssuesParams {
	return &GetBuildRelatedIssuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildRelatedIssuesParamsWithTimeout creates a new GetBuildRelatedIssuesParams object
// with the ability to set a timeout on a request.
func NewGetBuildRelatedIssuesParamsWithTimeout(timeout time.Duration) *GetBuildRelatedIssuesParams {
	return &GetBuildRelatedIssuesParams{
		timeout: timeout,
	}
}

// NewGetBuildRelatedIssuesParamsWithContext creates a new GetBuildRelatedIssuesParams object
// with the ability to set a context for a request.
func NewGetBuildRelatedIssuesParamsWithContext(ctx context.Context) *GetBuildRelatedIssuesParams {
	return &GetBuildRelatedIssuesParams{
		Context: ctx,
	}
}

// NewGetBuildRelatedIssuesParamsWithHTTPClient creates a new GetBuildRelatedIssuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildRelatedIssuesParamsWithHTTPClient(client *http.Client) *GetBuildRelatedIssuesParams {
	return &GetBuildRelatedIssuesParams{
		HTTPClient: client,
	}
}

/* GetBuildRelatedIssuesParams contains all the parameters to send to the API endpoint
   for the get build related issues operation.

   Typically these are written to a http.Request.
*/
type GetBuildRelatedIssuesParams struct {

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

// WithDefaults hydrates default values in the get build related issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildRelatedIssuesParams) WithDefaults() *GetBuildRelatedIssuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build related issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildRelatedIssuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build related issues params
func (o *GetBuildRelatedIssuesParams) WithTimeout(timeout time.Duration) *GetBuildRelatedIssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build related issues params
func (o *GetBuildRelatedIssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build related issues params
func (o *GetBuildRelatedIssuesParams) WithContext(ctx context.Context) *GetBuildRelatedIssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build related issues params
func (o *GetBuildRelatedIssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build related issues params
func (o *GetBuildRelatedIssuesParams) WithHTTPClient(client *http.Client) *GetBuildRelatedIssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build related issues params
func (o *GetBuildRelatedIssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get build related issues params
func (o *GetBuildRelatedIssuesParams) WithBuildLocator(buildLocator string) *GetBuildRelatedIssuesParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get build related issues params
func (o *GetBuildRelatedIssuesParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get build related issues params
func (o *GetBuildRelatedIssuesParams) WithFields(fields *string) *GetBuildRelatedIssuesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build related issues params
func (o *GetBuildRelatedIssuesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildRelatedIssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
