// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddGroup(params *AddGroupParams, opts ...ClientOption) (*AddGroupOK, error)

	AddRoleAtScopeToGroup(params *AddRoleAtScopeToGroupParams, opts ...ClientOption) (*AddRoleAtScopeToGroupOK, error)

	AddRoleToGroup(params *AddRoleToGroupParams, opts ...ClientOption) (*AddRoleToGroupOK, error)

	DeleteGroup(params *DeleteGroupParams, opts ...ClientOption) error

	GetAllGroups(params *GetAllGroupsParams, opts ...ClientOption) (*GetAllGroupsOK, error)

	GetGroupParentGroups(params *GetGroupParentGroupsParams, opts ...ClientOption) (*GetGroupParentGroupsOK, error)

	GetGroupProperties(params *GetGroupPropertiesParams, opts ...ClientOption) (*GetGroupPropertiesOK, error)

	GetGroupProperty(params *GetGroupPropertyParams, opts ...ClientOption) (*GetGroupPropertyOK, error)

	GetGroupRoleAtScope(params *GetGroupRoleAtScopeParams, opts ...ClientOption) (*GetGroupRoleAtScopeOK, error)

	GetGroupRoles(params *GetGroupRolesParams, opts ...ClientOption) (*GetGroupRolesOK, error)

	GetUserGroupOfGroup(params *GetUserGroupOfGroupParams, opts ...ClientOption) (*GetUserGroupOfGroupOK, error)

	RemoveGroupProperty(params *RemoveGroupPropertyParams, opts ...ClientOption) error

	RemoveRoleAtScopeFromGroup(params *RemoveRoleAtScopeFromGroupParams, opts ...ClientOption) error

	SetGroupParentGroups(params *SetGroupParentGroupsParams, opts ...ClientOption) (*SetGroupParentGroupsOK, error)

	SetGroupProperty(params *SetGroupPropertyParams, opts ...ClientOption) (*SetGroupPropertyOK, error)

	SetGroupRoles(params *SetGroupRolesParams, opts ...ClientOption) (*SetGroupRolesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddGroup adds a new user group
*/
func (a *Client) AddGroup(params *AddGroupParams, opts ...ClientOption) (*AddGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addGroup",
		Method:             "POST",
		PathPattern:        "/app/rest/userGroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddRoleAtScopeToGroup adds a role with the specific scope to the matching user group
*/
func (a *Client) AddRoleAtScopeToGroup(params *AddRoleAtScopeToGroupParams, opts ...ClientOption) (*AddRoleAtScopeToGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleAtScopeToGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addRoleAtScopeToGroup",
		Method:             "POST",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleAtScopeToGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRoleAtScopeToGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addRoleAtScopeToGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddRoleToGroup adds a role to the matching user group
*/
func (a *Client) AddRoleToGroup(params *AddRoleToGroupParams, opts ...ClientOption) (*AddRoleToGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleToGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addRoleToGroup",
		Method:             "POST",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleToGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRoleToGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addRoleToGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteGroup deletes user group matching the locator
*/
func (a *Client) DeleteGroup(params *DeleteGroupParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteGroup",
		Method:             "DELETE",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  GetAllGroups gets all user groups
*/
func (a *Client) GetAllGroups(params *GetAllGroupsParams, opts ...ClientOption) (*GetAllGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupParentGroups gets parent groups of the matching user group
*/
func (a *Client) GetGroupParentGroups(params *GetGroupParentGroupsParams, opts ...ClientOption) (*GetGroupParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupParentGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupParentGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupParentGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupParentGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupParentGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupProperties gets properties of the matching user group
*/
func (a *Client) GetGroupProperties(params *GetGroupPropertiesParams, opts ...ClientOption) (*GetGroupPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupPropertiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupProperties",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupPropertiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupProperties: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupProperty gets a property of the matching user group
*/
func (a *Client) GetGroupProperty(params *GetGroupPropertyParams, opts ...ClientOption) (*GetGroupPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupProperty",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties/{name}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupPropertyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupProperty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupRoleAtScope gets a role with the specific scope of the matching user group
*/
func (a *Client) GetGroupRoleAtScope(params *GetGroupRoleAtScopeParams, opts ...ClientOption) (*GetGroupRoleAtScopeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupRoleAtScopeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupRoleAtScope",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupRoleAtScopeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupRoleAtScopeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupRoleAtScope: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupRoles gets all roles of the matching user group
*/
func (a *Client) GetGroupRoles(params *GetGroupRolesParams, opts ...ClientOption) (*GetGroupRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupRoles",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserGroupOfGroup gets user group matching the locator
*/
func (a *Client) GetUserGroupOfGroup(params *GetUserGroupOfGroupParams, opts ...ClientOption) (*GetUserGroupOfGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserGroupOfGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserGroupOfGroup",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserGroupOfGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserGroupOfGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGroupOfGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveGroupProperty removes a property of the matching user group
*/
func (a *Client) RemoveGroupProperty(params *RemoveGroupPropertyParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveGroupPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeGroupProperty",
		Method:             "DELETE",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveGroupPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  RemoveRoleAtScopeFromGroup removes a role with the specific scope from the matching user group
*/
func (a *Client) RemoveRoleAtScopeFromGroup(params *RemoveRoleAtScopeFromGroupParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveRoleAtScopeFromGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeRoleAtScopeFromGroup",
		Method:             "DELETE",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveRoleAtScopeFromGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  SetGroupParentGroups updates parent groups of the matching user group
*/
func (a *Client) SetGroupParentGroups(params *SetGroupParentGroupsParams, opts ...ClientOption) (*SetGroupParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetGroupParentGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setGroupParentGroups",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetGroupParentGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetGroupParentGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setGroupParentGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetGroupProperty updates a property of the matching user group
*/
func (a *Client) SetGroupProperty(params *SetGroupPropertyParams, opts ...ClientOption) (*SetGroupPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetGroupPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setGroupProperty",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties/{name}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetGroupPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetGroupPropertyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setGroupProperty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetGroupRoles updates roles of the matching user group
*/
func (a *Client) SetGroupRoles(params *SetGroupRolesParams, opts ...ClientOption) (*SetGroupRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetGroupRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setGroupRoles",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetGroupRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetGroupRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setGroupRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
