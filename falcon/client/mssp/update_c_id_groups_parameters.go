// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateCIDGroupsParams creates a new UpdateCIDGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCIDGroupsParams() *UpdateCIDGroupsParams {
	return &UpdateCIDGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCIDGroupsParamsWithTimeout creates a new UpdateCIDGroupsParams object
// with the ability to set a timeout on a request.
func NewUpdateCIDGroupsParamsWithTimeout(timeout time.Duration) *UpdateCIDGroupsParams {
	return &UpdateCIDGroupsParams{
		timeout: timeout,
	}
}

// NewUpdateCIDGroupsParamsWithContext creates a new UpdateCIDGroupsParams object
// with the ability to set a context for a request.
func NewUpdateCIDGroupsParamsWithContext(ctx context.Context) *UpdateCIDGroupsParams {
	return &UpdateCIDGroupsParams{
		Context: ctx,
	}
}

// NewUpdateCIDGroupsParamsWithHTTPClient creates a new UpdateCIDGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCIDGroupsParamsWithHTTPClient(client *http.Client) *UpdateCIDGroupsParams {
	return &UpdateCIDGroupsParams{
		HTTPClient: client,
	}
}

/*
UpdateCIDGroupsParams contains all the parameters to send to the API endpoint

	for the update c ID groups operation.

	Typically these are written to a http.Request.
*/
type UpdateCIDGroupsParams struct {

	// Body.
	Body *models.DomainCIDGroupsRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update c ID groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCIDGroupsParams) WithDefaults() *UpdateCIDGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update c ID groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCIDGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update c ID groups params
func (o *UpdateCIDGroupsParams) WithTimeout(timeout time.Duration) *UpdateCIDGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update c ID groups params
func (o *UpdateCIDGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update c ID groups params
func (o *UpdateCIDGroupsParams) WithContext(ctx context.Context) *UpdateCIDGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update c ID groups params
func (o *UpdateCIDGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update c ID groups params
func (o *UpdateCIDGroupsParams) WithHTTPClient(client *http.Client) *UpdateCIDGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update c ID groups params
func (o *UpdateCIDGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update c ID groups params
func (o *UpdateCIDGroupsParams) WithBody(body *models.DomainCIDGroupsRequestV1) *UpdateCIDGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update c ID groups params
func (o *UpdateCIDGroupsParams) SetBody(body *models.DomainCIDGroupsRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCIDGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
