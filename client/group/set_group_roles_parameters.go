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

// NewSetGroupRolesParams creates a new SetGroupRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetGroupRolesParams() *SetGroupRolesParams {
	return &SetGroupRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetGroupRolesParamsWithTimeout creates a new SetGroupRolesParams object
// with the ability to set a timeout on a request.
func NewSetGroupRolesParamsWithTimeout(timeout time.Duration) *SetGroupRolesParams {
	return &SetGroupRolesParams{
		timeout: timeout,
	}
}

// NewSetGroupRolesParamsWithContext creates a new SetGroupRolesParams object
// with the ability to set a context for a request.
func NewSetGroupRolesParamsWithContext(ctx context.Context) *SetGroupRolesParams {
	return &SetGroupRolesParams{
		Context: ctx,
	}
}

// NewSetGroupRolesParamsWithHTTPClient creates a new SetGroupRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetGroupRolesParamsWithHTTPClient(client *http.Client) *SetGroupRolesParams {
	return &SetGroupRolesParams{
		HTTPClient: client,
	}
}

/* SetGroupRolesParams contains all the parameters to send to the API endpoint
   for the set group roles operation.

   Typically these are written to a http.Request.
*/
type SetGroupRolesParams struct {

	// Body.
	Body *models.Roles

	// GroupLocator.
	GroupLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetGroupRolesParams) WithDefaults() *SetGroupRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetGroupRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set group roles params
func (o *SetGroupRolesParams) WithTimeout(timeout time.Duration) *SetGroupRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set group roles params
func (o *SetGroupRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set group roles params
func (o *SetGroupRolesParams) WithContext(ctx context.Context) *SetGroupRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set group roles params
func (o *SetGroupRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set group roles params
func (o *SetGroupRolesParams) WithHTTPClient(client *http.Client) *SetGroupRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set group roles params
func (o *SetGroupRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set group roles params
func (o *SetGroupRolesParams) WithBody(body *models.Roles) *SetGroupRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set group roles params
func (o *SetGroupRolesParams) SetBody(body *models.Roles) {
	o.Body = body
}

// WithGroupLocator adds the groupLocator to the set group roles params
func (o *SetGroupRolesParams) WithGroupLocator(groupLocator string) *SetGroupRolesParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the set group roles params
func (o *SetGroupRolesParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetGroupRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupLocator
	if err := r.SetPathParam("groupLocator", o.GroupLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}