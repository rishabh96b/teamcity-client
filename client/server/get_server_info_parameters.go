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

// NewGetServerInfoParams creates a new GetServerInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerInfoParams() *GetServerInfoParams {
	return &GetServerInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerInfoParamsWithTimeout creates a new GetServerInfoParams object
// with the ability to set a timeout on a request.
func NewGetServerInfoParamsWithTimeout(timeout time.Duration) *GetServerInfoParams {
	return &GetServerInfoParams{
		timeout: timeout,
	}
}

// NewGetServerInfoParamsWithContext creates a new GetServerInfoParams object
// with the ability to set a context for a request.
func NewGetServerInfoParamsWithContext(ctx context.Context) *GetServerInfoParams {
	return &GetServerInfoParams{
		Context: ctx,
	}
}

// NewGetServerInfoParamsWithHTTPClient creates a new GetServerInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerInfoParamsWithHTTPClient(client *http.Client) *GetServerInfoParams {
	return &GetServerInfoParams{
		HTTPClient: client,
	}
}

/* GetServerInfoParams contains all the parameters to send to the API endpoint
   for the get server info operation.

   Typically these are written to a http.Request.
*/
type GetServerInfoParams struct {

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerInfoParams) WithDefaults() *GetServerInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server info params
func (o *GetServerInfoParams) WithTimeout(timeout time.Duration) *GetServerInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server info params
func (o *GetServerInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server info params
func (o *GetServerInfoParams) WithContext(ctx context.Context) *GetServerInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server info params
func (o *GetServerInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server info params
func (o *GetServerInfoParams) WithHTTPClient(client *http.Client) *GetServerInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server info params
func (o *GetServerInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get server info params
func (o *GetServerInfoParams) WithFields(fields *string) *GetServerInfoParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get server info params
func (o *GetServerInfoParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
