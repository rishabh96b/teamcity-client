// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

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

// NewDeleteProjectFromAgentPoolParams creates a new DeleteProjectFromAgentPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectFromAgentPoolParams() *DeleteProjectFromAgentPoolParams {
	return &DeleteProjectFromAgentPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectFromAgentPoolParamsWithTimeout creates a new DeleteProjectFromAgentPoolParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectFromAgentPoolParamsWithTimeout(timeout time.Duration) *DeleteProjectFromAgentPoolParams {
	return &DeleteProjectFromAgentPoolParams{
		timeout: timeout,
	}
}

// NewDeleteProjectFromAgentPoolParamsWithContext creates a new DeleteProjectFromAgentPoolParams object
// with the ability to set a context for a request.
func NewDeleteProjectFromAgentPoolParamsWithContext(ctx context.Context) *DeleteProjectFromAgentPoolParams {
	return &DeleteProjectFromAgentPoolParams{
		Context: ctx,
	}
}

// NewDeleteProjectFromAgentPoolParamsWithHTTPClient creates a new DeleteProjectFromAgentPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectFromAgentPoolParamsWithHTTPClient(client *http.Client) *DeleteProjectFromAgentPoolParams {
	return &DeleteProjectFromAgentPoolParams{
		HTTPClient: client,
	}
}

/* DeleteProjectFromAgentPoolParams contains all the parameters to send to the API endpoint
   for the delete project from agent pool operation.

   Typically these are written to a http.Request.
*/
type DeleteProjectFromAgentPoolParams struct {

	// AgentPoolLocator.
	//
	// Format: AgentPoolLocator
	AgentPoolLocator string

	// ProjectLocator.
	//
	// Format: ProjectLocator
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project from agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectFromAgentPoolParams) WithDefaults() *DeleteProjectFromAgentPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project from agent pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectFromAgentPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) WithTimeout(timeout time.Duration) *DeleteProjectFromAgentPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) WithContext(ctx context.Context) *DeleteProjectFromAgentPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) WithHTTPClient(client *http.Client) *DeleteProjectFromAgentPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) WithAgentPoolLocator(agentPoolLocator string) *DeleteProjectFromAgentPoolParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WithProjectLocator adds the projectLocator to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) WithProjectLocator(projectLocator string) *DeleteProjectFromAgentPoolParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the delete project from agent pool params
func (o *DeleteProjectFromAgentPoolParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectFromAgentPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
