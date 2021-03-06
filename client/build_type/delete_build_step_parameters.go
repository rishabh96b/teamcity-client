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

// NewDeleteBuildStepParams creates a new DeleteBuildStepParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildStepParams() *DeleteBuildStepParams {
	return &DeleteBuildStepParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildStepParamsWithTimeout creates a new DeleteBuildStepParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildStepParamsWithTimeout(timeout time.Duration) *DeleteBuildStepParams {
	return &DeleteBuildStepParams{
		timeout: timeout,
	}
}

// NewDeleteBuildStepParamsWithContext creates a new DeleteBuildStepParams object
// with the ability to set a context for a request.
func NewDeleteBuildStepParamsWithContext(ctx context.Context) *DeleteBuildStepParams {
	return &DeleteBuildStepParams{
		Context: ctx,
	}
}

// NewDeleteBuildStepParamsWithHTTPClient creates a new DeleteBuildStepParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildStepParamsWithHTTPClient(client *http.Client) *DeleteBuildStepParams {
	return &DeleteBuildStepParams{
		HTTPClient: client,
	}
}

/* DeleteBuildStepParams contains all the parameters to send to the API endpoint
   for the delete build step operation.

   Typically these are written to a http.Request.
*/
type DeleteBuildStepParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// StepID.
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete build step params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildStepParams) WithDefaults() *DeleteBuildStepParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build step params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildStepParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build step params
func (o *DeleteBuildStepParams) WithTimeout(timeout time.Duration) *DeleteBuildStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build step params
func (o *DeleteBuildStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build step params
func (o *DeleteBuildStepParams) WithContext(ctx context.Context) *DeleteBuildStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build step params
func (o *DeleteBuildStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build step params
func (o *DeleteBuildStepParams) WithHTTPClient(client *http.Client) *DeleteBuildStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build step params
func (o *DeleteBuildStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the delete build step params
func (o *DeleteBuildStepParams) WithBtLocator(btLocator string) *DeleteBuildStepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete build step params
func (o *DeleteBuildStepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithStepID adds the stepID to the delete build step params
func (o *DeleteBuildStepParams) WithStepID(stepID string) *DeleteBuildStepParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the delete build step params
func (o *DeleteBuildStepParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
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
