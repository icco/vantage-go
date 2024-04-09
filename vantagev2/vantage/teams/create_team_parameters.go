// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewCreateTeamParams creates a new CreateTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTeamParams() *CreateTeamParams {
	return &CreateTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTeamParamsWithTimeout creates a new CreateTeamParams object
// with the ability to set a timeout on a request.
func NewCreateTeamParamsWithTimeout(timeout time.Duration) *CreateTeamParams {
	return &CreateTeamParams{
		timeout: timeout,
	}
}

// NewCreateTeamParamsWithContext creates a new CreateTeamParams object
// with the ability to set a context for a request.
func NewCreateTeamParamsWithContext(ctx context.Context) *CreateTeamParams {
	return &CreateTeamParams{
		Context: ctx,
	}
}

// NewCreateTeamParamsWithHTTPClient creates a new CreateTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTeamParamsWithHTTPClient(client *http.Client) *CreateTeamParams {
	return &CreateTeamParams{
		HTTPClient: client,
	}
}

/*
CreateTeamParams contains all the parameters to send to the API endpoint

	for the create team operation.

	Typically these are written to a http.Request.
*/
type CreateTeamParams struct {

	// CreateTeam.
	CreateTeam *models.CreateTeam

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTeamParams) WithDefaults() *CreateTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create team params
func (o *CreateTeamParams) WithTimeout(timeout time.Duration) *CreateTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create team params
func (o *CreateTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create team params
func (o *CreateTeamParams) WithContext(ctx context.Context) *CreateTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create team params
func (o *CreateTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create team params
func (o *CreateTeamParams) WithHTTPClient(client *http.Client) *CreateTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create team params
func (o *CreateTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateTeam adds the createTeam to the create team params
func (o *CreateTeamParams) WithCreateTeam(createTeam *models.CreateTeam) *CreateTeamParams {
	o.SetCreateTeam(createTeam)
	return o
}

// SetCreateTeam adds the createTeam to the create team params
func (o *CreateTeamParams) SetCreateTeam(createTeam *models.CreateTeam) {
	o.CreateTeam = createTeam
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateTeam != nil {
		if err := r.SetBodyParam(o.CreateTeam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
