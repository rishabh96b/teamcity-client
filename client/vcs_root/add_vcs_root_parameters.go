// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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

	"github.com/rishabh96b/teamcity-client/models"
)

// NewAddVcsRootParams creates a new AddVcsRootParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVcsRootParams() *AddVcsRootParams {
	return &AddVcsRootParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVcsRootParamsWithTimeout creates a new AddVcsRootParams object
// with the ability to set a timeout on a request.
func NewAddVcsRootParamsWithTimeout(timeout time.Duration) *AddVcsRootParams {
	return &AddVcsRootParams{
		timeout: timeout,
	}
}

// NewAddVcsRootParamsWithContext creates a new AddVcsRootParams object
// with the ability to set a context for a request.
func NewAddVcsRootParamsWithContext(ctx context.Context) *AddVcsRootParams {
	return &AddVcsRootParams{
		Context: ctx,
	}
}

// NewAddVcsRootParamsWithHTTPClient creates a new AddVcsRootParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVcsRootParamsWithHTTPClient(client *http.Client) *AddVcsRootParams {
	return &AddVcsRootParams{
		HTTPClient: client,
	}
}

/* AddVcsRootParams contains all the parameters to send to the API endpoint
   for the add vcs root operation.

   Typically these are written to a http.Request.
*/
type AddVcsRootParams struct {

	// Body.
	Body *models.VcsRoot

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVcsRootParams) WithDefaults() *AddVcsRootParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add vcs root params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVcsRootParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add vcs root params
func (o *AddVcsRootParams) WithTimeout(timeout time.Duration) *AddVcsRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add vcs root params
func (o *AddVcsRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add vcs root params
func (o *AddVcsRootParams) WithContext(ctx context.Context) *AddVcsRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add vcs root params
func (o *AddVcsRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add vcs root params
func (o *AddVcsRootParams) WithHTTPClient(client *http.Client) *AddVcsRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add vcs root params
func (o *AddVcsRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add vcs root params
func (o *AddVcsRootParams) WithBody(body *models.VcsRoot) *AddVcsRootParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add vcs root params
func (o *AddVcsRootParams) SetBody(body *models.VcsRoot) {
	o.Body = body
}

// WithFields adds the fields to the add vcs root params
func (o *AddVcsRootParams) WithFields(fields *string) *AddVcsRootParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add vcs root params
func (o *AddVcsRootParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddVcsRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
