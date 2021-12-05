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

// NewDeleteBuildStepParametersParams creates a new DeleteBuildStepParametersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildStepParametersParams() *DeleteBuildStepParametersParams {
	return &DeleteBuildStepParametersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildStepParametersParamsWithTimeout creates a new DeleteBuildStepParametersParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildStepParametersParamsWithTimeout(timeout time.Duration) *DeleteBuildStepParametersParams {
	return &DeleteBuildStepParametersParams{
		timeout: timeout,
	}
}

// NewDeleteBuildStepParametersParamsWithContext creates a new DeleteBuildStepParametersParams object
// with the ability to set a context for a request.
func NewDeleteBuildStepParametersParamsWithContext(ctx context.Context) *DeleteBuildStepParametersParams {
	return &DeleteBuildStepParametersParams{
		Context: ctx,
	}
}

// NewDeleteBuildStepParametersParamsWithHTTPClient creates a new DeleteBuildStepParametersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildStepParametersParamsWithHTTPClient(client *http.Client) *DeleteBuildStepParametersParams {
	return &DeleteBuildStepParametersParams{
		HTTPClient: client,
	}
}

/* DeleteBuildStepParametersParams contains all the parameters to send to the API endpoint
   for the delete build step parameters operation.

   Typically these are written to a http.Request.
*/
type DeleteBuildStepParametersParams struct {

	// Body.
	Body *models.Properties

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	// StepID.
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete build step parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildStepParametersParams) WithDefaults() *DeleteBuildStepParametersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build step parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildStepParametersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithTimeout(timeout time.Duration) *DeleteBuildStepParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithContext(ctx context.Context) *DeleteBuildStepParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithHTTPClient(client *http.Client) *DeleteBuildStepParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithBody(body *models.Properties) *DeleteBuildStepParametersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetBody(body *models.Properties) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithBtLocator(btLocator string) *DeleteBuildStepParametersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithFields(fields *string) *DeleteBuildStepParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStepID adds the stepID to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) WithStepID(stepID string) *DeleteBuildStepParametersParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the delete build step parameters params
func (o *DeleteBuildStepParametersParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildStepParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stepId
	if err := r.SetPathParam("stepId", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
