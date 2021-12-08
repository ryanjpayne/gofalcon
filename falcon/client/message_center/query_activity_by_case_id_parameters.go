// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// NewQueryActivityByCaseIDParams creates a new QueryActivityByCaseIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryActivityByCaseIDParams() *QueryActivityByCaseIDParams {
	return &QueryActivityByCaseIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryActivityByCaseIDParamsWithTimeout creates a new QueryActivityByCaseIDParams object
// with the ability to set a timeout on a request.
func NewQueryActivityByCaseIDParamsWithTimeout(timeout time.Duration) *QueryActivityByCaseIDParams {
	return &QueryActivityByCaseIDParams{
		timeout: timeout,
	}
}

// NewQueryActivityByCaseIDParamsWithContext creates a new QueryActivityByCaseIDParams object
// with the ability to set a context for a request.
func NewQueryActivityByCaseIDParamsWithContext(ctx context.Context) *QueryActivityByCaseIDParams {
	return &QueryActivityByCaseIDParams{
		Context: ctx,
	}
}

// NewQueryActivityByCaseIDParamsWithHTTPClient creates a new QueryActivityByCaseIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryActivityByCaseIDParamsWithHTTPClient(client *http.Client) *QueryActivityByCaseIDParams {
	return &QueryActivityByCaseIDParams{
		HTTPClient: client,
	}
}

/* QueryActivityByCaseIDParams contains all the parameters to send to the API endpoint
   for the query activity by case ID operation.

   Typically these are written to a http.Request.
*/
type QueryActivityByCaseIDParams struct {

	/* CaseID.

	   Case ID
	*/
	CaseID string

	/* Filter.

	   Optional filter and sort criteria in the form of an FQL query.
	*/
	Filter *string

	/* Limit.

	   The maximum records to return. [1-500]
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *string

	/* Sort.

	   The property to sort on, followed by a dot (.), followed by the sort direction, either "asc" or "desc".
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query activity by case ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryActivityByCaseIDParams) WithDefaults() *QueryActivityByCaseIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query activity by case ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryActivityByCaseIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithTimeout(timeout time.Duration) *QueryActivityByCaseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithContext(ctx context.Context) *QueryActivityByCaseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithHTTPClient(client *http.Client) *QueryActivityByCaseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseID adds the caseID to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithCaseID(caseID string) *QueryActivityByCaseIDParams {
	o.SetCaseID(caseID)
	return o
}

// SetCaseID adds the caseId to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetCaseID(caseID string) {
	o.CaseID = caseID
}

// WithFilter adds the filter to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithFilter(filter *string) *QueryActivityByCaseIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithLimit(limit *int64) *QueryActivityByCaseIDParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithOffset(offset *string) *QueryActivityByCaseIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) WithSort(sort *string) *QueryActivityByCaseIDParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query activity by case ID params
func (o *QueryActivityByCaseIDParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryActivityByCaseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param case_id
	qrCaseID := o.CaseID
	qCaseID := qrCaseID
	if qCaseID != "" {

		if err := r.SetQueryParam("case_id", qCaseID); err != nil {
			return err
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
