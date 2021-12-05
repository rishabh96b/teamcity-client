// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetProjectParams creates a new GetProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectParams() *GetProjectParams {
	return &GetProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectParamsWithTimeout creates a new GetProjectParams object
// with the ability to set a timeout on a request.
func NewGetProjectParamsWithTimeout(timeout time.Duration) *GetProjectParams {
	return &GetProjectParams{
		timeout: timeout,
	}
}

// NewGetProjectParamsWithContext creates a new GetProjectParams object
// with the ability to set a context for a request.
func NewGetProjectParamsWithContext(ctx context.Context) *GetProjectParams {
	return &GetProjectParams{
		Context: ctx,
	}
}

// NewGetProjectParamsWithHTTPClient creates a new GetProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectParamsWithHTTPClient(client *http.Client) *GetProjectParams {
	return &GetProjectParams{
		HTTPClient: client,
	}
}

/* GetProjectParams contains all the parameters to send to the API endpoint
   for the get project operation.

   Typically these are written to a http.Request.
*/
type GetProjectParams struct {

	// Fields.
	Fields *string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) WithDefaults() *GetProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project params
func (o *GetProjectParams) WithTimeout(timeout time.Duration) *GetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project params
func (o *GetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project params
func (o *GetProjectParams) WithContext(ctx context.Context) *GetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project params
func (o *GetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) WithHTTPClient(client *http.Client) *GetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get project params
func (o *GetProjectParams) WithFields(fields *string) *GetProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get project params
func (o *GetProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the get project params
func (o *GetProjectParams) WithProjectLocator(projectLocator string) *GetProjectParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get project params
func (o *GetProjectParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}