// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewCreateBuildTypeParams creates a new CreateBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBuildTypeParams() *CreateBuildTypeParams {
	return &CreateBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBuildTypeParamsWithTimeout creates a new CreateBuildTypeParams object
// with the ability to set a timeout on a request.
func NewCreateBuildTypeParamsWithTimeout(timeout time.Duration) *CreateBuildTypeParams {
	return &CreateBuildTypeParams{
		timeout: timeout,
	}
}

// NewCreateBuildTypeParamsWithContext creates a new CreateBuildTypeParams object
// with the ability to set a context for a request.
func NewCreateBuildTypeParamsWithContext(ctx context.Context) *CreateBuildTypeParams {
	return &CreateBuildTypeParams{
		Context: ctx,
	}
}

// NewCreateBuildTypeParamsWithHTTPClient creates a new CreateBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBuildTypeParamsWithHTTPClient(client *http.Client) *CreateBuildTypeParams {
	return &CreateBuildTypeParams{
		HTTPClient: client,
	}
}

/* CreateBuildTypeParams contains all the parameters to send to the API endpoint
   for the create build type operation.

   Typically these are written to a http.Request.
*/
type CreateBuildTypeParams struct {

	// Body.
	Body *models.BuildType

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBuildTypeParams) WithDefaults() *CreateBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create build type params
func (o *CreateBuildTypeParams) WithTimeout(timeout time.Duration) *CreateBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create build type params
func (o *CreateBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create build type params
func (o *CreateBuildTypeParams) WithContext(ctx context.Context) *CreateBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create build type params
func (o *CreateBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create build type params
func (o *CreateBuildTypeParams) WithHTTPClient(client *http.Client) *CreateBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create build type params
func (o *CreateBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create build type params
func (o *CreateBuildTypeParams) WithBody(body *models.BuildType) *CreateBuildTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create build type params
func (o *CreateBuildTypeParams) SetBody(body *models.BuildType) {
	o.Body = body
}

// WithFields adds the fields to the create build type params
func (o *CreateBuildTypeParams) WithFields(fields *string) *CreateBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create build type params
func (o *CreateBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
