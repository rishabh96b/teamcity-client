// Code generated by go-swagger; DO NOT EDIT.

package build_queue

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

// NewGetQueuedBuildParams creates a new GetQueuedBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueuedBuildParams() *GetQueuedBuildParams {
	return &GetQueuedBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueuedBuildParamsWithTimeout creates a new GetQueuedBuildParams object
// with the ability to set a timeout on a request.
func NewGetQueuedBuildParamsWithTimeout(timeout time.Duration) *GetQueuedBuildParams {
	return &GetQueuedBuildParams{
		timeout: timeout,
	}
}

// NewGetQueuedBuildParamsWithContext creates a new GetQueuedBuildParams object
// with the ability to set a context for a request.
func NewGetQueuedBuildParamsWithContext(ctx context.Context) *GetQueuedBuildParams {
	return &GetQueuedBuildParams{
		Context: ctx,
	}
}

// NewGetQueuedBuildParamsWithHTTPClient creates a new GetQueuedBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueuedBuildParamsWithHTTPClient(client *http.Client) *GetQueuedBuildParams {
	return &GetQueuedBuildParams{
		HTTPClient: client,
	}
}

/* GetQueuedBuildParams contains all the parameters to send to the API endpoint
   for the get queued build operation.

   Typically these are written to a http.Request.
*/
type GetQueuedBuildParams struct {

	// Fields.
	Fields *string

	// QueuedBuildLocator.
	//
	// Format: BuildQueueLocator
	QueuedBuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get queued build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueuedBuildParams) WithDefaults() *GetQueuedBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get queued build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueuedBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get queued build params
func (o *GetQueuedBuildParams) WithTimeout(timeout time.Duration) *GetQueuedBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queued build params
func (o *GetQueuedBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queued build params
func (o *GetQueuedBuildParams) WithContext(ctx context.Context) *GetQueuedBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queued build params
func (o *GetQueuedBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queued build params
func (o *GetQueuedBuildParams) WithHTTPClient(client *http.Client) *GetQueuedBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queued build params
func (o *GetQueuedBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get queued build params
func (o *GetQueuedBuildParams) WithFields(fields *string) *GetQueuedBuildParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get queued build params
func (o *GetQueuedBuildParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithQueuedBuildLocator adds the queuedBuildLocator to the get queued build params
func (o *GetQueuedBuildParams) WithQueuedBuildLocator(queuedBuildLocator string) *GetQueuedBuildParams {
	o.SetQueuedBuildLocator(queuedBuildLocator)
	return o
}

// SetQueuedBuildLocator adds the queuedBuildLocator to the get queued build params
func (o *GetQueuedBuildParams) SetQueuedBuildLocator(queuedBuildLocator string) {
	o.QueuedBuildLocator = queuedBuildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueuedBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queuedBuildLocator
	if err := r.SetPathParam("queuedBuildLocator", o.QueuedBuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
