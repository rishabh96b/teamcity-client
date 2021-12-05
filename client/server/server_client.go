// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddLicenseKeys(params *AddLicenseKeysParams, opts ...ClientOption) (*AddLicenseKeysOK, error)

	DeleteLicenseKey(params *DeleteLicenseKeyParams, opts ...ClientOption) error

	DownloadFileOfServer(params *DownloadFileOfServerParams, opts ...ClientOption) error

	GetAllMetrics(params *GetAllMetricsParams, opts ...ClientOption) (*GetAllMetricsOK, error)

	GetAllPlugins(params *GetAllPluginsParams, opts ...ClientOption) (*GetAllPluginsOK, error)

	GetBackupStatus(params *GetBackupStatusParams, opts ...ClientOption) (*GetBackupStatusOK, error)

	GetFileMetadataOfServer(params *GetFileMetadataOfServerParams, opts ...ClientOption) (*GetFileMetadataOfServerOK, error)

	GetFilesListForSubpathOfServer(params *GetFilesListForSubpathOfServerParams, opts ...ClientOption) (*GetFilesListForSubpathOfServerOK, error)

	GetFilesListOfServer(params *GetFilesListOfServerParams, opts ...ClientOption) (*GetFilesListOfServerOK, error)

	GetLicenseKey(params *GetLicenseKeyParams, opts ...ClientOption) (*GetLicenseKeyOK, error)

	GetLicenseKeys(params *GetLicenseKeysParams, opts ...ClientOption) (*GetLicenseKeysOK, error)

	GetLicensingData(params *GetLicensingDataParams, opts ...ClientOption) (*GetLicensingDataOK, error)

	GetServerField(params *GetServerFieldParams, opts ...ClientOption) (*GetServerFieldOK, error)

	GetServerInfo(params *GetServerInfoParams, opts ...ClientOption) (*GetServerInfoOK, error)

	GetZippedFileOfServer(params *GetZippedFileOfServerParams, opts ...ClientOption) error

	StartBackup(params *StartBackupParams, opts ...ClientOption) (*StartBackupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddLicenseKeys adds license keys
*/
func (a *Client) AddLicenseKeys(params *AddLicenseKeysParams, opts ...ClientOption) (*AddLicenseKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLicenseKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addLicenseKeys",
		Method:             "POST",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddLicenseKeysReader{formats: a.formats},
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
	success, ok := result.(*AddLicenseKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addLicenseKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLicenseKey deletes a license key
*/
func (a *Client) DeleteLicenseKey(params *DeleteLicenseKeyParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLicenseKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteLicenseKey",
		Method:             "DELETE",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys/{licenseKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLicenseKeyReader{formats: a.formats},
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
  DownloadFileOfServer downloads specific file
*/
func (a *Client) DownloadFileOfServer(params *DownloadFileOfServerParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadFileOfServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadFileOfServer",
		Method:             "GET",
		PathPattern:        "/app/rest/server/files/{areaId}/files{path}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadFileOfServerReader{formats: a.formats},
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
  GetAllMetrics gets metrics
*/
func (a *Client) GetAllMetrics(params *GetAllMetricsParams, opts ...ClientOption) (*GetAllMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllMetrics",
		Method:             "GET",
		PathPattern:        "/app/rest/server/metrics",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetAllMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllPlugins gets all plugins
*/
func (a *Client) GetAllPlugins(params *GetAllPluginsParams, opts ...ClientOption) (*GetAllPluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPluginsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllPlugins",
		Method:             "GET",
		PathPattern:        "/app/rest/server/plugins",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllPluginsReader{formats: a.formats},
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
	success, ok := result.(*GetAllPluginsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllPlugins: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBackupStatus gets the latest backup status
*/
func (a *Client) GetBackupStatus(params *GetBackupStatusParams, opts ...ClientOption) (*GetBackupStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBackupStatus",
		Method:             "GET",
		PathPattern:        "/app/rest/server/backup",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupStatusReader{formats: a.formats},
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
	success, ok := result.(*GetBackupStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBackupStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFileMetadataOfServer gets metadata of specific file
*/
func (a *Client) GetFileMetadataOfServer(params *GetFileMetadataOfServerParams, opts ...ClientOption) (*GetFileMetadataOfServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileMetadataOfServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileMetadataOfServer",
		Method:             "GET",
		PathPattern:        "/app/rest/server/files/{areaId}/metadata{path}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileMetadataOfServerReader{formats: a.formats},
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
	success, ok := result.(*GetFileMetadataOfServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFileMetadataOfServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFilesListForSubpathOfServer lists files under this path
*/
func (a *Client) GetFilesListForSubpathOfServer(params *GetFilesListForSubpathOfServerParams, opts ...ClientOption) (*GetFilesListForSubpathOfServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesListForSubpathOfServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFilesListForSubpathOfServer",
		Method:             "GET",
		PathPattern:        "/app/rest/server/files/{areaId}/{path}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilesListForSubpathOfServerReader{formats: a.formats},
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
	success, ok := result.(*GetFilesListForSubpathOfServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFilesListForSubpathOfServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFilesListOfServer lists all files
*/
func (a *Client) GetFilesListOfServer(params *GetFilesListOfServerParams, opts ...ClientOption) (*GetFilesListOfServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesListOfServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFilesListOfServer",
		Method:             "GET",
		PathPattern:        "/app/rest/server/files/{areaId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilesListOfServerReader{formats: a.formats},
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
	success, ok := result.(*GetFilesListOfServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFilesListOfServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLicenseKey gets a license key
*/
func (a *Client) GetLicenseKey(params *GetLicenseKeyParams, opts ...ClientOption) (*GetLicenseKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseKey",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys/{licenseKey}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeyReader{formats: a.formats},
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
	success, ok := result.(*GetLicenseKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLicenseKeys gets all license keys
*/
func (a *Client) GetLicenseKeys(params *GetLicenseKeysParams, opts ...ClientOption) (*GetLicenseKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseKeys",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeysReader{formats: a.formats},
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
	success, ok := result.(*GetLicenseKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLicensingData gets the licensing data
*/
func (a *Client) GetLicensingData(params *GetLicensingDataParams, opts ...ClientOption) (*GetLicensingDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicensingDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicensingData",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicensingDataReader{formats: a.formats},
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
	success, ok := result.(*GetLicensingDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicensingData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServerField gets a field of the server info
*/
func (a *Client) GetServerField(params *GetServerFieldParams, opts ...ClientOption) (*GetServerFieldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerFieldParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServerField",
		Method:             "GET",
		PathPattern:        "/app/rest/server/{field}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerFieldReader{formats: a.formats},
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
	success, ok := result.(*GetServerFieldOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServerField: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServerInfo gets the server info
*/
func (a *Client) GetServerInfo(params *GetServerInfoParams, opts ...ClientOption) (*GetServerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServerInfo",
		Method:             "GET",
		PathPattern:        "/app/rest/server",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerInfoReader{formats: a.formats},
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
	success, ok := result.(*GetServerInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServerInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZippedFileOfServer gets specific file zipped
*/
func (a *Client) GetZippedFileOfServer(params *GetZippedFileOfServerParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZippedFileOfServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getZippedFileOfServer",
		Method:             "GET",
		PathPattern:        "/app/rest/server/files/{areaId}/archived{path}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZippedFileOfServerReader{formats: a.formats},
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
  StartBackup starts a new backup
*/
func (a *Client) StartBackup(params *StartBackupParams, opts ...ClientOption) (*StartBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startBackup",
		Method:             "POST",
		PathPattern:        "/app/rest/server/backup",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartBackupReader{formats: a.formats},
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
	success, ok := result.(*StartBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}