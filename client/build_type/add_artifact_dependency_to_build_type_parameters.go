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

// NewAddArtifactDependencyToBuildTypeParams creates a new AddArtifactDependencyToBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddArtifactDependencyToBuildTypeParams() *AddArtifactDependencyToBuildTypeParams {
	return &AddArtifactDependencyToBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddArtifactDependencyToBuildTypeParamsWithTimeout creates a new AddArtifactDependencyToBuildTypeParams object
// with the ability to set a timeout on a request.
func NewAddArtifactDependencyToBuildTypeParamsWithTimeout(timeout time.Duration) *AddArtifactDependencyToBuildTypeParams {
	return &AddArtifactDependencyToBuildTypeParams{
		timeout: timeout,
	}
}

// NewAddArtifactDependencyToBuildTypeParamsWithContext creates a new AddArtifactDependencyToBuildTypeParams object
// with the ability to set a context for a request.
func NewAddArtifactDependencyToBuildTypeParamsWithContext(ctx context.Context) *AddArtifactDependencyToBuildTypeParams {
	return &AddArtifactDependencyToBuildTypeParams{
		Context: ctx,
	}
}

// NewAddArtifactDependencyToBuildTypeParamsWithHTTPClient creates a new AddArtifactDependencyToBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddArtifactDependencyToBuildTypeParamsWithHTTPClient(client *http.Client) *AddArtifactDependencyToBuildTypeParams {
	return &AddArtifactDependencyToBuildTypeParams{
		HTTPClient: client,
	}
}

/* AddArtifactDependencyToBuildTypeParams contains all the parameters to send to the API endpoint
   for the add artifact dependency to build type operation.

   Typically these are written to a http.Request.
*/
type AddArtifactDependencyToBuildTypeParams struct {

	// Body.
	Body *models.ArtifactDependency

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add artifact dependency to build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddArtifactDependencyToBuildTypeParams) WithDefaults() *AddArtifactDependencyToBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add artifact dependency to build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddArtifactDependencyToBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithTimeout(timeout time.Duration) *AddArtifactDependencyToBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithContext(ctx context.Context) *AddArtifactDependencyToBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithHTTPClient(client *http.Client) *AddArtifactDependencyToBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithBody(body *models.ArtifactDependency) *AddArtifactDependencyToBuildTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetBody(body *models.ArtifactDependency) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithBtLocator(btLocator string) *AddArtifactDependencyToBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) WithFields(fields *string) *AddArtifactDependencyToBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add artifact dependency to build type params
func (o *AddArtifactDependencyToBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddArtifactDependencyToBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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