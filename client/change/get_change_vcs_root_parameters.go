// Code generated by go-swagger; DO NOT EDIT.

package change

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

// NewGetChangeVcsRootParams creates a new GetChangeVcsRootParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChangeVcsRootParams() *GetChangeVcsRootParams {
	return &GetChangeVcsRootParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChangeVcsRootParamsWithTimeout creates a new GetChangeVcsRootParams object
// with the ability to set a timeout on a request.
func NewGetChangeVcsRootParamsWithTimeout(timeout time.Duration) *GetChangeVcsRootParams {
	return &GetChangeVcsRootParams{
		timeout: timeout,
	}
}

// NewGetChangeVcsRootParamsWithContext creates a new GetChangeVcsRootParams object
// with the ability to set a context for a request.
func NewGetChangeVcsRootParamsWithContext(ctx context.Context) *GetChangeVcsRootParams {
	return &GetChangeVcsRootParams{
		Context: ctx,
	}
}

// NewGetChangeVcsRootParamsWithHTTPClient creates a new GetChangeVcsRootParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChangeVcsRootParamsWithHTTPClient(client *http.Client) *GetChangeVcsRootParams {
	return &GetChangeVcsRootParams{
		HTTPClient: client,
	}
}

/* GetChangeVcsRootParams contains all the parameters to send to the API endpoint
   for the get change vcs root operation.

   Typically these are written to a http.Request.
*/
type GetChangeVcsRootParams struct {

	// ChangeLocator.
	//
	// Format: ChangeLocator
	ChangeLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get change vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChangeVcsRootParams) WithDefaults() *GetChangeVcsRootParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get change vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChangeVcsRootParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get change vcs root params
func (o *GetChangeVcsRootParams) WithTimeout(timeout time.Duration) *GetChangeVcsRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get change vcs root params
func (o *GetChangeVcsRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get change vcs root params
func (o *GetChangeVcsRootParams) WithContext(ctx context.Context) *GetChangeVcsRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get change vcs root params
func (o *GetChangeVcsRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get change vcs root params
func (o *GetChangeVcsRootParams) WithHTTPClient(client *http.Client) *GetChangeVcsRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get change vcs root params
func (o *GetChangeVcsRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeLocator adds the changeLocator to the get change vcs root params
func (o *GetChangeVcsRootParams) WithChangeLocator(changeLocator string) *GetChangeVcsRootParams {
	o.SetChangeLocator(changeLocator)
	return o
}

// SetChangeLocator adds the changeLocator to the get change vcs root params
func (o *GetChangeVcsRootParams) SetChangeLocator(changeLocator string) {
	o.ChangeLocator = changeLocator
}

// WithFields adds the fields to the get change vcs root params
func (o *GetChangeVcsRootParams) WithFields(fields *string) *GetChangeVcsRootParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get change vcs root params
func (o *GetChangeVcsRootParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetChangeVcsRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param changeLocator
	if err := r.SetPathParam("changeLocator", o.ChangeLocator); err != nil {
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
