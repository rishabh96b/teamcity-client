// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

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

// NewTriggerCommitHookNotificationParams creates a new TriggerCommitHookNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerCommitHookNotificationParams() *TriggerCommitHookNotificationParams {
	return &TriggerCommitHookNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerCommitHookNotificationParamsWithTimeout creates a new TriggerCommitHookNotificationParams object
// with the ability to set a timeout on a request.
func NewTriggerCommitHookNotificationParamsWithTimeout(timeout time.Duration) *TriggerCommitHookNotificationParams {
	return &TriggerCommitHookNotificationParams{
		timeout: timeout,
	}
}

// NewTriggerCommitHookNotificationParamsWithContext creates a new TriggerCommitHookNotificationParams object
// with the ability to set a context for a request.
func NewTriggerCommitHookNotificationParamsWithContext(ctx context.Context) *TriggerCommitHookNotificationParams {
	return &TriggerCommitHookNotificationParams{
		Context: ctx,
	}
}

// NewTriggerCommitHookNotificationParamsWithHTTPClient creates a new TriggerCommitHookNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerCommitHookNotificationParamsWithHTTPClient(client *http.Client) *TriggerCommitHookNotificationParams {
	return &TriggerCommitHookNotificationParams{
		HTTPClient: client,
	}
}

/* TriggerCommitHookNotificationParams contains all the parameters to send to the API endpoint
   for the trigger commit hook notification operation.

   Typically these are written to a http.Request.
*/
type TriggerCommitHookNotificationParams struct {

	// Locator.
	//
	// Format: VcsRootInstanceLocator
	Locator *string

	// OkOnNothingFound.
	OkOnNothingFound *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger commit hook notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerCommitHookNotificationParams) WithDefaults() *TriggerCommitHookNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger commit hook notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerCommitHookNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) WithTimeout(timeout time.Duration) *TriggerCommitHookNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) WithContext(ctx context.Context) *TriggerCommitHookNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) WithHTTPClient(client *http.Client) *TriggerCommitHookNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocator adds the locator to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) WithLocator(locator *string) *TriggerCommitHookNotificationParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithOkOnNothingFound adds the okOnNothingFound to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) WithOkOnNothingFound(okOnNothingFound *bool) *TriggerCommitHookNotificationParams {
	o.SetOkOnNothingFound(okOnNothingFound)
	return o
}

// SetOkOnNothingFound adds the okOnNothingFound to the trigger commit hook notification params
func (o *TriggerCommitHookNotificationParams) SetOkOnNothingFound(okOnNothingFound *bool) {
	o.OkOnNothingFound = okOnNothingFound
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerCommitHookNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.OkOnNothingFound != nil {

		// query param okOnNothingFound
		var qrOkOnNothingFound bool

		if o.OkOnNothingFound != nil {
			qrOkOnNothingFound = *o.OkOnNothingFound
		}
		qOkOnNothingFound := swag.FormatBool(qrOkOnNothingFound)
		if qOkOnNothingFound != "" {

			if err := r.SetQueryParam("okOnNothingFound", qOkOnNothingFound); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
