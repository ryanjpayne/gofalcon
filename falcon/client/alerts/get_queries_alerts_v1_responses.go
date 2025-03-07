// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// GetQueriesAlertsV1Reader is a Reader for the GetQueriesAlertsV1 structure.
type GetQueriesAlertsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueriesAlertsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueriesAlertsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQueriesAlertsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQueriesAlertsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQueriesAlertsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQueriesAlertsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetQueriesAlertsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueriesAlertsV1OK creates a GetQueriesAlertsV1OK with default headers values
func NewGetQueriesAlertsV1OK() *GetQueriesAlertsV1OK {
	return &GetQueriesAlertsV1OK{}
}

/*
GetQueriesAlertsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetQueriesAlertsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get queries alerts v1 o k response has a 2xx status code
func (o *GetQueriesAlertsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get queries alerts v1 o k response has a 3xx status code
func (o *GetQueriesAlertsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries alerts v1 o k response has a 4xx status code
func (o *GetQueriesAlertsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queries alerts v1 o k response has a 5xx status code
func (o *GetQueriesAlertsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries alerts v1 o k response a status code equal to that given
func (o *GetQueriesAlertsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get queries alerts v1 o k response
func (o *GetQueriesAlertsV1OK) Code() int {
	return 200
}

func (o *GetQueriesAlertsV1OK) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *GetQueriesAlertsV1OK) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *GetQueriesAlertsV1OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetQueriesAlertsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesAlertsV1BadRequest creates a GetQueriesAlertsV1BadRequest with default headers values
func NewGetQueriesAlertsV1BadRequest() *GetQueriesAlertsV1BadRequest {
	return &GetQueriesAlertsV1BadRequest{}
}

/*
GetQueriesAlertsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetQueriesAlertsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get queries alerts v1 bad request response has a 2xx status code
func (o *GetQueriesAlertsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queries alerts v1 bad request response has a 3xx status code
func (o *GetQueriesAlertsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries alerts v1 bad request response has a 4xx status code
func (o *GetQueriesAlertsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get queries alerts v1 bad request response has a 5xx status code
func (o *GetQueriesAlertsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries alerts v1 bad request response a status code equal to that given
func (o *GetQueriesAlertsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get queries alerts v1 bad request response
func (o *GetQueriesAlertsV1BadRequest) Code() int {
	return 400
}

func (o *GetQueriesAlertsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetQueriesAlertsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetQueriesAlertsV1BadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetQueriesAlertsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesAlertsV1Forbidden creates a GetQueriesAlertsV1Forbidden with default headers values
func NewGetQueriesAlertsV1Forbidden() *GetQueriesAlertsV1Forbidden {
	return &GetQueriesAlertsV1Forbidden{}
}

/*
GetQueriesAlertsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetQueriesAlertsV1Forbidden struct {

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

// IsSuccess returns true when this get queries alerts v1 forbidden response has a 2xx status code
func (o *GetQueriesAlertsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queries alerts v1 forbidden response has a 3xx status code
func (o *GetQueriesAlertsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries alerts v1 forbidden response has a 4xx status code
func (o *GetQueriesAlertsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get queries alerts v1 forbidden response has a 5xx status code
func (o *GetQueriesAlertsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries alerts v1 forbidden response a status code equal to that given
func (o *GetQueriesAlertsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get queries alerts v1 forbidden response
func (o *GetQueriesAlertsV1Forbidden) Code() int {
	return 403
}

func (o *GetQueriesAlertsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetQueriesAlertsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetQueriesAlertsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetQueriesAlertsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetQueriesAlertsV1TooManyRequests creates a GetQueriesAlertsV1TooManyRequests with default headers values
func NewGetQueriesAlertsV1TooManyRequests() *GetQueriesAlertsV1TooManyRequests {
	return &GetQueriesAlertsV1TooManyRequests{}
}

/*
GetQueriesAlertsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetQueriesAlertsV1TooManyRequests struct {

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

// IsSuccess returns true when this get queries alerts v1 too many requests response has a 2xx status code
func (o *GetQueriesAlertsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queries alerts v1 too many requests response has a 3xx status code
func (o *GetQueriesAlertsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries alerts v1 too many requests response has a 4xx status code
func (o *GetQueriesAlertsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get queries alerts v1 too many requests response has a 5xx status code
func (o *GetQueriesAlertsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get queries alerts v1 too many requests response a status code equal to that given
func (o *GetQueriesAlertsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get queries alerts v1 too many requests response
func (o *GetQueriesAlertsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetQueriesAlertsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQueriesAlertsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQueriesAlertsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetQueriesAlertsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetQueriesAlertsV1InternalServerError creates a GetQueriesAlertsV1InternalServerError with default headers values
func NewGetQueriesAlertsV1InternalServerError() *GetQueriesAlertsV1InternalServerError {
	return &GetQueriesAlertsV1InternalServerError{}
}

/*
GetQueriesAlertsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetQueriesAlertsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get queries alerts v1 internal server error response has a 2xx status code
func (o *GetQueriesAlertsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queries alerts v1 internal server error response has a 3xx status code
func (o *GetQueriesAlertsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queries alerts v1 internal server error response has a 4xx status code
func (o *GetQueriesAlertsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queries alerts v1 internal server error response has a 5xx status code
func (o *GetQueriesAlertsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get queries alerts v1 internal server error response a status code equal to that given
func (o *GetQueriesAlertsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get queries alerts v1 internal server error response
func (o *GetQueriesAlertsV1InternalServerError) Code() int {
	return 500
}

func (o *GetQueriesAlertsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetQueriesAlertsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] getQueriesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetQueriesAlertsV1InternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetQueriesAlertsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesAlertsV1Default creates a GetQueriesAlertsV1Default with default headers values
func NewGetQueriesAlertsV1Default(code int) *GetQueriesAlertsV1Default {
	return &GetQueriesAlertsV1Default{
		_statusCode: code,
	}
}

/*
GetQueriesAlertsV1Default describes a response with status code -1, with default header values.

OK
*/
type GetQueriesAlertsV1Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get queries alerts v1 default response has a 2xx status code
func (o *GetQueriesAlertsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get queries alerts v1 default response has a 3xx status code
func (o *GetQueriesAlertsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get queries alerts v1 default response has a 4xx status code
func (o *GetQueriesAlertsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get queries alerts v1 default response has a 5xx status code
func (o *GetQueriesAlertsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get queries alerts v1 default response a status code equal to that given
func (o *GetQueriesAlertsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get queries alerts v1 default response
func (o *GetQueriesAlertsV1Default) Code() int {
	return o._statusCode
}

func (o *GetQueriesAlertsV1Default) Error() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] GetQueriesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesAlertsV1Default) String() string {
	return fmt.Sprintf("[GET /alerts/queries/alerts/v1][%d] GetQueriesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueriesAlertsV1Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetQueriesAlertsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
