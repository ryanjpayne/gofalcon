// Code generated by go-swagger; DO NOT EDIT.

package quarantine

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

// ActionUpdateCountReader is a Reader for the ActionUpdateCount structure.
type ActionUpdateCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionUpdateCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionUpdateCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewActionUpdateCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewActionUpdateCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewActionUpdateCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionUpdateCountOK creates a ActionUpdateCountOK with default headers values
func NewActionUpdateCountOK() *ActionUpdateCountOK {
	return &ActionUpdateCountOK{}
}

/*
ActionUpdateCountOK describes a response with status code 200, with default header values.

OK
*/
type ActionUpdateCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this action update count o k response has a 2xx status code
func (o *ActionUpdateCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action update count o k response has a 3xx status code
func (o *ActionUpdateCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action update count o k response has a 4xx status code
func (o *ActionUpdateCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action update count o k response has a 5xx status code
func (o *ActionUpdateCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action update count o k response a status code equal to that given
func (o *ActionUpdateCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the action update count o k response
func (o *ActionUpdateCountOK) Code() int {
	return 200
}

func (o *ActionUpdateCountOK) Error() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountOK  %+v", 200, o.Payload)
}

func (o *ActionUpdateCountOK) String() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountOK  %+v", 200, o.Payload)
}

func (o *ActionUpdateCountOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *ActionUpdateCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionUpdateCountForbidden creates a ActionUpdateCountForbidden with default headers values
func NewActionUpdateCountForbidden() *ActionUpdateCountForbidden {
	return &ActionUpdateCountForbidden{}
}

/*
ActionUpdateCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ActionUpdateCountForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this action update count forbidden response has a 2xx status code
func (o *ActionUpdateCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action update count forbidden response has a 3xx status code
func (o *ActionUpdateCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action update count forbidden response has a 4xx status code
func (o *ActionUpdateCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this action update count forbidden response has a 5xx status code
func (o *ActionUpdateCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this action update count forbidden response a status code equal to that given
func (o *ActionUpdateCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the action update count forbidden response
func (o *ActionUpdateCountForbidden) Code() int {
	return 403
}

func (o *ActionUpdateCountForbidden) Error() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountForbidden  %+v", 403, o.Payload)
}

func (o *ActionUpdateCountForbidden) String() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountForbidden  %+v", 403, o.Payload)
}

func (o *ActionUpdateCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ActionUpdateCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewActionUpdateCountTooManyRequests creates a ActionUpdateCountTooManyRequests with default headers values
func NewActionUpdateCountTooManyRequests() *ActionUpdateCountTooManyRequests {
	return &ActionUpdateCountTooManyRequests{}
}

/*
ActionUpdateCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ActionUpdateCountTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this action update count too many requests response has a 2xx status code
func (o *ActionUpdateCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action update count too many requests response has a 3xx status code
func (o *ActionUpdateCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action update count too many requests response has a 4xx status code
func (o *ActionUpdateCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this action update count too many requests response has a 5xx status code
func (o *ActionUpdateCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this action update count too many requests response a status code equal to that given
func (o *ActionUpdateCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the action update count too many requests response
func (o *ActionUpdateCountTooManyRequests) Code() int {
	return 429
}

func (o *ActionUpdateCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ActionUpdateCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] actionUpdateCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ActionUpdateCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ActionUpdateCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewActionUpdateCountDefault creates a ActionUpdateCountDefault with default headers values
func NewActionUpdateCountDefault(code int) *ActionUpdateCountDefault {
	return &ActionUpdateCountDefault{
		_statusCode: code,
	}
}

/*
ActionUpdateCountDefault describes a response with status code -1, with default header values.

OK
*/
type ActionUpdateCountDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this action update count default response has a 2xx status code
func (o *ActionUpdateCountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this action update count default response has a 3xx status code
func (o *ActionUpdateCountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this action update count default response has a 4xx status code
func (o *ActionUpdateCountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this action update count default response has a 5xx status code
func (o *ActionUpdateCountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this action update count default response a status code equal to that given
func (o *ActionUpdateCountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the action update count default response
func (o *ActionUpdateCountDefault) Code() int {
	return o._statusCode
}

func (o *ActionUpdateCountDefault) Error() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] ActionUpdateCount default  %+v", o._statusCode, o.Payload)
}

func (o *ActionUpdateCountDefault) String() string {
	return fmt.Sprintf("[GET /quarantine/aggregates/action-update-count/v1][%d] ActionUpdateCount default  %+v", o._statusCode, o.Payload)
}

func (o *ActionUpdateCountDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *ActionUpdateCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
