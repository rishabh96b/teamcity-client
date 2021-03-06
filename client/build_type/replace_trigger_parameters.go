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

	"github.com/rishabh96b/teamcity-client/models"
)

// NewReplaceTriggerParams creates a new ReplaceTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceTriggerParams() *ReplaceTriggerParams {
	return &ReplaceTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceTriggerParamsWithTimeout creates a new ReplaceTriggerParams object
// with the ability to set a timeout on a request.
func NewReplaceTriggerParamsWithTimeout(timeout time.Duration) *ReplaceTriggerParams {
	return &ReplaceTriggerParams{
		timeout: timeout,
	}
}

// NewReplaceTriggerParamsWithContext creates a new ReplaceTriggerParams object
// with the ability to set a context for a request.
func NewReplaceTriggerParamsWithContext(ctx context.Context) *ReplaceTriggerParams {
	return &ReplaceTriggerParams{
		Context: ctx,
	}
}

// NewReplaceTriggerParamsWithHTTPClient creates a new ReplaceTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceTriggerParamsWithHTTPClient(client *http.Client) *ReplaceTriggerParams {
	return &ReplaceTriggerParams{
		HTTPClient: client,
	}
}

/* ReplaceTriggerParams contains all the parameters to send to the API endpoint
   for the replace trigger operation.

   Typically these are written to a http.Request.
*/
type ReplaceTriggerParams struct {

	// Body.
	Body *models.Trigger

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	// TriggerLocator.
	TriggerLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceTriggerParams) WithDefaults() *ReplaceTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace trigger params
func (o *ReplaceTriggerParams) WithTimeout(timeout time.Duration) *ReplaceTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace trigger params
func (o *ReplaceTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace trigger params
func (o *ReplaceTriggerParams) WithContext(ctx context.Context) *ReplaceTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace trigger params
func (o *ReplaceTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace trigger params
func (o *ReplaceTriggerParams) WithHTTPClient(client *http.Client) *ReplaceTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace trigger params
func (o *ReplaceTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace trigger params
func (o *ReplaceTriggerParams) WithBody(body *models.Trigger) *ReplaceTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace trigger params
func (o *ReplaceTriggerParams) SetBody(body *models.Trigger) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace trigger params
func (o *ReplaceTriggerParams) WithBtLocator(btLocator string) *ReplaceTriggerParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace trigger params
func (o *ReplaceTriggerParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace trigger params
func (o *ReplaceTriggerParams) WithFields(fields *string) *ReplaceTriggerParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace trigger params
func (o *ReplaceTriggerParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithTriggerLocator adds the triggerLocator to the replace trigger params
func (o *ReplaceTriggerParams) WithTriggerLocator(triggerLocator string) *ReplaceTriggerParams {
	o.SetTriggerLocator(triggerLocator)
	return o
}

// SetTriggerLocator adds the triggerLocator to the replace trigger params
func (o *ReplaceTriggerParams) SetTriggerLocator(triggerLocator string) {
	o.TriggerLocator = triggerLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param triggerLocator
	if err := r.SetPathParam("triggerLocator", o.TriggerLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
