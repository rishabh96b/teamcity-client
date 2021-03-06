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

// NewAddBuildTypeParams creates a new AddBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBuildTypeParams() *AddBuildTypeParams {
	return &AddBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBuildTypeParamsWithTimeout creates a new AddBuildTypeParams object
// with the ability to set a timeout on a request.
func NewAddBuildTypeParamsWithTimeout(timeout time.Duration) *AddBuildTypeParams {
	return &AddBuildTypeParams{
		timeout: timeout,
	}
}

// NewAddBuildTypeParamsWithContext creates a new AddBuildTypeParams object
// with the ability to set a context for a request.
func NewAddBuildTypeParamsWithContext(ctx context.Context) *AddBuildTypeParams {
	return &AddBuildTypeParams{
		Context: ctx,
	}
}

// NewAddBuildTypeParamsWithHTTPClient creates a new AddBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBuildTypeParamsWithHTTPClient(client *http.Client) *AddBuildTypeParams {
	return &AddBuildTypeParams{
		HTTPClient: client,
	}
}

/* AddBuildTypeParams contains all the parameters to send to the API endpoint
   for the add build type operation.

   Typically these are written to a http.Request.
*/
type AddBuildTypeParams struct {

	// Body.
	Body *models.NewBuildTypeDescription

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

// WithDefaults hydrates default values in the add build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBuildTypeParams) WithDefaults() *AddBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add build type params
func (o *AddBuildTypeParams) WithTimeout(timeout time.Duration) *AddBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add build type params
func (o *AddBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add build type params
func (o *AddBuildTypeParams) WithContext(ctx context.Context) *AddBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add build type params
func (o *AddBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add build type params
func (o *AddBuildTypeParams) WithHTTPClient(client *http.Client) *AddBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add build type params
func (o *AddBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add build type params
func (o *AddBuildTypeParams) WithBody(body *models.NewBuildTypeDescription) *AddBuildTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add build type params
func (o *AddBuildTypeParams) SetBody(body *models.NewBuildTypeDescription) {
	o.Body = body
}

// WithFields adds the fields to the add build type params
func (o *AddBuildTypeParams) WithFields(fields *string) *AddBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add build type params
func (o *AddBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the add build type params
func (o *AddBuildTypeParams) WithProjectLocator(projectLocator string) *AddBuildTypeParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the add build type params
func (o *AddBuildTypeParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *AddBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
