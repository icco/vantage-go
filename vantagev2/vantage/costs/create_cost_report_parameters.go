// Code generated by go-swagger; DO NOT EDIT.

package costs

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
	"github.com/go-openapi/swag"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// NewCreateCostReportParams creates a new CreateCostReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCostReportParams() *CreateCostReportParams {
	return &CreateCostReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCostReportParamsWithTimeout creates a new CreateCostReportParams object
// with the ability to set a timeout on a request.
func NewCreateCostReportParamsWithTimeout(timeout time.Duration) *CreateCostReportParams {
	return &CreateCostReportParams{
		timeout: timeout,
	}
}

// NewCreateCostReportParamsWithContext creates a new CreateCostReportParams object
// with the ability to set a context for a request.
func NewCreateCostReportParamsWithContext(ctx context.Context) *CreateCostReportParams {
	return &CreateCostReportParams{
		Context: ctx,
	}
}

// NewCreateCostReportParamsWithHTTPClient creates a new CreateCostReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCostReportParamsWithHTTPClient(client *http.Client) *CreateCostReportParams {
	return &CreateCostReportParams{
		HTTPClient: client,
	}
}

/*
CreateCostReportParams contains all the parameters to send to the API endpoint

	for the create cost report operation.

	Typically these are written to a http.Request.
*/
type CreateCostReportParams struct {

	// CostReports.
	CostReports *models.PostCostReports

	/* Groupings.

	   Grouping values for aggregating costs on the report. Valid groupings: account_id, billing_account_id, charge_type, cost_category, cost_subcategory, provider, region, resource_id, service, tag:<tag_value>. If providing multiple groupings, join as comma separated values: groupings=provider,service,region
	*/
	Groupings []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCostReportParams) WithDefaults() *CreateCostReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCostReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cost report params
func (o *CreateCostReportParams) WithTimeout(timeout time.Duration) *CreateCostReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cost report params
func (o *CreateCostReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cost report params
func (o *CreateCostReportParams) WithContext(ctx context.Context) *CreateCostReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cost report params
func (o *CreateCostReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cost report params
func (o *CreateCostReportParams) WithHTTPClient(client *http.Client) *CreateCostReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cost report params
func (o *CreateCostReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCostReports adds the costReports to the create cost report params
func (o *CreateCostReportParams) WithCostReports(costReports *models.PostCostReports) *CreateCostReportParams {
	o.SetCostReports(costReports)
	return o
}

// SetCostReports adds the costReports to the create cost report params
func (o *CreateCostReportParams) SetCostReports(costReports *models.PostCostReports) {
	o.CostReports = costReports
}

// WithGroupings adds the groupings to the create cost report params
func (o *CreateCostReportParams) WithGroupings(groupings []string) *CreateCostReportParams {
	o.SetGroupings(groupings)
	return o
}

// SetGroupings adds the groupings to the create cost report params
func (o *CreateCostReportParams) SetGroupings(groupings []string) {
	o.Groupings = groupings
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCostReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CostReports != nil {
		if err := r.SetBodyParam(o.CostReports); err != nil {
			return err
		}
	}

	if o.Groupings != nil {

		// binding items for groupings
		joinedGroupings := o.bindParamGroupings(reg)

		// query array param groupings
		if err := r.SetQueryParam("groupings", joinedGroupings...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateCostReport binds the parameter groupings
func (o *CreateCostReportParams) bindParamGroupings(formats strfmt.Registry) []string {
	groupingsIR := o.Groupings

	var groupingsIC []string
	for _, groupingsIIR := range groupingsIR { // explode []string

		groupingsIIV := groupingsIIR // string as string
		groupingsIC = append(groupingsIC, groupingsIIV)
	}

	// items.CollectionFormat: "csv"
	groupingsIS := swag.JoinByFormat(groupingsIC, "csv")

	return groupingsIS
}
