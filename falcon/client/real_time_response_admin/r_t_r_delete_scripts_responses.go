// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRDeleteScriptsReader is a Reader for the RTRDeleteScripts structure.
type RTRDeleteScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeleteScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRDeleteScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeleteScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeleteScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRDeleteScriptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeleteScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRDeleteScriptsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRDeleteScriptsOK creates a RTRDeleteScriptsOK with default headers values
func NewRTRDeleteScriptsOK() *RTRDeleteScriptsOK {
	return &RTRDeleteScriptsOK{}
}

/*
RTRDeleteScriptsOK describes a response with status code 200, with default header values.

OK
*/
type RTRDeleteScriptsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete scripts o k response has a 2xx status code
func (o *RTRDeleteScriptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r delete scripts o k response has a 3xx status code
func (o *RTRDeleteScriptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete scripts o k response has a 4xx status code
func (o *RTRDeleteScriptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete scripts o k response has a 5xx status code
func (o *RTRDeleteScriptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete scripts o k response a status code equal to that given
func (o *RTRDeleteScriptsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r delete scripts o k response
func (o *RTRDeleteScriptsOK) Code() int {
	return 200
}

func (o *RTRDeleteScriptsOK) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRDeleteScriptsOK) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRDeleteScriptsOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteScriptsBadRequest creates a RTRDeleteScriptsBadRequest with default headers values
func NewRTRDeleteScriptsBadRequest() *RTRDeleteScriptsBadRequest {
	return &RTRDeleteScriptsBadRequest{}
}

/*
RTRDeleteScriptsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRDeleteScriptsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete scripts bad request response has a 2xx status code
func (o *RTRDeleteScriptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete scripts bad request response has a 3xx status code
func (o *RTRDeleteScriptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete scripts bad request response has a 4xx status code
func (o *RTRDeleteScriptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete scripts bad request response has a 5xx status code
func (o *RTRDeleteScriptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete scripts bad request response a status code equal to that given
func (o *RTRDeleteScriptsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r delete scripts bad request response
func (o *RTRDeleteScriptsBadRequest) Code() int {
	return 400
}

func (o *RTRDeleteScriptsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteScriptsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteScriptsBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteScriptsForbidden creates a RTRDeleteScriptsForbidden with default headers values
func NewRTRDeleteScriptsForbidden() *RTRDeleteScriptsForbidden {
	return &RTRDeleteScriptsForbidden{}
}

/*
RTRDeleteScriptsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRDeleteScriptsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete scripts forbidden response has a 2xx status code
func (o *RTRDeleteScriptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete scripts forbidden response has a 3xx status code
func (o *RTRDeleteScriptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete scripts forbidden response has a 4xx status code
func (o *RTRDeleteScriptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete scripts forbidden response has a 5xx status code
func (o *RTRDeleteScriptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete scripts forbidden response a status code equal to that given
func (o *RTRDeleteScriptsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r delete scripts forbidden response
func (o *RTRDeleteScriptsForbidden) Code() int {
	return 403
}

func (o *RTRDeleteScriptsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteScriptsForbidden) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteScriptsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteScriptsNotFound creates a RTRDeleteScriptsNotFound with default headers values
func NewRTRDeleteScriptsNotFound() *RTRDeleteScriptsNotFound {
	return &RTRDeleteScriptsNotFound{}
}

/*
RTRDeleteScriptsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRDeleteScriptsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete scripts not found response has a 2xx status code
func (o *RTRDeleteScriptsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete scripts not found response has a 3xx status code
func (o *RTRDeleteScriptsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete scripts not found response has a 4xx status code
func (o *RTRDeleteScriptsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete scripts not found response has a 5xx status code
func (o *RTRDeleteScriptsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete scripts not found response a status code equal to that given
func (o *RTRDeleteScriptsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r delete scripts not found response
func (o *RTRDeleteScriptsNotFound) Code() int {
	return 404
}

func (o *RTRDeleteScriptsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteScriptsNotFound) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteScriptsNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteScriptsTooManyRequests creates a RTRDeleteScriptsTooManyRequests with default headers values
func NewRTRDeleteScriptsTooManyRequests() *RTRDeleteScriptsTooManyRequests {
	return &RTRDeleteScriptsTooManyRequests{}
}

/*
RTRDeleteScriptsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRDeleteScriptsTooManyRequests struct {

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

// IsSuccess returns true when this r t r delete scripts too many requests response has a 2xx status code
func (o *RTRDeleteScriptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete scripts too many requests response has a 3xx status code
func (o *RTRDeleteScriptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete scripts too many requests response has a 4xx status code
func (o *RTRDeleteScriptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete scripts too many requests response has a 5xx status code
func (o *RTRDeleteScriptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete scripts too many requests response a status code equal to that given
func (o *RTRDeleteScriptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r delete scripts too many requests response
func (o *RTRDeleteScriptsTooManyRequests) Code() int {
	return 429
}

func (o *RTRDeleteScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteScriptsTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] rTRDeleteScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteScriptsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteScriptsDefault creates a RTRDeleteScriptsDefault with default headers values
func NewRTRDeleteScriptsDefault(code int) *RTRDeleteScriptsDefault {
	return &RTRDeleteScriptsDefault{
		_statusCode: code,
	}
}

/*
RTRDeleteScriptsDefault describes a response with status code -1, with default header values.

OK
*/
type RTRDeleteScriptsDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete scripts default response has a 2xx status code
func (o *RTRDeleteScriptsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r delete scripts default response has a 3xx status code
func (o *RTRDeleteScriptsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r delete scripts default response has a 4xx status code
func (o *RTRDeleteScriptsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r delete scripts default response has a 5xx status code
func (o *RTRDeleteScriptsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r delete scripts default response a status code equal to that given
func (o *RTRDeleteScriptsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the r t r delete scripts default response
func (o *RTRDeleteScriptsDefault) Code() int {
	return o._statusCode
}

func (o *RTRDeleteScriptsDefault) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] RTR-DeleteScripts default  %+v", o._statusCode, o.Payload)
}

func (o *RTRDeleteScriptsDefault) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/scripts/v1][%d] RTR-DeleteScripts default  %+v", o._statusCode, o.Payload)
}

func (o *RTRDeleteScriptsDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteScriptsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
