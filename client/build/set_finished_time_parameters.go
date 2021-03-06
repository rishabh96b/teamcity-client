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

// NewSetFinishedTimeParams creates a new SetFinishedTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetFinishedTimeParams() *SetFinishedTimeParams {
	return &SetFinishedTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetFinishedTimeParamsWithTimeout creates a new SetFinishedTimeParams object
// with the ability to set a timeout on a request.
func NewSetFinishedTimeParamsWithTimeout(timeout time.Duration) *SetFinishedTimeParams {
	return &SetFinishedTimeParams{
		timeout: timeout,
	}
}

// NewSetFinishedTimeParamsWithContext creates a new SetFinishedTimeParams object
// with the ability to set a context for a request.
func NewSetFinishedTimeParamsWithContext(ctx context.Context) *SetFinishedTimeParams {
	return &SetFinishedTimeParams{
		Context: ctx,
	}
}

// NewSetFinishedTimeParamsWithHTTPClient creates a new SetFinishedTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetFinishedTimeParamsWithHTTPClient(client *http.Client) *SetFinishedTimeParams {
	return &SetFinishedTimeParams{
		HTTPClient: client,
	}
}

/* SetFinishedTimeParams contains all the parameters to send to the API endpoint
   for the set finished time operation.

   Typically these are written to a http.Request.
*/
type SetFinishedTimeParams struct {

	// BuildLocator.
	//
	// Format: BuildLocator
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set finished time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetFinishedTimeParams) WithDefaults() *SetFinishedTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set finished time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetFinishedTimeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set finished time params
func (o *SetFinishedTimeParams) WithTimeout(timeout time.Duration) *SetFinishedTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set finished time params
func (o *SetFinishedTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set finished time params
func (o *SetFinishedTimeParams) WithContext(ctx context.Context) *SetFinishedTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set finished time params
func (o *SetFinishedTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set finished time params
func (o *SetFinishedTimeParams) WithHTTPClient(client *http.Client) *SetFinishedTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set finished time params
func (o *SetFinishedTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the set finished time params
func (o *SetFinishedTimeParams) WithBuildLocator(buildLocator string) *SetFinishedTimeParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the set finished time params
func (o *SetFinishedTimeParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetFinishedTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
