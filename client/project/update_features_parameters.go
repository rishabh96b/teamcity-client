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

// NewUpdateFeaturesParams creates a new UpdateFeaturesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFeaturesParams() *UpdateFeaturesParams {
	return &UpdateFeaturesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFeaturesParamsWithTimeout creates a new UpdateFeaturesParams object
// with the ability to set a timeout on a request.
func NewUpdateFeaturesParamsWithTimeout(timeout time.Duration) *UpdateFeaturesParams {
	return &UpdateFeaturesParams{
		timeout: timeout,
	}
}

// NewUpdateFeaturesParamsWithContext creates a new UpdateFeaturesParams object
// with the ability to set a context for a request.
func NewUpdateFeaturesParamsWithContext(ctx context.Context) *UpdateFeaturesParams {
	return &UpdateFeaturesParams{
		Context: ctx,
	}
}

// NewUpdateFeaturesParamsWithHTTPClient creates a new UpdateFeaturesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFeaturesParamsWithHTTPClient(client *http.Client) *UpdateFeaturesParams {
	return &UpdateFeaturesParams{
		HTTPClient: client,
	}
}

/* UpdateFeaturesParams contains all the parameters to send to the API endpoint
   for the update features operation.

   Typically these are written to a http.Request.
*/
type UpdateFeaturesParams struct {

	// Body.
	Body *models.ProjectFeatures

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

// WithDefaults hydrates default values in the update features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFeaturesParams) WithDefaults() *UpdateFeaturesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFeaturesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update features params
func (o *UpdateFeaturesParams) WithTimeout(timeout time.Duration) *UpdateFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update features params
func (o *UpdateFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update features params
func (o *UpdateFeaturesParams) WithContext(ctx context.Context) *UpdateFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update features params
func (o *UpdateFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update features params
func (o *UpdateFeaturesParams) WithHTTPClient(client *http.Client) *UpdateFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update features params
func (o *UpdateFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update features params
func (o *UpdateFeaturesParams) WithBody(body *models.ProjectFeatures) *UpdateFeaturesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update features params
func (o *UpdateFeaturesParams) SetBody(body *models.ProjectFeatures) {
	o.Body = body
}

// WithFields adds the fields to the update features params
func (o *UpdateFeaturesParams) WithFields(fields *string) *UpdateFeaturesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update features params
func (o *UpdateFeaturesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the update features params
func (o *UpdateFeaturesParams) WithProjectLocator(projectLocator string) *UpdateFeaturesParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the update features params
func (o *UpdateFeaturesParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
