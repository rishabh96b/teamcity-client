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
	"github.com/go-openapi/swag"
)

// NewRemoveTemplateParams creates a new RemoveTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTemplateParams() *RemoveTemplateParams {
	return &RemoveTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTemplateParamsWithTimeout creates a new RemoveTemplateParams object
// with the ability to set a timeout on a request.
func NewRemoveTemplateParamsWithTimeout(timeout time.Duration) *RemoveTemplateParams {
	return &RemoveTemplateParams{
		timeout: timeout,
	}
}

// NewRemoveTemplateParamsWithContext creates a new RemoveTemplateParams object
// with the ability to set a context for a request.
func NewRemoveTemplateParamsWithContext(ctx context.Context) *RemoveTemplateParams {
	return &RemoveTemplateParams{
		Context: ctx,
	}
}

// NewRemoveTemplateParamsWithHTTPClient creates a new RemoveTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTemplateParamsWithHTTPClient(client *http.Client) *RemoveTemplateParams {
	return &RemoveTemplateParams{
		HTTPClient: client,
	}
}

/* RemoveTemplateParams contains all the parameters to send to the API endpoint
   for the remove template operation.

   Typically these are written to a http.Request.
*/
type RemoveTemplateParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// InlineSettings.
	InlineSettings *bool

	// TemplateLocator.
	//
	// Format: BuildTypeLocator
	TemplateLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTemplateParams) WithDefaults() *RemoveTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove template params
func (o *RemoveTemplateParams) WithTimeout(timeout time.Duration) *RemoveTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove template params
func (o *RemoveTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove template params
func (o *RemoveTemplateParams) WithContext(ctx context.Context) *RemoveTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove template params
func (o *RemoveTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove template params
func (o *RemoveTemplateParams) WithHTTPClient(client *http.Client) *RemoveTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove template params
func (o *RemoveTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the remove template params
func (o *RemoveTemplateParams) WithBtLocator(btLocator string) *RemoveTemplateParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the remove template params
func (o *RemoveTemplateParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithInlineSettings adds the inlineSettings to the remove template params
func (o *RemoveTemplateParams) WithInlineSettings(inlineSettings *bool) *RemoveTemplateParams {
	o.SetInlineSettings(inlineSettings)
	return o
}

// SetInlineSettings adds the inlineSettings to the remove template params
func (o *RemoveTemplateParams) SetInlineSettings(inlineSettings *bool) {
	o.InlineSettings = inlineSettings
}

// WithTemplateLocator adds the templateLocator to the remove template params
func (o *RemoveTemplateParams) WithTemplateLocator(templateLocator string) *RemoveTemplateParams {
	o.SetTemplateLocator(templateLocator)
	return o
}

// SetTemplateLocator adds the templateLocator to the remove template params
func (o *RemoveTemplateParams) SetTemplateLocator(templateLocator string) {
	o.TemplateLocator = templateLocator
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if o.InlineSettings != nil {

		// query param inlineSettings
		var qrInlineSettings bool

		if o.InlineSettings != nil {
			qrInlineSettings = *o.InlineSettings
		}
		qInlineSettings := swag.FormatBool(qrInlineSettings)
		if qInlineSettings != "" {

			if err := r.SetQueryParam("inlineSettings", qInlineSettings); err != nil {
				return err
			}
		}
	}

	// path param templateLocator
	if err := r.SetPathParam("templateLocator", o.TemplateLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
