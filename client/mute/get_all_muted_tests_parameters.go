// Code generated by go-swagger; DO NOT EDIT.

package mute

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

// NewGetAllMutedTestsParams creates a new GetAllMutedTestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllMutedTestsParams() *GetAllMutedTestsParams {
	return &GetAllMutedTestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMutedTestsParamsWithTimeout creates a new GetAllMutedTestsParams object
// with the ability to set a timeout on a request.
func NewGetAllMutedTestsParamsWithTimeout(timeout time.Duration) *GetAllMutedTestsParams {
	return &GetAllMutedTestsParams{
		timeout: timeout,
	}
}

// NewGetAllMutedTestsParamsWithContext creates a new GetAllMutedTestsParams object
// with the ability to set a context for a request.
func NewGetAllMutedTestsParamsWithContext(ctx context.Context) *GetAllMutedTestsParams {
	return &GetAllMutedTestsParams{
		Context: ctx,
	}
}

// NewGetAllMutedTestsParamsWithHTTPClient creates a new GetAllMutedTestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllMutedTestsParamsWithHTTPClient(client *http.Client) *GetAllMutedTestsParams {
	return &GetAllMutedTestsParams{
		HTTPClient: client,
	}
}

/* GetAllMutedTestsParams contains all the parameters to send to the API endpoint
   for the get all muted tests operation.

   Typically these are written to a http.Request.
*/
type GetAllMutedTestsParams struct {

	// Fields.
	Fields *string

	// Locator.
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all muted tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMutedTestsParams) WithDefaults() *GetAllMutedTestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all muted tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMutedTestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all muted tests params
func (o *GetAllMutedTestsParams) WithTimeout(timeout time.Duration) *GetAllMutedTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all muted tests params
func (o *GetAllMutedTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all muted tests params
func (o *GetAllMutedTestsParams) WithContext(ctx context.Context) *GetAllMutedTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all muted tests params
func (o *GetAllMutedTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all muted tests params
func (o *GetAllMutedTestsParams) WithHTTPClient(client *http.Client) *GetAllMutedTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all muted tests params
func (o *GetAllMutedTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get all muted tests params
func (o *GetAllMutedTestsParams) WithFields(fields *string) *GetAllMutedTestsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all muted tests params
func (o *GetAllMutedTestsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get all muted tests params
func (o *GetAllMutedTestsParams) WithLocator(locator *string) *GetAllMutedTestsParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get all muted tests params
func (o *GetAllMutedTestsParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMutedTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string

		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {

			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
