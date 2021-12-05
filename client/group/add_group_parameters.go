// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewAddGroupParams creates a new AddGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddGroupParams() *AddGroupParams {
	return &AddGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddGroupParamsWithTimeout creates a new AddGroupParams object
// with the ability to set a timeout on a request.
func NewAddGroupParamsWithTimeout(timeout time.Duration) *AddGroupParams {
	return &AddGroupParams{
		timeout: timeout,
	}
}

// NewAddGroupParamsWithContext creates a new AddGroupParams object
// with the ability to set a context for a request.
func NewAddGroupParamsWithContext(ctx context.Context) *AddGroupParams {
	return &AddGroupParams{
		Context: ctx,
	}
}

// NewAddGroupParamsWithHTTPClient creates a new AddGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddGroupParamsWithHTTPClient(client *http.Client) *AddGroupParams {
	return &AddGroupParams{
		HTTPClient: client,
	}
}

/* AddGroupParams contains all the parameters to send to the API endpoint
   for the add group operation.

   Typically these are written to a http.Request.
*/
type AddGroupParams struct {

	// Body.
	Body *models.Group

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddGroupParams) WithDefaults() *AddGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add group params
func (o *AddGroupParams) WithTimeout(timeout time.Duration) *AddGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add group params
func (o *AddGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add group params
func (o *AddGroupParams) WithContext(ctx context.Context) *AddGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add group params
func (o *AddGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add group params
func (o *AddGroupParams) WithHTTPClient(client *http.Client) *AddGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add group params
func (o *AddGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add group params
func (o *AddGroupParams) WithBody(body *models.Group) *AddGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add group params
func (o *AddGroupParams) SetBody(body *models.Group) {
	o.Body = body
}

// WithFields adds the fields to the add group params
func (o *AddGroupParams) WithFields(fields *string) *AddGroupParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add group params
func (o *AddGroupParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
