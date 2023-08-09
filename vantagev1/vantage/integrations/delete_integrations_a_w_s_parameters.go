// Code generated by go-swagger; DO NOT EDIT.

package integrations

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
)

// NewDeleteIntegrationsAWSParams creates a new DeleteIntegrationsAWSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIntegrationsAWSParams() *DeleteIntegrationsAWSParams {
	return &DeleteIntegrationsAWSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIntegrationsAWSParamsWithTimeout creates a new DeleteIntegrationsAWSParams object
// with the ability to set a timeout on a request.
func NewDeleteIntegrationsAWSParamsWithTimeout(timeout time.Duration) *DeleteIntegrationsAWSParams {
	return &DeleteIntegrationsAWSParams{
		timeout: timeout,
	}
}

// NewDeleteIntegrationsAWSParamsWithContext creates a new DeleteIntegrationsAWSParams object
// with the ability to set a context for a request.
func NewDeleteIntegrationsAWSParamsWithContext(ctx context.Context) *DeleteIntegrationsAWSParams {
	return &DeleteIntegrationsAWSParams{
		Context: ctx,
	}
}

// NewDeleteIntegrationsAWSParamsWithHTTPClient creates a new DeleteIntegrationsAWSParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIntegrationsAWSParamsWithHTTPClient(client *http.Client) *DeleteIntegrationsAWSParams {
	return &DeleteIntegrationsAWSParams{
		HTTPClient: client,
	}
}

/*
DeleteIntegrationsAWSParams contains all the parameters to send to the API endpoint

	for the delete integrations a w s operation.

	Typically these are written to a http.Request.
*/
type DeleteIntegrationsAWSParams struct {

	// AccessCredentialID.
	//
	// Format: int32
	AccessCredentialID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete integrations a w s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIntegrationsAWSParams) WithDefaults() *DeleteIntegrationsAWSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete integrations a w s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIntegrationsAWSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) WithTimeout(timeout time.Duration) *DeleteIntegrationsAWSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) WithContext(ctx context.Context) *DeleteIntegrationsAWSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) WithHTTPClient(client *http.Client) *DeleteIntegrationsAWSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessCredentialID adds the accessCredentialID to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) WithAccessCredentialID(accessCredentialID int32) *DeleteIntegrationsAWSParams {
	o.SetAccessCredentialID(accessCredentialID)
	return o
}

// SetAccessCredentialID adds the accessCredentialId to the delete integrations a w s params
func (o *DeleteIntegrationsAWSParams) SetAccessCredentialID(accessCredentialID int32) {
	o.AccessCredentialID = accessCredentialID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIntegrationsAWSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_credential_id
	if err := r.SetPathParam("access_credential_id", swag.FormatInt32(o.AccessCredentialID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
