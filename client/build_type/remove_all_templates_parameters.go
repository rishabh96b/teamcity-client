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

// NewRemoveAllTemplatesParams creates a new RemoveAllTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveAllTemplatesParams() *RemoveAllTemplatesParams {
	return &RemoveAllTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAllTemplatesParamsWithTimeout creates a new RemoveAllTemplatesParams object
// with the ability to set a timeout on a request.
func NewRemoveAllTemplatesParamsWithTimeout(timeout time.Duration) *RemoveAllTemplatesParams {
	return &RemoveAllTemplatesParams{
		timeout: timeout,
	}
}

// NewRemoveAllTemplatesParamsWithContext creates a new RemoveAllTemplatesParams object
// with the ability to set a context for a request.
func NewRemoveAllTemplatesParamsWithContext(ctx context.Context) *RemoveAllTemplatesParams {
	return &RemoveAllTemplatesParams{
		Context: ctx,
	}
}

// NewRemoveAllTemplatesParamsWithHTTPClient creates a new RemoveAllTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveAllTemplatesParamsWithHTTPClient(client *http.Client) *RemoveAllTemplatesParams {
	return &RemoveAllTemplatesParams{
		HTTPClient: client,
	}
}

/* RemoveAllTemplatesParams contains all the parameters to send to the API endpoint
   for the remove all templates operation.

   Typically these are written to a http.Request.
*/
type RemoveAllTemplatesParams struct {

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// InlineSettings.
	InlineSettings *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove all templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAllTemplatesParams) WithDefaults() *RemoveAllTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove all templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAllTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove all templates params
func (o *RemoveAllTemplatesParams) WithTimeout(timeout time.Duration) *RemoveAllTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove all templates params
func (o *RemoveAllTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove all templates params
func (o *RemoveAllTemplatesParams) WithContext(ctx context.Context) *RemoveAllTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove all templates params
func (o *RemoveAllTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove all templates params
func (o *RemoveAllTemplatesParams) WithHTTPClient(client *http.Client) *RemoveAllTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove all templates params
func (o *RemoveAllTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the remove all templates params
func (o *RemoveAllTemplatesParams) WithBtLocator(btLocator string) *RemoveAllTemplatesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the remove all templates params
func (o *RemoveAllTemplatesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithInlineSettings adds the inlineSettings to the remove all templates params
func (o *RemoveAllTemplatesParams) WithInlineSettings(inlineSettings *bool) *RemoveAllTemplatesParams {
	o.SetInlineSettings(inlineSettings)
	return o
}

// SetInlineSettings adds the inlineSettings to the remove all templates params
func (o *RemoveAllTemplatesParams) SetInlineSettings(inlineSettings *bool) {
	o.InlineSettings = inlineSettings
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAllTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
