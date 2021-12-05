// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetAllBranchesParams creates a new GetAllBranchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllBranchesParams() *GetAllBranchesParams {
	return &GetAllBranchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllBranchesParamsWithTimeout creates a new GetAllBranchesParams object
// with the ability to set a timeout on a request.
func NewGetAllBranchesParamsWithTimeout(timeout time.Duration) *GetAllBranchesParams {
	return &GetAllBranchesParams{
		timeout: timeout,
	}
}

// NewGetAllBranchesParamsWithContext creates a new GetAllBranchesParams object
// with the ability to set a context for a request.
func NewGetAllBranchesParamsWithContext(ctx context.Context) *GetAllBranchesParams {
	return &GetAllBranchesParams{
		Context: ctx,
	}
}

// NewGetAllBranchesParamsWithHTTPClient creates a new GetAllBranchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllBranchesParamsWithHTTPClient(client *http.Client) *GetAllBranchesParams {
	return &GetAllBranchesParams{
		HTTPClient: client,
	}
}

/* GetAllBranchesParams contains all the parameters to send to the API endpoint
   for the get all branches operation.

   Typically these are written to a http.Request.
*/
type GetAllBranchesParams struct {

	// Fields.
	Fields *string

	// Locator.
	Locator *string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBranchesParams) WithDefaults() *GetAllBranchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllBranchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all branches params
func (o *GetAllBranchesParams) WithTimeout(timeout time.Duration) *GetAllBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all branches params
func (o *GetAllBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all branches params
func (o *GetAllBranchesParams) WithContext(ctx context.Context) *GetAllBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all branches params
func (o *GetAllBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all branches params
func (o *GetAllBranchesParams) WithHTTPClient(client *http.Client) *GetAllBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all branches params
func (o *GetAllBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get all branches params
func (o *GetAllBranchesParams) WithFields(fields *string) *GetAllBranchesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all branches params
func (o *GetAllBranchesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get all branches params
func (o *GetAllBranchesParams) WithLocator(locator *string) *GetAllBranchesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get all branches params
func (o *GetAllBranchesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithProjectLocator adds the projectLocator to the get all branches params
func (o *GetAllBranchesParams) WithProjectLocator(projectLocator string) *GetAllBranchesParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get all branches params
func (o *GetAllBranchesParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
