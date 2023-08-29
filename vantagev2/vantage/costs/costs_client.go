// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new costs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for costs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCostReport(params *CreateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostReportCreated, error)

	CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderCreated, error)

	CreateSavedFilter(params *CreateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSavedFilterCreated, error)

	DeleteCostReport(params *DeleteCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCostReportNoContent, error)

	DeleteSavedFilter(params *DeleteSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSavedFilterNoContent, error)

	GetCostReport(params *GetCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportOK, error)

	GetCostReports(params *GetCostReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportsOK, error)

	GetFolder(params *GetFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderOK, error)

	GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error)

	GetSavedFilter(params *GetSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFilterOK, error)

	GetSavedFilters(params *GetSavedFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFiltersOK, error)

	UpdateCostReport(params *UpdateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCostReportOK, error)

	UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error)

	UpdateSavedFilter(params *UpdateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSavedFilterCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCostReport Create a CostReport.
*/
func (a *Client) CreateCostReport(params *CreateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCostReport",
		Method:             "POST",
		PathPattern:        "/cost_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCostReportReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateCostReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateFolder Create a Folder for CostReports.
*/
func (a *Client) CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFolder",
		Method:             "POST",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateFolderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSavedFilter Create a SavedFilter for CostReports.
*/
func (a *Client) CreateSavedFilter(params *CreateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSavedFilterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSavedFilter",
		Method:             "POST",
		PathPattern:        "/saved_filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSavedFilterReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateSavedFilterCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCostReport Delete a CostReport.
*/
func (a *Client) DeleteCostReport(params *DeleteCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCostReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCostReport",
		Method:             "DELETE",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCostReportReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteCostReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSavedFilter Delete a SavedFilter for CostReports.
*/
func (a *Client) DeleteSavedFilter(params *DeleteSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSavedFilterNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSavedFilter",
		Method:             "DELETE",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSavedFilterReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteSavedFilterNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCostReport Return a CostReport.
*/
func (a *Client) GetCostReport(params *GetCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCostReport",
		Method:             "GET",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostReportReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetCostReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCostReports Return all CostReports.
*/
func (a *Client) GetCostReports(params *GetCostReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostReportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCostReports",
		Method:             "GET",
		PathPattern:        "/cost_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetCostReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCostReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolder Return a specific Folder for CostReports.
*/
func (a *Client) GetFolder(params *GetFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolder",
		Method:             "GET",
		PathPattern:        "/folders/{folder_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolders Return all Folders for CostReports.
*/
func (a *Client) GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoldersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolders",
		Method:             "GET",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoldersReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSavedFilter Return a specific SavedFilter.
*/
func (a *Client) GetSavedFilter(params *GetSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSavedFilter",
		Method:             "GET",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSavedFilterReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetSavedFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSavedFilters Return all SavedFilters that can be applied to a CostReport.
*/
func (a *Client) GetSavedFilters(params *GetSavedFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSavedFiltersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSavedFilters",
		Method:             "GET",
		PathPattern:        "/saved_filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSavedFiltersReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetSavedFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSavedFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCostReport Update a CostReport.
*/
func (a *Client) UpdateCostReport(params *UpdateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCostReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCostReport",
		Method:             "PUT",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCostReportReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateCostReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFolder Update a Folder for CostReports.
*/
func (a *Client) UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFolder",
		Method:             "PUT",
		PathPattern:        "/folders/{folder_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSavedFilter Update a SavedFilter for CostReports.
*/
func (a *Client) UpdateSavedFilter(params *UpdateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSavedFilterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSavedFilter",
		Method:             "PUT",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSavedFilterReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateSavedFilterCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
