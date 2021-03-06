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

// NewAddAgentRequirementToBuildTypeParams creates a new AddAgentRequirementToBuildTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAgentRequirementToBuildTypeParams() *AddAgentRequirementToBuildTypeParams {
	return &AddAgentRequirementToBuildTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAgentRequirementToBuildTypeParamsWithTimeout creates a new AddAgentRequirementToBuildTypeParams object
// with the ability to set a timeout on a request.
func NewAddAgentRequirementToBuildTypeParamsWithTimeout(timeout time.Duration) *AddAgentRequirementToBuildTypeParams {
	return &AddAgentRequirementToBuildTypeParams{
		timeout: timeout,
	}
}

// NewAddAgentRequirementToBuildTypeParamsWithContext creates a new AddAgentRequirementToBuildTypeParams object
// with the ability to set a context for a request.
func NewAddAgentRequirementToBuildTypeParamsWithContext(ctx context.Context) *AddAgentRequirementToBuildTypeParams {
	return &AddAgentRequirementToBuildTypeParams{
		Context: ctx,
	}
}

// NewAddAgentRequirementToBuildTypeParamsWithHTTPClient creates a new AddAgentRequirementToBuildTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAgentRequirementToBuildTypeParamsWithHTTPClient(client *http.Client) *AddAgentRequirementToBuildTypeParams {
	return &AddAgentRequirementToBuildTypeParams{
		HTTPClient: client,
	}
}

/* AddAgentRequirementToBuildTypeParams contains all the parameters to send to the API endpoint
   for the add agent requirement to build type operation.

   Typically these are written to a http.Request.
*/
type AddAgentRequirementToBuildTypeParams struct {

	// Body.
	Body *models.AgentRequirement

	// BtLocator.
	//
	// Format: BuildTypeLocator
	BtLocator string

	// Fields.
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add agent requirement to build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAgentRequirementToBuildTypeParams) WithDefaults() *AddAgentRequirementToBuildTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add agent requirement to build type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAgentRequirementToBuildTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithTimeout(timeout time.Duration) *AddAgentRequirementToBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithContext(ctx context.Context) *AddAgentRequirementToBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithHTTPClient(client *http.Client) *AddAgentRequirementToBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithBody(body *models.AgentRequirement) *AddAgentRequirementToBuildTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetBody(body *models.AgentRequirement) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithBtLocator(btLocator string) *AddAgentRequirementToBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) WithFields(fields *string) *AddAgentRequirementToBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add agent requirement to build type params
func (o *AddAgentRequirementToBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddAgentRequirementToBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
