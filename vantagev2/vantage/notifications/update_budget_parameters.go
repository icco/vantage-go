// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// NewUpdateBudgetParams creates a new UpdateBudgetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBudgetParams() *UpdateBudgetParams {
	return &UpdateBudgetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBudgetParamsWithTimeout creates a new UpdateBudgetParams object
// with the ability to set a timeout on a request.
func NewUpdateBudgetParamsWithTimeout(timeout time.Duration) *UpdateBudgetParams {
	return &UpdateBudgetParams{
		timeout: timeout,
	}
}

// NewUpdateBudgetParamsWithContext creates a new UpdateBudgetParams object
// with the ability to set a context for a request.
func NewUpdateBudgetParamsWithContext(ctx context.Context) *UpdateBudgetParams {
	return &UpdateBudgetParams{
		Context: ctx,
	}
}

// NewUpdateBudgetParamsWithHTTPClient creates a new UpdateBudgetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBudgetParamsWithHTTPClient(client *http.Client) *UpdateBudgetParams {
	return &UpdateBudgetParams{
		HTTPClient: client,
	}
}

/*
UpdateBudgetParams contains all the parameters to send to the API endpoint

	for the update budget operation.

	Typically these are written to a http.Request.
*/
type UpdateBudgetParams struct {

	// BudgetToken.
	BudgetToken string

	// UpdateBudget.
	UpdateBudget *models.UpdateBudget

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update budget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBudgetParams) WithDefaults() *UpdateBudgetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update budget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBudgetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update budget params
func (o *UpdateBudgetParams) WithTimeout(timeout time.Duration) *UpdateBudgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update budget params
func (o *UpdateBudgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update budget params
func (o *UpdateBudgetParams) WithContext(ctx context.Context) *UpdateBudgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update budget params
func (o *UpdateBudgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update budget params
func (o *UpdateBudgetParams) WithHTTPClient(client *http.Client) *UpdateBudgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update budget params
func (o *UpdateBudgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBudgetToken adds the budgetToken to the update budget params
func (o *UpdateBudgetParams) WithBudgetToken(budgetToken string) *UpdateBudgetParams {
	o.SetBudgetToken(budgetToken)
	return o
}

// SetBudgetToken adds the budgetToken to the update budget params
func (o *UpdateBudgetParams) SetBudgetToken(budgetToken string) {
	o.BudgetToken = budgetToken
}

// WithUpdateBudget adds the updateBudget to the update budget params
func (o *UpdateBudgetParams) WithUpdateBudget(updateBudget *models.UpdateBudget) *UpdateBudgetParams {
	o.SetUpdateBudget(updateBudget)
	return o
}

// SetUpdateBudget adds the updateBudget to the update budget params
func (o *UpdateBudgetParams) SetUpdateBudget(updateBudget *models.UpdateBudget) {
	o.UpdateBudget = updateBudget
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBudgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param budget_token
	if err := r.SetPathParam("budget_token", o.BudgetToken); err != nil {
		return err
	}
	if o.UpdateBudget != nil {
		if err := r.SetBodyParam(o.UpdateBudget); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
