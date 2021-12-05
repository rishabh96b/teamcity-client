// Code generated by go-swagger; DO NOT EDIT.

package mute

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

// NewMuteMultipleTestsParams creates a new MuteMultipleTestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMuteMultipleTestsParams() *MuteMultipleTestsParams {
	return &MuteMultipleTestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMuteMultipleTestsParamsWithTimeout creates a new MuteMultipleTestsParams object
// with the ability to set a timeout on a request.
func NewMuteMultipleTestsParamsWithTimeout(timeout time.Duration) *MuteMultipleTestsParams {
	return &MuteMultipleTestsParams{
		timeout: timeout,
	}
}

// NewMuteMultipleTestsParamsWithContext creates a new MuteMultipleTestsParams object
// with the ability to set a context for a request.
func NewMuteMultipleTestsParamsWithContext(ctx context.Context) *MuteMultipleTestsParams {
	return &MuteMultipleTestsParams{
		Context: ctx,
	}
}

// NewMuteMultipleTestsParamsWithHTTPClient creates a new MuteMultipleTestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewMuteMultipleTestsParamsWithHTTPClient(client *http.Client) *MuteMultipleTestsParams {
	return &MuteMultipleTestsParams{
		HTTPClient: client,
	}
}

/* MuteMultipleTestsParams contains all the parameters to send to the API endpoint
   for the mute multiple tests operation.

   Typically these are written to a http.Request.
*/
type MuteMultipleTestsParams struct {

	// Body.
	Body *models.Mutes

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mute multiple tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MuteMultipleTestsParams) WithDefaults() *MuteMultipleTestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mute multiple tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MuteMultipleTestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mute multiple tests params
func (o *MuteMultipleTestsParams) WithTimeout(timeout time.Duration) *MuteMultipleTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mute multiple tests params
func (o *MuteMultipleTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mute multiple tests params
func (o *MuteMultipleTestsParams) WithContext(ctx context.Context) *MuteMultipleTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mute multiple tests params
func (o *MuteMultipleTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mute multiple tests params
func (o *MuteMultipleTestsParams) WithHTTPClient(client *http.Client) *MuteMultipleTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mute multiple tests params
func (o *MuteMultipleTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the mute multiple tests params
func (o *MuteMultipleTestsParams) WithBody(body *models.Mutes) *MuteMultipleTestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the mute multiple tests params
func (o *MuteMultipleTestsParams) SetBody(body *models.Mutes) {
	o.Body = body
}

// WithFields adds the fields to the mute multiple tests params
func (o *MuteMultipleTestsParams) WithFields(fields *string) *MuteMultipleTestsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the mute multiple tests params
func (o *MuteMultipleTestsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *MuteMultipleTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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