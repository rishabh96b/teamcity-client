// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetAllMetricsParams creates a new GetAllMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllMetricsParams() *GetAllMetricsParams {
	return &GetAllMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMetricsParamsWithTimeout creates a new GetAllMetricsParams object
// with the ability to set a timeout on a request.
func NewGetAllMetricsParamsWithTimeout(timeout time.Duration) *GetAllMetricsParams {
	return &GetAllMetricsParams{
		timeout: timeout,
	}
}

// NewGetAllMetricsParamsWithContext creates a new GetAllMetricsParams object
// with the ability to set a context for a request.
func NewGetAllMetricsParamsWithContext(ctx context.Context) *GetAllMetricsParams {
	return &GetAllMetricsParams{
		Context: ctx,
	}
}

// NewGetAllMetricsParamsWithHTTPClient creates a new GetAllMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllMetricsParamsWithHTTPClient(client *http.Client) *GetAllMetricsParams {
	return &GetAllMetricsParams{
		HTTPClient: client,
	}
}

/* GetAllMetricsParams contains all the parameters to send to the API endpoint
   for the get all metrics operation.

   Typically these are written to a http.Request.
*/
type GetAllMetricsParams struct {

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMetricsParams) WithDefaults() *GetAllMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all metrics params
func (o *GetAllMetricsParams) WithTimeout(timeout time.Duration) *GetAllMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all metrics params
func (o *GetAllMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all metrics params
func (o *GetAllMetricsParams) WithContext(ctx context.Context) *GetAllMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all metrics params
func (o *GetAllMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all metrics params
func (o *GetAllMetricsParams) WithHTTPClient(client *http.Client) *GetAllMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all metrics params
func (o *GetAllMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get all metrics params
func (o *GetAllMetricsParams) WithFields(fields *string) *GetAllMetricsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all metrics params
func (o *GetAllMetricsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
