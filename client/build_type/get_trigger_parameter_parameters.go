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

// NewGetTriggerParameterParams creates a new GetTriggerParameterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTriggerParameterParams() *GetTriggerParameterParams {
	return &GetTriggerParameterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTriggerParameterParamsWithTimeout creates a new GetTriggerParameterParams object
// with the ability to set a timeout on a request.
func NewGetTriggerParameterParamsWithTimeout(timeout time.Duration) *GetTriggerParameterParams {
	return &GetTriggerParameterParams{
		timeout: timeout,
	}
}

// NewGetTriggerParameterParamsWithContext creates a new GetTriggerParameterParams object
// with the ability to set a context for a request.
func NewGetTriggerParameterParamsWithContext(ctx context.Context) *GetTriggerParameterParams {
	return &GetTriggerParameterParams{
		Context: ctx,
	}
}

// NewGetTriggerParameterParamsWithHTTPClient creates a new GetTriggerParameterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTriggerParameterParamsWithHTTPClient(client *http.Client) *GetTriggerParameterParams {
	return &GetTriggerParameterParams{
		HTTPClient: client,
	}
}

/* GetTriggerParameterParams contains all the parameters to send to the API endpoint
   for the get trigger parameter operation.

   Typically these are written to a http.Request.
*/
type GetTriggerParameterParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// FieldName.
	FieldName string

	// TriggerLocator.
	TriggerLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trigger parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTriggerParameterParams) WithDefaults() *GetTriggerParameterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trigger parameter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTriggerParameterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get trigger parameter params
func (o *GetTriggerParameterParams) WithTimeout(timeout time.Duration) *GetTriggerParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trigger parameter params
func (o *GetTriggerParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trigger parameter params
func (o *GetTriggerParameterParams) WithContext(ctx context.Context) *GetTriggerParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trigger parameter params
func (o *GetTriggerParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trigger parameter params
func (o *GetTriggerParameterParams) WithHTTPClient(client *http.Client) *GetTriggerParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trigger parameter params
func (o *GetTriggerParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get trigger parameter params
func (o *GetTriggerParameterParams) WithBtLocator(btLocator string) *GetTriggerParameterParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get trigger parameter params
func (o *GetTriggerParameterParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFieldName adds the fieldName to the get trigger parameter params
func (o *GetTriggerParameterParams) WithFieldName(fieldName string) *GetTriggerParameterParams {
	o.SetFieldName(fieldName)
	return o
}

// SetFieldName adds the fieldName to the get trigger parameter params
func (o *GetTriggerParameterParams) SetFieldName(fieldName string) {
	o.FieldName = fieldName
}

// WithTriggerLocator adds the triggerLocator to the get trigger parameter params
func (o *GetTriggerParameterParams) WithTriggerLocator(triggerLocator string) *GetTriggerParameterParams {
	o.SetTriggerLocator(triggerLocator)
	return o
}

// SetTriggerLocator adds the triggerLocator to the get trigger parameter params
func (o *GetTriggerParameterParams) SetTriggerLocator(triggerLocator string) {
	o.TriggerLocator = triggerLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetTriggerParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param fieldName
	if err := r.SetPathParam("fieldName", o.FieldName); err != nil {
		return err
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
