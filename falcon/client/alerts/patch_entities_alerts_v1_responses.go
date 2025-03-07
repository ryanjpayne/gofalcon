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

// PatchEntitiesAlertsV1Reader is a Reader for the PatchEntitiesAlertsV1 structure.
type PatchEntitiesAlertsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEntitiesAlertsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEntitiesAlertsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchEntitiesAlertsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEntitiesAlertsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchEntitiesAlertsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEntitiesAlertsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEntitiesAlertsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEntitiesAlertsV1OK creates a PatchEntitiesAlertsV1OK with default headers values
func NewPatchEntitiesAlertsV1OK() *PatchEntitiesAlertsV1OK {
	return &PatchEntitiesAlertsV1OK{}
}

/*
PatchEntitiesAlertsV1OK describes a response with status code 200, with default header values.

OK
*/
type PatchEntitiesAlertsV1OK struct {

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

// IsSuccess returns true when this patch entities alerts v1 o k response has a 2xx status code
func (o *PatchEntitiesAlertsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch entities alerts v1 o k response has a 3xx status code
func (o *PatchEntitiesAlertsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v1 o k response has a 4xx status code
func (o *PatchEntitiesAlertsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch entities alerts v1 o k response has a 5xx status code
func (o *PatchEntitiesAlertsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v1 o k response a status code equal to that given
func (o *PatchEntitiesAlertsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch entities alerts v1 o k response
func (o *PatchEntitiesAlertsV1OK) Code() int {
	return 200
}

func (o *PatchEntitiesAlertsV1OK) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *PatchEntitiesAlertsV1OK) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *PatchEntitiesAlertsV1OK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV1BadRequest creates a PatchEntitiesAlertsV1BadRequest with default headers values
func NewPatchEntitiesAlertsV1BadRequest() *PatchEntitiesAlertsV1BadRequest {
	return &PatchEntitiesAlertsV1BadRequest{}
}

/*
PatchEntitiesAlertsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchEntitiesAlertsV1BadRequest struct {

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

// IsSuccess returns true when this patch entities alerts v1 bad request response has a 2xx status code
func (o *PatchEntitiesAlertsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v1 bad request response has a 3xx status code
func (o *PatchEntitiesAlertsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v1 bad request response has a 4xx status code
func (o *PatchEntitiesAlertsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v1 bad request response has a 5xx status code
func (o *PatchEntitiesAlertsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v1 bad request response a status code equal to that given
func (o *PatchEntitiesAlertsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch entities alerts v1 bad request response
func (o *PatchEntitiesAlertsV1BadRequest) Code() int {
	return 400
}

func (o *PatchEntitiesAlertsV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *PatchEntitiesAlertsV1BadRequest) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *PatchEntitiesAlertsV1BadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV1Forbidden creates a PatchEntitiesAlertsV1Forbidden with default headers values
func NewPatchEntitiesAlertsV1Forbidden() *PatchEntitiesAlertsV1Forbidden {
	return &PatchEntitiesAlertsV1Forbidden{}
}

/*
PatchEntitiesAlertsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchEntitiesAlertsV1Forbidden struct {

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

// IsSuccess returns true when this patch entities alerts v1 forbidden response has a 2xx status code
func (o *PatchEntitiesAlertsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v1 forbidden response has a 3xx status code
func (o *PatchEntitiesAlertsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v1 forbidden response has a 4xx status code
func (o *PatchEntitiesAlertsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v1 forbidden response has a 5xx status code
func (o *PatchEntitiesAlertsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v1 forbidden response a status code equal to that given
func (o *PatchEntitiesAlertsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch entities alerts v1 forbidden response
func (o *PatchEntitiesAlertsV1Forbidden) Code() int {
	return 403
}

func (o *PatchEntitiesAlertsV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *PatchEntitiesAlertsV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *PatchEntitiesAlertsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV1TooManyRequests creates a PatchEntitiesAlertsV1TooManyRequests with default headers values
func NewPatchEntitiesAlertsV1TooManyRequests() *PatchEntitiesAlertsV1TooManyRequests {
	return &PatchEntitiesAlertsV1TooManyRequests{}
}

/*
PatchEntitiesAlertsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchEntitiesAlertsV1TooManyRequests struct {

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

// IsSuccess returns true when this patch entities alerts v1 too many requests response has a 2xx status code
func (o *PatchEntitiesAlertsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v1 too many requests response has a 3xx status code
func (o *PatchEntitiesAlertsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v1 too many requests response has a 4xx status code
func (o *PatchEntitiesAlertsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v1 too many requests response has a 5xx status code
func (o *PatchEntitiesAlertsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v1 too many requests response a status code equal to that given
func (o *PatchEntitiesAlertsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch entities alerts v1 too many requests response
func (o *PatchEntitiesAlertsV1TooManyRequests) Code() int {
	return 429
}

func (o *PatchEntitiesAlertsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchEntitiesAlertsV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchEntitiesAlertsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV1InternalServerError creates a PatchEntitiesAlertsV1InternalServerError with default headers values
func NewPatchEntitiesAlertsV1InternalServerError() *PatchEntitiesAlertsV1InternalServerError {
	return &PatchEntitiesAlertsV1InternalServerError{}
}

/*
PatchEntitiesAlertsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchEntitiesAlertsV1InternalServerError struct {

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

// IsSuccess returns true when this patch entities alerts v1 internal server error response has a 2xx status code
func (o *PatchEntitiesAlertsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v1 internal server error response has a 3xx status code
func (o *PatchEntitiesAlertsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v1 internal server error response has a 4xx status code
func (o *PatchEntitiesAlertsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch entities alerts v1 internal server error response has a 5xx status code
func (o *PatchEntitiesAlertsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch entities alerts v1 internal server error response a status code equal to that given
func (o *PatchEntitiesAlertsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch entities alerts v1 internal server error response
func (o *PatchEntitiesAlertsV1InternalServerError) Code() int {
	return 500
}

func (o *PatchEntitiesAlertsV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEntitiesAlertsV1InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] patchEntitiesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEntitiesAlertsV1InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV1Default creates a PatchEntitiesAlertsV1Default with default headers values
func NewPatchEntitiesAlertsV1Default(code int) *PatchEntitiesAlertsV1Default {
	return &PatchEntitiesAlertsV1Default{
		_statusCode: code,
	}
}

/*
PatchEntitiesAlertsV1Default describes a response with status code -1, with default header values.

OK
*/
type PatchEntitiesAlertsV1Default struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this patch entities alerts v1 default response has a 2xx status code
func (o *PatchEntitiesAlertsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch entities alerts v1 default response has a 3xx status code
func (o *PatchEntitiesAlertsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch entities alerts v1 default response has a 4xx status code
func (o *PatchEntitiesAlertsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch entities alerts v1 default response has a 5xx status code
func (o *PatchEntitiesAlertsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch entities alerts v1 default response a status code equal to that given
func (o *PatchEntitiesAlertsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch entities alerts v1 default response
func (o *PatchEntitiesAlertsV1Default) Code() int {
	return o._statusCode
}

func (o *PatchEntitiesAlertsV1Default) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] PatchEntitiesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEntitiesAlertsV1Default) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v1][%d] PatchEntitiesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEntitiesAlertsV1Default) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
