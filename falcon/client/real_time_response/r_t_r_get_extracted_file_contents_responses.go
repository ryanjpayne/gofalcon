// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRGetExtractedFileContentsReader is a Reader for the RTRGetExtractedFileContents structure.
type RTRGetExtractedFileContentsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *RTRGetExtractedFileContentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRGetExtractedFileContentsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRGetExtractedFileContentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRGetExtractedFileContentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRGetExtractedFileContentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRGetExtractedFileContentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRGetExtractedFileContentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRGetExtractedFileContentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRGetExtractedFileContentsOK creates a RTRGetExtractedFileContentsOK with default headers values
func NewRTRGetExtractedFileContentsOK(writer io.Writer) *RTRGetExtractedFileContentsOK {
	return &RTRGetExtractedFileContentsOK{

		Payload: writer,
	}
}

/*
RTRGetExtractedFileContentsOK describes a response with status code 200, with default header values.

OK
*/
type RTRGetExtractedFileContentsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload io.Writer
}

// IsSuccess returns true when this r t r get extracted file contents o k response has a 2xx status code
func (o *RTRGetExtractedFileContentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r get extracted file contents o k response has a 3xx status code
func (o *RTRGetExtractedFileContentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents o k response has a 4xx status code
func (o *RTRGetExtractedFileContentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r get extracted file contents o k response has a 5xx status code
func (o *RTRGetExtractedFileContentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r get extracted file contents o k response a status code equal to that given
func (o *RTRGetExtractedFileContentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r get extracted file contents o k response
func (o *RTRGetExtractedFileContentsOK) Code() int {
	return 200
}

func (o *RTRGetExtractedFileContentsOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsOK  %+v", 200, o.Payload)
}

func (o *RTRGetExtractedFileContentsOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsOK  %+v", 200, o.Payload)
}

func (o *RTRGetExtractedFileContentsOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRGetExtractedFileContentsBadRequest creates a RTRGetExtractedFileContentsBadRequest with default headers values
func NewRTRGetExtractedFileContentsBadRequest() *RTRGetExtractedFileContentsBadRequest {
	return &RTRGetExtractedFileContentsBadRequest{}
}

/*
RTRGetExtractedFileContentsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRGetExtractedFileContentsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r get extracted file contents bad request response has a 2xx status code
func (o *RTRGetExtractedFileContentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r get extracted file contents bad request response has a 3xx status code
func (o *RTRGetExtractedFileContentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents bad request response has a 4xx status code
func (o *RTRGetExtractedFileContentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r get extracted file contents bad request response has a 5xx status code
func (o *RTRGetExtractedFileContentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r get extracted file contents bad request response a status code equal to that given
func (o *RTRGetExtractedFileContentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r get extracted file contents bad request response
func (o *RTRGetExtractedFileContentsBadRequest) Code() int {
	return 400
}

func (o *RTRGetExtractedFileContentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRGetExtractedFileContentsBadRequest) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRGetExtractedFileContentsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetExtractedFileContentsForbidden creates a RTRGetExtractedFileContentsForbidden with default headers values
func NewRTRGetExtractedFileContentsForbidden() *RTRGetExtractedFileContentsForbidden {
	return &RTRGetExtractedFileContentsForbidden{}
}

/*
RTRGetExtractedFileContentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRGetExtractedFileContentsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r get extracted file contents forbidden response has a 2xx status code
func (o *RTRGetExtractedFileContentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r get extracted file contents forbidden response has a 3xx status code
func (o *RTRGetExtractedFileContentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents forbidden response has a 4xx status code
func (o *RTRGetExtractedFileContentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r get extracted file contents forbidden response has a 5xx status code
func (o *RTRGetExtractedFileContentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r get extracted file contents forbidden response a status code equal to that given
func (o *RTRGetExtractedFileContentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r get extracted file contents forbidden response
func (o *RTRGetExtractedFileContentsForbidden) Code() int {
	return 403
}

func (o *RTRGetExtractedFileContentsForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsForbidden  %+v", 403, o.Payload)
}

func (o *RTRGetExtractedFileContentsForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsForbidden  %+v", 403, o.Payload)
}

func (o *RTRGetExtractedFileContentsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetExtractedFileContentsNotFound creates a RTRGetExtractedFileContentsNotFound with default headers values
func NewRTRGetExtractedFileContentsNotFound() *RTRGetExtractedFileContentsNotFound {
	return &RTRGetExtractedFileContentsNotFound{}
}

/*
RTRGetExtractedFileContentsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRGetExtractedFileContentsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r get extracted file contents not found response has a 2xx status code
func (o *RTRGetExtractedFileContentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r get extracted file contents not found response has a 3xx status code
func (o *RTRGetExtractedFileContentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents not found response has a 4xx status code
func (o *RTRGetExtractedFileContentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r get extracted file contents not found response has a 5xx status code
func (o *RTRGetExtractedFileContentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r get extracted file contents not found response a status code equal to that given
func (o *RTRGetExtractedFileContentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r get extracted file contents not found response
func (o *RTRGetExtractedFileContentsNotFound) Code() int {
	return 404
}

func (o *RTRGetExtractedFileContentsNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsNotFound  %+v", 404, o.Payload)
}

func (o *RTRGetExtractedFileContentsNotFound) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsNotFound  %+v", 404, o.Payload)
}

func (o *RTRGetExtractedFileContentsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetExtractedFileContentsTooManyRequests creates a RTRGetExtractedFileContentsTooManyRequests with default headers values
func NewRTRGetExtractedFileContentsTooManyRequests() *RTRGetExtractedFileContentsTooManyRequests {
	return &RTRGetExtractedFileContentsTooManyRequests{}
}

/*
RTRGetExtractedFileContentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRGetExtractedFileContentsTooManyRequests struct {

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

// IsSuccess returns true when this r t r get extracted file contents too many requests response has a 2xx status code
func (o *RTRGetExtractedFileContentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r get extracted file contents too many requests response has a 3xx status code
func (o *RTRGetExtractedFileContentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents too many requests response has a 4xx status code
func (o *RTRGetExtractedFileContentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r get extracted file contents too many requests response has a 5xx status code
func (o *RTRGetExtractedFileContentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r get extracted file contents too many requests response a status code equal to that given
func (o *RTRGetExtractedFileContentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r get extracted file contents too many requests response
func (o *RTRGetExtractedFileContentsTooManyRequests) Code() int {
	return 429
}

func (o *RTRGetExtractedFileContentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRGetExtractedFileContentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRGetExtractedFileContentsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetExtractedFileContentsInternalServerError creates a RTRGetExtractedFileContentsInternalServerError with default headers values
func NewRTRGetExtractedFileContentsInternalServerError() *RTRGetExtractedFileContentsInternalServerError {
	return &RTRGetExtractedFileContentsInternalServerError{}
}

/*
RTRGetExtractedFileContentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RTRGetExtractedFileContentsInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r get extracted file contents internal server error response has a 2xx status code
func (o *RTRGetExtractedFileContentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r get extracted file contents internal server error response has a 3xx status code
func (o *RTRGetExtractedFileContentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r get extracted file contents internal server error response has a 4xx status code
func (o *RTRGetExtractedFileContentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r get extracted file contents internal server error response has a 5xx status code
func (o *RTRGetExtractedFileContentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r get extracted file contents internal server error response a status code equal to that given
func (o *RTRGetExtractedFileContentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r get extracted file contents internal server error response
func (o *RTRGetExtractedFileContentsInternalServerError) Code() int {
	return 500
}

func (o *RTRGetExtractedFileContentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRGetExtractedFileContentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] rTRGetExtractedFileContentsInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRGetExtractedFileContentsInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetExtractedFileContentsDefault creates a RTRGetExtractedFileContentsDefault with default headers values
func NewRTRGetExtractedFileContentsDefault(code int) *RTRGetExtractedFileContentsDefault {
	return &RTRGetExtractedFileContentsDefault{
		_statusCode: code,
	}
}

/*
RTRGetExtractedFileContentsDefault describes a response with status code -1, with default header values.

OK
*/
type RTRGetExtractedFileContentsDefault struct {
	_statusCode int

	Payload []int64
}

// IsSuccess returns true when this r t r get extracted file contents default response has a 2xx status code
func (o *RTRGetExtractedFileContentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r get extracted file contents default response has a 3xx status code
func (o *RTRGetExtractedFileContentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r get extracted file contents default response has a 4xx status code
func (o *RTRGetExtractedFileContentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r get extracted file contents default response has a 5xx status code
func (o *RTRGetExtractedFileContentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r get extracted file contents default response a status code equal to that given
func (o *RTRGetExtractedFileContentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the r t r get extracted file contents default response
func (o *RTRGetExtractedFileContentsDefault) Code() int {
	return o._statusCode
}

func (o *RTRGetExtractedFileContentsDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] RTR-GetExtractedFileContents default  %+v", o._statusCode, o.Payload)
}

func (o *RTRGetExtractedFileContentsDefault) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/extracted-file-contents/v1][%d] RTR-GetExtractedFileContents default  %+v", o._statusCode, o.Payload)
}

func (o *RTRGetExtractedFileContentsDefault) GetPayload() []int64 {
	return o.Payload
}

func (o *RTRGetExtractedFileContentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
