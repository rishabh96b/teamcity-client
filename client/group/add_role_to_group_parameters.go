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

// NewAddRoleToGroupParams creates a new AddRoleToGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRoleToGroupParams() *AddRoleToGroupParams {
	return &AddRoleToGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRoleToGroupParamsWithTimeout creates a new AddRoleToGroupParams object
// with the ability to set a timeout on a request.
func NewAddRoleToGroupParamsWithTimeout(timeout time.Duration) *AddRoleToGroupParams {
	return &AddRoleToGroupParams{
		timeout: timeout,
	}
}

// NewAddRoleToGroupParamsWithContext creates a new AddRoleToGroupParams object
// with the ability to set a context for a request.
func NewAddRoleToGroupParamsWithContext(ctx context.Context) *AddRoleToGroupParams {
	return &AddRoleToGroupParams{
		Context: ctx,
	}
}

// NewAddRoleToGroupParamsWithHTTPClient creates a new AddRoleToGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRoleToGroupParamsWithHTTPClient(client *http.Client) *AddRoleToGroupParams {
	return &AddRoleToGroupParams{
		HTTPClient: client,
	}
}

/* AddRoleToGroupParams contains all the parameters to send to the API endpoint
   for the add role to group operation.

   Typically these are written to a http.Request.
*/
type AddRoleToGroupParams struct {

	// Body.
	Body *models.Role

	// GroupLocator.
	GroupLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add role to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleToGroupParams) WithDefaults() *AddRoleToGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add role to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleToGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add role to group params
func (o *AddRoleToGroupParams) WithTimeout(timeout time.Duration) *AddRoleToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add role to group params
func (o *AddRoleToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add role to group params
func (o *AddRoleToGroupParams) WithContext(ctx context.Context) *AddRoleToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add role to group params
func (o *AddRoleToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add role to group params
func (o *AddRoleToGroupParams) WithHTTPClient(client *http.Client) *AddRoleToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add role to group params
func (o *AddRoleToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add role to group params
func (o *AddRoleToGroupParams) WithBody(body *models.Role) *AddRoleToGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add role to group params
func (o *AddRoleToGroupParams) SetBody(body *models.Role) {
	o.Body = body
}

// WithGroupLocator adds the groupLocator to the add role to group params
func (o *AddRoleToGroupParams) WithGroupLocator(groupLocator string) *AddRoleToGroupParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the add role to group params
func (o *AddRoleToGroupParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WriteToRequest writes these params to a swagger request
func (o *AddRoleToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
