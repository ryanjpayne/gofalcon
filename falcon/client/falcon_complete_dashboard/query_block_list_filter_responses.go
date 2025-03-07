// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryBlockListFilterReader is a Reader for the QueryBlockListFilter structure.
type QueryBlockListFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryBlockListFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryBlockListFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryBlockListFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryBlockListFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryBlockListFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryBlockListFilterOK creates a QueryBlockListFilterOK with default headers values
func NewQueryBlockListFilterOK() *QueryBlockListFilterOK {
	return &QueryBlockListFilterOK{}
}

/*
QueryBlockListFilterOK describes a response with status code 200, with default header values.

OK
*/
type QueryBlockListFilterOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query block list filter o k response has a 2xx status code
func (o *QueryBlockListFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query block list filter o k response has a 3xx status code
func (o *QueryBlockListFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query block list filter o k response has a 4xx status code
func (o *QueryBlockListFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query block list filter o k response has a 5xx status code
func (o *QueryBlockListFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query block list filter o k response a status code equal to that given
func (o *QueryBlockListFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query block list filter o k response
func (o *QueryBlockListFilterOK) Code() int {
	return 200
}

func (o *QueryBlockListFilterOK) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterOK  %+v", 200, o.Payload)
}

func (o *QueryBlockListFilterOK) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterOK  %+v", 200, o.Payload)
}

func (o *QueryBlockListFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryBlockListFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryBlockListFilterForbidden creates a QueryBlockListFilterForbidden with default headers values
func NewQueryBlockListFilterForbidden() *QueryBlockListFilterForbidden {
	return &QueryBlockListFilterForbidden{}
}

/*
QueryBlockListFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryBlockListFilterForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query block list filter forbidden response has a 2xx status code
func (o *QueryBlockListFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query block list filter forbidden response has a 3xx status code
func (o *QueryBlockListFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query block list filter forbidden response has a 4xx status code
func (o *QueryBlockListFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query block list filter forbidden response has a 5xx status code
func (o *QueryBlockListFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query block list filter forbidden response a status code equal to that given
func (o *QueryBlockListFilterForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query block list filter forbidden response
func (o *QueryBlockListFilterForbidden) Code() int {
	return 403
}

func (o *QueryBlockListFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryBlockListFilterForbidden) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryBlockListFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryBlockListFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryBlockListFilterTooManyRequests creates a QueryBlockListFilterTooManyRequests with default headers values
func NewQueryBlockListFilterTooManyRequests() *QueryBlockListFilterTooManyRequests {
	return &QueryBlockListFilterTooManyRequests{}
}

/*
QueryBlockListFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryBlockListFilterTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query block list filter too many requests response has a 2xx status code
func (o *QueryBlockListFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query block list filter too many requests response has a 3xx status code
func (o *QueryBlockListFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query block list filter too many requests response has a 4xx status code
func (o *QueryBlockListFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query block list filter too many requests response has a 5xx status code
func (o *QueryBlockListFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query block list filter too many requests response a status code equal to that given
func (o *QueryBlockListFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query block list filter too many requests response
func (o *QueryBlockListFilterTooManyRequests) Code() int {
	return 429
}

func (o *QueryBlockListFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryBlockListFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] queryBlockListFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryBlockListFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryBlockListFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryBlockListFilterDefault creates a QueryBlockListFilterDefault with default headers values
func NewQueryBlockListFilterDefault(code int) *QueryBlockListFilterDefault {
	return &QueryBlockListFilterDefault{
		_statusCode: code,
	}
}

/*
QueryBlockListFilterDefault describes a response with status code -1, with default header values.

OK
*/
type QueryBlockListFilterDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query block list filter default response has a 2xx status code
func (o *QueryBlockListFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query block list filter default response has a 3xx status code
func (o *QueryBlockListFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query block list filter default response has a 4xx status code
func (o *QueryBlockListFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query block list filter default response has a 5xx status code
func (o *QueryBlockListFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query block list filter default response a status code equal to that given
func (o *QueryBlockListFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query block list filter default response
func (o *QueryBlockListFilterDefault) Code() int {
	return o._statusCode
}

func (o *QueryBlockListFilterDefault) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] QueryBlockListFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryBlockListFilterDefault) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/blocklist/v1][%d] QueryBlockListFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryBlockListFilterDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryBlockListFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
