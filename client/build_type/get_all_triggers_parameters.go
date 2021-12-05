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

// NewGetAllTriggersParams creates a new GetAllTriggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllTriggersParams() *GetAllTriggersParams {
	return &GetAllTriggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllTriggersParamsWithTimeout creates a new GetAllTriggersParams object
// with the ability to set a timeout on a request.
func NewGetAllTriggersParamsWithTimeout(timeout time.Duration) *GetAllTriggersParams {
	return &GetAllTriggersParams{
		timeout: timeout,
	}
}

// NewGetAllTriggersParamsWithContext creates a new GetAllTriggersParams object
// with the ability to set a context for a request.
func NewGetAllTriggersParamsWithContext(ctx context.Context) *GetAllTriggersParams {
	return &GetAllTriggersParams{
		Context: ctx,
	}
}

// NewGetAllTriggersParamsWithHTTPClient creates a new GetAllTriggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllTriggersParamsWithHTTPClient(client *http.Client) *GetAllTriggersParams {
	return &GetAllTriggersParams{
		HTTPClient: client,
	}
}

/* GetAllTriggersParams contains all the parameters to send to the API endpoint
   for the get all triggers operation.

   Typically these are written to a http.Request.
*/
type GetAllTriggersParams struct {

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

// WithDefaults hydrates default values in the get all triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllTriggersParams) WithDefaults() *GetAllTriggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllTriggersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all triggers params
func (o *GetAllTriggersParams) WithTimeout(timeout time.Duration) *GetAllTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all triggers params
func (o *GetAllTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all triggers params
func (o *GetAllTriggersParams) WithContext(ctx context.Context) *GetAllTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all triggers params
func (o *GetAllTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all triggers params
func (o *GetAllTriggersParams) WithHTTPClient(client *http.Client) *GetAllTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all triggers params
func (o *GetAllTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get all triggers params
func (o *GetAllTriggersParams) WithBtLocator(btLocator string) *GetAllTriggersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get all triggers params
func (o *GetAllTriggersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get all triggers params
func (o *GetAllTriggersParams) WithFields(fields *string) *GetAllTriggersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all triggers params
func (o *GetAllTriggersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
