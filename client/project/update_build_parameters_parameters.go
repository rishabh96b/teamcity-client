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

	"github.com/rishabh96b/teamcity-client/models"
)

// NewUpdateBuildParametersParams creates a new UpdateBuildParametersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBuildParametersParams() *UpdateBuildParametersParams {
	return &UpdateBuildParametersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBuildParametersParamsWithTimeout creates a new UpdateBuildParametersParams object
// with the ability to set a timeout on a request.
func NewUpdateBuildParametersParamsWithTimeout(timeout time.Duration) *UpdateBuildParametersParams {
	return &UpdateBuildParametersParams{
		timeout: timeout,
	}
}

// NewUpdateBuildParametersParamsWithContext creates a new UpdateBuildParametersParams object
// with the ability to set a context for a request.
func NewUpdateBuildParametersParamsWithContext(ctx context.Context) *UpdateBuildParametersParams {
	return &UpdateBuildParametersParams{
		Context: ctx,
	}
}

// NewUpdateBuildParametersParamsWithHTTPClient creates a new UpdateBuildParametersParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBuildParametersParamsWithHTTPClient(client *http.Client) *UpdateBuildParametersParams {
	return &UpdateBuildParametersParams{
		HTTPClient: client,
	}
}

/* UpdateBuildParametersParams contains all the parameters to send to the API endpoint
   for the update build parameters operation.

   Typically these are written to a http.Request.
*/
type UpdateBuildParametersParams struct {

	// Body.
	Body *models.Properties

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

// WithDefaults hydrates default values in the update build parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBuildParametersParams) WithDefaults() *UpdateBuildParametersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update build parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBuildParametersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update build parameters params
func (o *UpdateBuildParametersParams) WithTimeout(timeout time.Duration) *UpdateBuildParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update build parameters params
func (o *UpdateBuildParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update build parameters params
func (o *UpdateBuildParametersParams) WithContext(ctx context.Context) *UpdateBuildParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update build parameters params
func (o *UpdateBuildParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update build parameters params
func (o *UpdateBuildParametersParams) WithHTTPClient(client *http.Client) *UpdateBuildParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update build parameters params
func (o *UpdateBuildParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update build parameters params
func (o *UpdateBuildParametersParams) WithBody(body *models.Properties) *UpdateBuildParametersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update build parameters params
func (o *UpdateBuildParametersParams) SetBody(body *models.Properties) {
	o.Body = body
}

// WithFields adds the fields to the update build parameters params
func (o *UpdateBuildParametersParams) WithFields(fields *string) *UpdateBuildParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update build parameters params
func (o *UpdateBuildParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the update build parameters params
func (o *UpdateBuildParametersParams) WithProjectLocator(projectLocator string) *UpdateBuildParametersParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the update build parameters params
func (o *UpdateBuildParametersParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBuildParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}