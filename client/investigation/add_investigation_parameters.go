// Code generated by go-swagger; DO NOT EDIT.

package investigation

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

// NewAddInvestigationParams creates a new AddInvestigationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddInvestigationParams() *AddInvestigationParams {
	return &AddInvestigationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddInvestigationParamsWithTimeout creates a new AddInvestigationParams object
// with the ability to set a timeout on a request.
func NewAddInvestigationParamsWithTimeout(timeout time.Duration) *AddInvestigationParams {
	return &AddInvestigationParams{
		timeout: timeout,
	}
}

// NewAddInvestigationParamsWithContext creates a new AddInvestigationParams object
// with the ability to set a context for a request.
func NewAddInvestigationParamsWithContext(ctx context.Context) *AddInvestigationParams {
	return &AddInvestigationParams{
		Context: ctx,
	}
}

// NewAddInvestigationParamsWithHTTPClient creates a new AddInvestigationParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddInvestigationParamsWithHTTPClient(client *http.Client) *AddInvestigationParams {
	return &AddInvestigationParams{
		HTTPClient: client,
	}
}

/* AddInvestigationParams contains all the parameters to send to the API endpoint
   for the add investigation operation.

   Typically these are written to a http.Request.
*/
type AddInvestigationParams struct {

	// Body.
	Body *models.Investigation

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add investigation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddInvestigationParams) WithDefaults() *AddInvestigationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add investigation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddInvestigationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add investigation params
func (o *AddInvestigationParams) WithTimeout(timeout time.Duration) *AddInvestigationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add investigation params
func (o *AddInvestigationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add investigation params
func (o *AddInvestigationParams) WithContext(ctx context.Context) *AddInvestigationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add investigation params
func (o *AddInvestigationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add investigation params
func (o *AddInvestigationParams) WithHTTPClient(client *http.Client) *AddInvestigationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add investigation params
func (o *AddInvestigationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add investigation params
func (o *AddInvestigationParams) WithBody(body *models.Investigation) *AddInvestigationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add investigation params
func (o *AddInvestigationParams) SetBody(body *models.Investigation) {
	o.Body = body
}

// WithFields adds the fields to the add investigation params
func (o *AddInvestigationParams) WithFields(fields *string) *AddInvestigationParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add investigation params
func (o *AddInvestigationParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddInvestigationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
