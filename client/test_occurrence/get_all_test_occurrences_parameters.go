// Code generated by go-swagger; DO NOT EDIT.

package test_occurrence

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

// NewGetAllTestOccurrencesParams creates a new GetAllTestOccurrencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllTestOccurrencesParams() *GetAllTestOccurrencesParams {
	return &GetAllTestOccurrencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllTestOccurrencesParamsWithTimeout creates a new GetAllTestOccurrencesParams object
// with the ability to set a timeout on a request.
func NewGetAllTestOccurrencesParamsWithTimeout(timeout time.Duration) *GetAllTestOccurrencesParams {
	return &GetAllTestOccurrencesParams{
		timeout: timeout,
	}
}

// NewGetAllTestOccurrencesParamsWithContext creates a new GetAllTestOccurrencesParams object
// with the ability to set a context for a request.
func NewGetAllTestOccurrencesParamsWithContext(ctx context.Context) *GetAllTestOccurrencesParams {
	return &GetAllTestOccurrencesParams{
		Context: ctx,
	}
}

// NewGetAllTestOccurrencesParamsWithHTTPClient creates a new GetAllTestOccurrencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllTestOccurrencesParamsWithHTTPClient(client *http.Client) *GetAllTestOccurrencesParams {
	return &GetAllTestOccurrencesParams{
		HTTPClient: client,
	}
}

/* GetAllTestOccurrencesParams contains all the parameters to send to the API endpoint
   for the get all test occurrences operation.

   Typically these are written to a http.Request.
*/
type GetAllTestOccurrencesParams struct {

	// Fields.
	Fields *string

	// Locator.
	//
	// Format: TestOccurrenceLocator
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all test occurrences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllTestOccurrencesParams) WithDefaults() *GetAllTestOccurrencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all test occurrences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllTestOccurrencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) WithTimeout(timeout time.Duration) *GetAllTestOccurrencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) WithContext(ctx context.Context) *GetAllTestOccurrencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) WithHTTPClient(client *http.Client) *GetAllTestOccurrencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) WithFields(fields *string) *GetAllTestOccurrencesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) WithLocator(locator *string) *GetAllTestOccurrencesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get all test occurrences params
func (o *GetAllTestOccurrencesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllTestOccurrencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
