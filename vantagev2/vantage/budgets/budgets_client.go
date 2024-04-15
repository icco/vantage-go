// Code generated by go-swagger; DO NOT EDIT.

package budgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new budgets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for budgets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBudget(params *CreateBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBudgetCreated, error)

	DeleteBudget(params *DeleteBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBudgetNoContent, error)

	GetBudget(params *GetBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBudgetOK, error)

	GetBudgets(params *GetBudgetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBudgetsOK, error)

	UpdateBudget(params *UpdateBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBudgetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBudget Create a Budget.
*/
func (a *Client) CreateBudget(params *CreateBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBudgetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBudgetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBudget",
		Method:             "POST",
		PathPattern:        "/budgets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBudgetReader{formats: a.formats},
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
	success, ok := result.(*CreateBudgetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBudget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBudget Delete a Budget.
*/
func (a *Client) DeleteBudget(params *DeleteBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBudgetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBudgetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBudget",
		Method:             "DELETE",
		PathPattern:        "/budgets/{budget_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBudgetReader{formats: a.formats},
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
	success, ok := result.(*DeleteBudgetNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBudget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBudget Return a Budget.
*/
func (a *Client) GetBudget(params *GetBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBudgetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBudgetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBudget",
		Method:             "GET",
		PathPattern:        "/budgets/{budget_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBudgetReader{formats: a.formats},
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
	success, ok := result.(*GetBudgetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBudget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBudgets Return all Budgets.
*/
func (a *Client) GetBudgets(params *GetBudgetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBudgetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBudgetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBudgets",
		Method:             "GET",
		PathPattern:        "/budgets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBudgetsReader{formats: a.formats},
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
	success, ok := result.(*GetBudgetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBudgets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBudget Update a Budget.
*/
func (a *Client) UpdateBudget(params *UpdateBudgetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBudgetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBudgetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBudget",
		Method:             "PUT",
		PathPattern:        "/budgets/{budget_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBudgetReader{formats: a.formats},
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
	success, ok := result.(*UpdateBudgetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBudget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
