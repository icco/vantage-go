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

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// NewPutIntegrationsAWSParams creates a new PutIntegrationsAWSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutIntegrationsAWSParams() *PutIntegrationsAWSParams {
	return &PutIntegrationsAWSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutIntegrationsAWSParamsWithTimeout creates a new PutIntegrationsAWSParams object
// with the ability to set a timeout on a request.
func NewPutIntegrationsAWSParamsWithTimeout(timeout time.Duration) *PutIntegrationsAWSParams {
	return &PutIntegrationsAWSParams{
		timeout: timeout,
	}
}

// NewPutIntegrationsAWSParamsWithContext creates a new PutIntegrationsAWSParams object
// with the ability to set a context for a request.
func NewPutIntegrationsAWSParamsWithContext(ctx context.Context) *PutIntegrationsAWSParams {
	return &PutIntegrationsAWSParams{
		Context: ctx,
	}
}

// NewPutIntegrationsAWSParamsWithHTTPClient creates a new PutIntegrationsAWSParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutIntegrationsAWSParamsWithHTTPClient(client *http.Client) *PutIntegrationsAWSParams {
	return &PutIntegrationsAWSParams{
		HTTPClient: client,
	}
}

/*
PutIntegrationsAWSParams contains all the parameters to send to the API endpoint

	for the put integrations a w s operation.

	Typically these are written to a http.Request.
*/
type PutIntegrationsAWSParams struct {

	// AccessCredentialID.
	//
	// Format: int32
	AccessCredentialID int32

	// PutIntegrationsAWS.
	PutIntegrationsAWS *models.PutIntegrationsAWS

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put integrations a w s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIntegrationsAWSParams) WithDefaults() *PutIntegrationsAWSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put integrations a w s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIntegrationsAWSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put integrations a w s params
func (o *PutIntegrationsAWSParams) WithTimeout(timeout time.Duration) *PutIntegrationsAWSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put integrations a w s params
func (o *PutIntegrationsAWSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put integrations a w s params
func (o *PutIntegrationsAWSParams) WithContext(ctx context.Context) *PutIntegrationsAWSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put integrations a w s params
func (o *PutIntegrationsAWSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put integrations a w s params
func (o *PutIntegrationsAWSParams) WithHTTPClient(client *http.Client) *PutIntegrationsAWSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put integrations a w s params
func (o *PutIntegrationsAWSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessCredentialID adds the accessCredentialID to the put integrations a w s params
func (o *PutIntegrationsAWSParams) WithAccessCredentialID(accessCredentialID int32) *PutIntegrationsAWSParams {
	o.SetAccessCredentialID(accessCredentialID)
	return o
}

// SetAccessCredentialID adds the accessCredentialId to the put integrations a w s params
func (o *PutIntegrationsAWSParams) SetAccessCredentialID(accessCredentialID int32) {
	o.AccessCredentialID = accessCredentialID
}

// WithPutIntegrationsAWS adds the putIntegrationsAWS to the put integrations a w s params
func (o *PutIntegrationsAWSParams) WithPutIntegrationsAWS(putIntegrationsAWS *models.PutIntegrationsAWS) *PutIntegrationsAWSParams {
	o.SetPutIntegrationsAWS(putIntegrationsAWS)
	return o
}

// SetPutIntegrationsAWS adds the putIntegrationsAWS to the put integrations a w s params
func (o *PutIntegrationsAWSParams) SetPutIntegrationsAWS(putIntegrationsAWS *models.PutIntegrationsAWS) {
	o.PutIntegrationsAWS = putIntegrationsAWS
}

// WriteToRequest writes these params to a swagger request
func (o *PutIntegrationsAWSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_credential_id
	if err := r.SetPathParam("access_credential_id", swag.FormatInt32(o.AccessCredentialID)); err != nil {
		return err
	}
	if o.PutIntegrationsAWS != nil {
		if err := r.SetBodyParam(o.PutIntegrationsAWS); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
