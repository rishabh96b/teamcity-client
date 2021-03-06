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

// NewGetAllBuildTypesParams creates a new GetAllBuildTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllBuildTypesParams() *GetAllBuildTypesParams {
	return &GetAllBuildTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllBuildTypesParamsWithTimeout creates a new GetAllBuildTypesParams object
// with the ability to set a timeout on a request.
func NewGetAllBuildTypesParamsWithTimeout(timeout time.Duration) *GetAllBuildTypesParams {
	return &GetAllBuildTypesParams{
		timeout: timeout,
	}
}

// NewGetAllBuildTypesParamsWithContext creates a new GetAllBuildTypesParams object
// with the ability to set a context for a request.
func NewGetAllBuildTypesParamsWithContext(ctx context.Context) *GetAllBuildTypesParams {
	return &GetAllBuildTypesParams{
		Context: ctx,
	}
}

// NewGetAllBuildTypesParamsWithHTTPClient creates a new GetAllBuildTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllBuildTypesParamsWithHTTPClient(client *http.Client) *GetAllBuildTypesParams {
	return &GetAllBuildTypesParams{
		HTTPClient: client,
	}
}

/* GetAllBuildTypesParams contains all the parameters to send to the API endpoint
   for the get all build types operation.

   Typically these are written to a http.Request.
*/
type GetAllBuildTypesParams struct {

	// Fields.
	Fields *string

	// Locator.
	//
	// Format: BuildTypeLocator
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all build types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBuildTypesParams) WithDefaults() *GetAllBuildTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all build types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBuildTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all build types params
func (o *GetAllBuildTypesParams) WithTimeout(timeout time.Duration) *GetAllBuildTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all build types params
func (o *GetAllBuildTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all build types params
func (o *GetAllBuildTypesParams) WithContext(ctx context.Context) *GetAllBuildTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all build types params
func (o *GetAllBuildTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all build types params
func (o *GetAllBuildTypesParams) WithHTTPClient(client *http.Client) *GetAllBuildTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all build types params
func (o *GetAllBuildTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get all build types params
func (o *GetAllBuildTypesParams) WithFields(fields *string) *GetAllBuildTypesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all build types params
func (o *GetAllBuildTypesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get all build types params
func (o *GetAllBuildTypesParams) WithLocator(locator *string) *GetAllBuildTypesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get all build types params
func (o *GetAllBuildTypesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllBuildTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
