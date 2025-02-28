// Code generated by go-swagger; DO NOT EDIT.

package virtual_tags

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

// NewGetVirtualTagConfigParams creates a new GetVirtualTagConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualTagConfigParams() *GetVirtualTagConfigParams {
	return &GetVirtualTagConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualTagConfigParamsWithTimeout creates a new GetVirtualTagConfigParams object
// with the ability to set a timeout on a request.
func NewGetVirtualTagConfigParamsWithTimeout(timeout time.Duration) *GetVirtualTagConfigParams {
	return &GetVirtualTagConfigParams{
		timeout: timeout,
	}
}

// NewGetVirtualTagConfigParamsWithContext creates a new GetVirtualTagConfigParams object
// with the ability to set a context for a request.
func NewGetVirtualTagConfigParamsWithContext(ctx context.Context) *GetVirtualTagConfigParams {
	return &GetVirtualTagConfigParams{
		Context: ctx,
	}
}

// NewGetVirtualTagConfigParamsWithHTTPClient creates a new GetVirtualTagConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualTagConfigParamsWithHTTPClient(client *http.Client) *GetVirtualTagConfigParams {
	return &GetVirtualTagConfigParams{
		HTTPClient: client,
	}
}

/*
GetVirtualTagConfigParams contains all the parameters to send to the API endpoint

	for the get virtual tag config operation.

	Typically these are written to a http.Request.
*/
type GetVirtualTagConfigParams struct {

	// VirtualTagConfigToken.
	VirtualTagConfigToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual tag config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualTagConfigParams) WithDefaults() *GetVirtualTagConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual tag config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualTagConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get virtual tag config params
func (o *GetVirtualTagConfigParams) WithTimeout(timeout time.Duration) *GetVirtualTagConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual tag config params
func (o *GetVirtualTagConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual tag config params
func (o *GetVirtualTagConfigParams) WithContext(ctx context.Context) *GetVirtualTagConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual tag config params
func (o *GetVirtualTagConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual tag config params
func (o *GetVirtualTagConfigParams) WithHTTPClient(client *http.Client) *GetVirtualTagConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual tag config params
func (o *GetVirtualTagConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualTagConfigToken adds the virtualTagConfigToken to the get virtual tag config params
func (o *GetVirtualTagConfigParams) WithVirtualTagConfigToken(virtualTagConfigToken string) *GetVirtualTagConfigParams {
	o.SetVirtualTagConfigToken(virtualTagConfigToken)
	return o
}

// SetVirtualTagConfigToken adds the virtualTagConfigToken to the get virtual tag config params
func (o *GetVirtualTagConfigParams) SetVirtualTagConfigToken(virtualTagConfigToken string) {
	o.VirtualTagConfigToken = virtualTagConfigToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualTagConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtual_tag_config_token
	if err := r.SetPathParam("virtual_tag_config_token", o.VirtualTagConfigToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
