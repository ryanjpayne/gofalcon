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

// RTRListPutFilesReader is a Reader for the RTRListPutFiles structure.
type RTRListPutFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListPutFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListPutFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListPutFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListPutFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListPutFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListPutFilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRListPutFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRListPutFilesOK creates a RTRListPutFilesOK with default headers values
func NewRTRListPutFilesOK() *RTRListPutFilesOK {
	return &RTRListPutFilesOK{}
}

/*
RTRListPutFilesOK describes a response with status code 200, with default header values.

OK
*/
type RTRListPutFilesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.BinservclientMsaPutFileResponse
}

// IsSuccess returns true when this r t r list put files o k response has a 2xx status code
func (o *RTRListPutFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r list put files o k response has a 3xx status code
func (o *RTRListPutFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list put files o k response has a 4xx status code
func (o *RTRListPutFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list put files o k response has a 5xx status code
func (o *RTRListPutFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list put files o k response a status code equal to that given
func (o *RTRListPutFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r list put files o k response
func (o *RTRListPutFilesOK) Code() int {
	return 200
}

func (o *RTRListPutFilesOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRListPutFilesOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRListPutFilesOK) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListPutFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesBadRequest creates a RTRListPutFilesBadRequest with default headers values
func NewRTRListPutFilesBadRequest() *RTRListPutFilesBadRequest {
	return &RTRListPutFilesBadRequest{}
}

/*
RTRListPutFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListPutFilesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list put files bad request response has a 2xx status code
func (o *RTRListPutFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list put files bad request response has a 3xx status code
func (o *RTRListPutFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list put files bad request response has a 4xx status code
func (o *RTRListPutFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list put files bad request response has a 5xx status code
func (o *RTRListPutFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list put files bad request response a status code equal to that given
func (o *RTRListPutFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r list put files bad request response
func (o *RTRListPutFilesBadRequest) Code() int {
	return 400
}

func (o *RTRListPutFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListPutFilesBadRequest) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListPutFilesBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListPutFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesForbidden creates a RTRListPutFilesForbidden with default headers values
func NewRTRListPutFilesForbidden() *RTRListPutFilesForbidden {
	return &RTRListPutFilesForbidden{}
}

/*
RTRListPutFilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListPutFilesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r list put files forbidden response has a 2xx status code
func (o *RTRListPutFilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list put files forbidden response has a 3xx status code
func (o *RTRListPutFilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list put files forbidden response has a 4xx status code
func (o *RTRListPutFilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list put files forbidden response has a 5xx status code
func (o *RTRListPutFilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list put files forbidden response a status code equal to that given
func (o *RTRListPutFilesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r list put files forbidden response
func (o *RTRListPutFilesForbidden) Code() int {
	return 403
}

func (o *RTRListPutFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRListPutFilesForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRListPutFilesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListPutFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListPutFilesNotFound creates a RTRListPutFilesNotFound with default headers values
func NewRTRListPutFilesNotFound() *RTRListPutFilesNotFound {
	return &RTRListPutFilesNotFound{}
}

/*
RTRListPutFilesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListPutFilesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list put files not found response has a 2xx status code
func (o *RTRListPutFilesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list put files not found response has a 3xx status code
func (o *RTRListPutFilesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list put files not found response has a 4xx status code
func (o *RTRListPutFilesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list put files not found response has a 5xx status code
func (o *RTRListPutFilesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list put files not found response a status code equal to that given
func (o *RTRListPutFilesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r list put files not found response
func (o *RTRListPutFilesNotFound) Code() int {
	return 404
}

func (o *RTRListPutFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRListPutFilesNotFound) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRListPutFilesNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListPutFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesTooManyRequests creates a RTRListPutFilesTooManyRequests with default headers values
func NewRTRListPutFilesTooManyRequests() *RTRListPutFilesTooManyRequests {
	return &RTRListPutFilesTooManyRequests{}
}

/*
RTRListPutFilesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListPutFilesTooManyRequests struct {

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

// IsSuccess returns true when this r t r list put files too many requests response has a 2xx status code
func (o *RTRListPutFilesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list put files too many requests response has a 3xx status code
func (o *RTRListPutFilesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list put files too many requests response has a 4xx status code
func (o *RTRListPutFilesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list put files too many requests response has a 5xx status code
func (o *RTRListPutFilesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list put files too many requests response a status code equal to that given
func (o *RTRListPutFilesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r list put files too many requests response
func (o *RTRListPutFilesTooManyRequests) Code() int {
	return 429
}

func (o *RTRListPutFilesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListPutFilesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListPutFilesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListPutFilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListPutFilesDefault creates a RTRListPutFilesDefault with default headers values
func NewRTRListPutFilesDefault(code int) *RTRListPutFilesDefault {
	return &RTRListPutFilesDefault{
		_statusCode: code,
	}
}

/*
RTRListPutFilesDefault describes a response with status code -1, with default header values.

OK
*/
type RTRListPutFilesDefault struct {
	_statusCode int

	Payload *models.BinservclientMsaPutFileResponse
}

// IsSuccess returns true when this r t r list put files default response has a 2xx status code
func (o *RTRListPutFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r list put files default response has a 3xx status code
func (o *RTRListPutFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r list put files default response has a 4xx status code
func (o *RTRListPutFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r list put files default response has a 5xx status code
func (o *RTRListPutFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r list put files default response a status code equal to that given
func (o *RTRListPutFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the r t r list put files default response
func (o *RTRListPutFilesDefault) Code() int {
	return o._statusCode
}

func (o *RTRListPutFilesDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] RTR-ListPut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRListPutFilesDefault) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] RTR-ListPut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRListPutFilesDefault) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListPutFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
