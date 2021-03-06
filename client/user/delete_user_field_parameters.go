// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewDeleteUserFieldParams creates a new DeleteUserFieldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserFieldParams() *DeleteUserFieldParams {
	return &DeleteUserFieldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserFieldParamsWithTimeout creates a new DeleteUserFieldParams object
// with the ability to set a timeout on a request.
func NewDeleteUserFieldParamsWithTimeout(timeout time.Duration) *DeleteUserFieldParams {
	return &DeleteUserFieldParams{
		timeout: timeout,
	}
}

// NewDeleteUserFieldParamsWithContext creates a new DeleteUserFieldParams object
// with the ability to set a context for a request.
func NewDeleteUserFieldParamsWithContext(ctx context.Context) *DeleteUserFieldParams {
	return &DeleteUserFieldParams{
		Context: ctx,
	}
}

// NewDeleteUserFieldParamsWithHTTPClient creates a new DeleteUserFieldParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserFieldParamsWithHTTPClient(client *http.Client) *DeleteUserFieldParams {
	return &DeleteUserFieldParams{
		HTTPClient: client,
	}
}

/* DeleteUserFieldParams contains all the parameters to send to the API endpoint
   for the delete user field operation.

   Typically these are written to a http.Request.
*/
type DeleteUserFieldParams struct {

	// Field.
	Field string

	// UserLocator.
	//
	// Format: UserLocator
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFieldParams) WithDefaults() *DeleteUserFieldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user field params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFieldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user field params
func (o *DeleteUserFieldParams) WithTimeout(timeout time.Duration) *DeleteUserFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user field params
func (o *DeleteUserFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user field params
func (o *DeleteUserFieldParams) WithContext(ctx context.Context) *DeleteUserFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user field params
func (o *DeleteUserFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user field params
func (o *DeleteUserFieldParams) WithHTTPClient(client *http.Client) *DeleteUserFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user field params
func (o *DeleteUserFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the delete user field params
func (o *DeleteUserFieldParams) WithField(field string) *DeleteUserFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the delete user field params
func (o *DeleteUserFieldParams) SetField(field string) {
	o.Field = field
}

// WithUserLocator adds the userLocator to the delete user field params
func (o *DeleteUserFieldParams) WithUserLocator(userLocator string) *DeleteUserFieldParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the delete user field params
func (o *DeleteUserFieldParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
