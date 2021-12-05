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

// NewGetVcsRootInstancesOfBuildTypeParams creates a new GetVcsRootInstancesOfBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsRootInstancesOfBuildTypeParams() *GetVcsRootInstancesOfBuildTypeParams {
	return &GetVcsRootInstancesOfBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootInstancesOfBuildTypeParamsWithTimeout creates a new GetVcsRootInstancesOfBuildTypeParams object
// with the ability to set a timeout on a request.
func NewGetVcsRootInstancesOfBuildTypeParamsWithTimeout(timeout time.Duration) *GetVcsRootInstancesOfBuildTypeParams {
	return &GetVcsRootInstancesOfBuildTypeParams{
		timeout: timeout,
	}
}

// NewGetVcsRootInstancesOfBuildTypeParamsWithContext creates a new GetVcsRootInstancesOfBuildTypeParams object
// with the ability to set a context for a request.
func NewGetVcsRootInstancesOfBuildTypeParamsWithContext(ctx context.Context) *GetVcsRootInstancesOfBuildTypeParams {
	return &GetVcsRootInstancesOfBuildTypeParams{
		Context: ctx,
	}
}

// NewGetVcsRootInstancesOfBuildTypeParamsWithHTTPClient creates a new GetVcsRootInstancesOfBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsRootInstancesOfBuildTypeParamsWithHTTPClient(client *http.Client) *GetVcsRootInstancesOfBuildTypeParams {
	return &GetVcsRootInstancesOfBuildTypeParams{
		HTTPClient: client,
	}
}

/* GetVcsRootInstancesOfBuildTypeParams contains all the parameters to send to the API endpoint
   for the get vcs root instances of build type operation.

   Typically these are written to a http.Request.
*/
type GetVcsRootInstancesOfBuildTypeParams struct {

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

// WithDefaults hydrates default values in the get vcs root instances of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstancesOfBuildTypeParams) WithDefaults() *GetVcsRootInstancesOfBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs root instances of build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsRootInstancesOfBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) WithTimeout(timeout time.Duration) *GetVcsRootInstancesOfBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) WithContext(ctx context.Context) *GetVcsRootInstancesOfBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) WithHTTPClient(client *http.Client) *GetVcsRootInstancesOfBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) WithBtLocator(btLocator string) *GetVcsRootInstancesOfBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) WithFields(fields *string) *GetVcsRootInstancesOfBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root instances of build type params
func (o *GetVcsRootInstancesOfBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootInstancesOfBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
