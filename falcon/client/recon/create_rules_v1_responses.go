// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// CreateRulesV1Reader is a Reader for the CreateRulesV1 structure.
type CreateRulesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRulesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRulesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRulesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRulesV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRulesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateRulesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRulesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRulesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRulesV1OK creates a CreateRulesV1OK with default headers values
func NewCreateRulesV1OK() *CreateRulesV1OK {
	return &CreateRulesV1OK{}
}

/*
CreateRulesV1OK describes a response with status code 200, with default header values.

OK
*/
type CreateRulesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRulesEntitiesResponseV1
}

// IsSuccess returns true when this create rules v1 o k response has a 2xx status code
func (o *CreateRulesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create rules v1 o k response has a 3xx status code
func (o *CreateRulesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 o k response has a 4xx status code
func (o *CreateRulesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create rules v1 o k response has a 5xx status code
func (o *CreateRulesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this create rules v1 o k response a status code equal to that given
func (o *CreateRulesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create rules v1 o k response
func (o *CreateRulesV1OK) Code() int {
	return 200
}

func (o *CreateRulesV1OK) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1OK  %+v", 200, o.Payload)
}

func (o *CreateRulesV1OK) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1OK  %+v", 200, o.Payload)
}

func (o *CreateRulesV1OK) GetPayload() *models.DomainRulesEntitiesResponseV1 {
	return o.Payload
}

func (o *CreateRulesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRulesEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRulesV1BadRequest creates a CreateRulesV1BadRequest with default headers values
func NewCreateRulesV1BadRequest() *CreateRulesV1BadRequest {
	return &CreateRulesV1BadRequest{}
}

/*
CreateRulesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateRulesV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this create rules v1 bad request response has a 2xx status code
func (o *CreateRulesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create rules v1 bad request response has a 3xx status code
func (o *CreateRulesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 bad request response has a 4xx status code
func (o *CreateRulesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create rules v1 bad request response has a 5xx status code
func (o *CreateRulesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create rules v1 bad request response a status code equal to that given
func (o *CreateRulesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create rules v1 bad request response
func (o *CreateRulesV1BadRequest) Code() int {
	return 400
}

func (o *CreateRulesV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateRulesV1BadRequest) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateRulesV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *CreateRulesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRulesV1Unauthorized creates a CreateRulesV1Unauthorized with default headers values
func NewCreateRulesV1Unauthorized() *CreateRulesV1Unauthorized {
	return &CreateRulesV1Unauthorized{}
}

/*
CreateRulesV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRulesV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this create rules v1 unauthorized response has a 2xx status code
func (o *CreateRulesV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create rules v1 unauthorized response has a 3xx status code
func (o *CreateRulesV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 unauthorized response has a 4xx status code
func (o *CreateRulesV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create rules v1 unauthorized response has a 5xx status code
func (o *CreateRulesV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create rules v1 unauthorized response a status code equal to that given
func (o *CreateRulesV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create rules v1 unauthorized response
func (o *CreateRulesV1Unauthorized) Code() int {
	return 401
}

func (o *CreateRulesV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CreateRulesV1Unauthorized) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CreateRulesV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *CreateRulesV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRulesV1Forbidden creates a CreateRulesV1Forbidden with default headers values
func NewCreateRulesV1Forbidden() *CreateRulesV1Forbidden {
	return &CreateRulesV1Forbidden{}
}

/*
CreateRulesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRulesV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this create rules v1 forbidden response has a 2xx status code
func (o *CreateRulesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create rules v1 forbidden response has a 3xx status code
func (o *CreateRulesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 forbidden response has a 4xx status code
func (o *CreateRulesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create rules v1 forbidden response has a 5xx status code
func (o *CreateRulesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create rules v1 forbidden response a status code equal to that given
func (o *CreateRulesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create rules v1 forbidden response
func (o *CreateRulesV1Forbidden) Code() int {
	return 403
}

func (o *CreateRulesV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateRulesV1Forbidden) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateRulesV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *CreateRulesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRulesV1TooManyRequests creates a CreateRulesV1TooManyRequests with default headers values
func NewCreateRulesV1TooManyRequests() *CreateRulesV1TooManyRequests {
	return &CreateRulesV1TooManyRequests{}
}

/*
CreateRulesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateRulesV1TooManyRequests struct {

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

// IsSuccess returns true when this create rules v1 too many requests response has a 2xx status code
func (o *CreateRulesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create rules v1 too many requests response has a 3xx status code
func (o *CreateRulesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 too many requests response has a 4xx status code
func (o *CreateRulesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create rules v1 too many requests response has a 5xx status code
func (o *CreateRulesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create rules v1 too many requests response a status code equal to that given
func (o *CreateRulesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create rules v1 too many requests response
func (o *CreateRulesV1TooManyRequests) Code() int {
	return 429
}

func (o *CreateRulesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateRulesV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateRulesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateRulesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRulesV1InternalServerError creates a CreateRulesV1InternalServerError with default headers values
func NewCreateRulesV1InternalServerError() *CreateRulesV1InternalServerError {
	return &CreateRulesV1InternalServerError{}
}

/*
CreateRulesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateRulesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this create rules v1 internal server error response has a 2xx status code
func (o *CreateRulesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create rules v1 internal server error response has a 3xx status code
func (o *CreateRulesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create rules v1 internal server error response has a 4xx status code
func (o *CreateRulesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create rules v1 internal server error response has a 5xx status code
func (o *CreateRulesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create rules v1 internal server error response a status code equal to that given
func (o *CreateRulesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create rules v1 internal server error response
func (o *CreateRulesV1InternalServerError) Code() int {
	return 500
}

func (o *CreateRulesV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRulesV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] createRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRulesV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *CreateRulesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRulesV1Default creates a CreateRulesV1Default with default headers values
func NewCreateRulesV1Default(code int) *CreateRulesV1Default {
	return &CreateRulesV1Default{
		_statusCode: code,
	}
}

/*
CreateRulesV1Default describes a response with status code -1, with default header values.

OK
*/
type CreateRulesV1Default struct {
	_statusCode int

	Payload *models.DomainRulesEntitiesResponseV1
}

// IsSuccess returns true when this create rules v1 default response has a 2xx status code
func (o *CreateRulesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create rules v1 default response has a 3xx status code
func (o *CreateRulesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create rules v1 default response has a 4xx status code
func (o *CreateRulesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create rules v1 default response has a 5xx status code
func (o *CreateRulesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create rules v1 default response a status code equal to that given
func (o *CreateRulesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create rules v1 default response
func (o *CreateRulesV1Default) Code() int {
	return o._statusCode
}

func (o *CreateRulesV1Default) Error() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] CreateRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRulesV1Default) String() string {
	return fmt.Sprintf("[POST /recon/entities/rules/v1][%d] CreateRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRulesV1Default) GetPayload() *models.DomainRulesEntitiesResponseV1 {
	return o.Payload
}

func (o *CreateRulesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainRulesEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
