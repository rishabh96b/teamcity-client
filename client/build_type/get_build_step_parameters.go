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
)

// NewGetBuildStepParams creates a new GetBuildStepParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBuildStepParams() *GetBuildStepParams {
	return &GetBuildStepParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildStepParamsWithTimeout creates a new GetBuildStepParams object
// with the ability to set a timeout on a request.
func NewGetBuildStepParamsWithTimeout(timeout time.Duration) *GetBuildStepParams {
	return &GetBuildStepParams{
		timeout: timeout,
	}
}

// NewGetBuildStepParamsWithContext creates a new GetBuildStepParams object
// with the ability to set a context for a request.
func NewGetBuildStepParamsWithContext(ctx context.Context) *GetBuildStepParams {
	return &GetBuildStepParams{
		Context: ctx,
	}
}

// NewGetBuildStepParamsWithHTTPClient creates a new GetBuildStepParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBuildStepParamsWithHTTPClient(client *http.Client) *GetBuildStepParams {
	return &GetBuildStepParams{
		HTTPClient: client,
	}
}

/* GetBuildStepParams contains all the parameters to send to the API endpoint
   for the get build step operation.

   Typically these are written to a http.Request.
*/
type GetBuildStepParams struct {

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

// WithDefaults hydrates default values in the get build step params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildStepParams) WithDefaults() *GetBuildStepParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get build step params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBuildStepParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get build step params
func (o *GetBuildStepParams) WithTimeout(timeout time.Duration) *GetBuildStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build step params
func (o *GetBuildStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build step params
func (o *GetBuildStepParams) WithContext(ctx context.Context) *GetBuildStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build step params
func (o *GetBuildStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build step params
func (o *GetBuildStepParams) WithHTTPClient(client *http.Client) *GetBuildStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build step params
func (o *GetBuildStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get build step params
func (o *GetBuildStepParams) WithBtLocator(btLocator string) *GetBuildStepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get build step params
func (o *GetBuildStepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get build step params
func (o *GetBuildStepParams) WithFields(fields *string) *GetBuildStepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get build step params
func (o *GetBuildStepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStepID adds the stepID to the get build step params
func (o *GetBuildStepParams) WithStepID(stepID string) *GetBuildStepParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get build step params
func (o *GetBuildStepParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
