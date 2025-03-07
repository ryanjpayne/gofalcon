// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

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

// NewUpdateSensorVisibilityExclusionsV1Params creates a new UpdateSensorVisibilityExclusionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSensorVisibilityExclusionsV1Params() *UpdateSensorVisibilityExclusionsV1Params {
	return &UpdateSensorVisibilityExclusionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSensorVisibilityExclusionsV1ParamsWithTimeout creates a new UpdateSensorVisibilityExclusionsV1Params object
// with the ability to set a timeout on a request.
func NewUpdateSensorVisibilityExclusionsV1ParamsWithTimeout(timeout time.Duration) *UpdateSensorVisibilityExclusionsV1Params {
	return &UpdateSensorVisibilityExclusionsV1Params{
		timeout: timeout,
	}
}

// NewUpdateSensorVisibilityExclusionsV1ParamsWithContext creates a new UpdateSensorVisibilityExclusionsV1Params object
// with the ability to set a context for a request.
func NewUpdateSensorVisibilityExclusionsV1ParamsWithContext(ctx context.Context) *UpdateSensorVisibilityExclusionsV1Params {
	return &UpdateSensorVisibilityExclusionsV1Params{
		Context: ctx,
	}
}

// NewUpdateSensorVisibilityExclusionsV1ParamsWithHTTPClient creates a new UpdateSensorVisibilityExclusionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSensorVisibilityExclusionsV1ParamsWithHTTPClient(client *http.Client) *UpdateSensorVisibilityExclusionsV1Params {
	return &UpdateSensorVisibilityExclusionsV1Params{
		HTTPClient: client,
	}
}

/*
UpdateSensorVisibilityExclusionsV1Params contains all the parameters to send to the API endpoint

	for the update sensor visibility exclusions v1 operation.

	Typically these are written to a http.Request.
*/
type UpdateSensorVisibilityExclusionsV1Params struct {

	// Body.
	Body *models.RequestsSvExclusionUpdateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update sensor visibility exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSensorVisibilityExclusionsV1Params) WithDefaults() *UpdateSensorVisibilityExclusionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update sensor visibility exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSensorVisibilityExclusionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) WithTimeout(timeout time.Duration) *UpdateSensorVisibilityExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) WithContext(ctx context.Context) *UpdateSensorVisibilityExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) WithHTTPClient(client *http.Client) *UpdateSensorVisibilityExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) WithBody(body *models.RequestsSvExclusionUpdateReqV1) *UpdateSensorVisibilityExclusionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update sensor visibility exclusions v1 params
func (o *UpdateSensorVisibilityExclusionsV1Params) SetBody(body *models.RequestsSvExclusionUpdateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSensorVisibilityExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
