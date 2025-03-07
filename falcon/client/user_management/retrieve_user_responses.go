// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// RetrieveUserReader is a Reader for the RetrieveUser structure.
type RetrieveUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrieveUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrieveUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRetrieveUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRetrieveUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetrieveUserOK creates a RetrieveUserOK with default headers values
func NewRetrieveUserOK() *RetrieveUserOK {
	return &RetrieveUserOK{}
}

/*
RetrieveUserOK describes a response with status code 200, with default header values.

OK
*/
type RetrieveUserOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIUserMetadataResponse
}

// IsSuccess returns true when this retrieve user o k response has a 2xx status code
func (o *RetrieveUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve user o k response has a 3xx status code
func (o *RetrieveUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user o k response has a 4xx status code
func (o *RetrieveUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve user o k response has a 5xx status code
func (o *RetrieveUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user o k response a status code equal to that given
func (o *RetrieveUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retrieve user o k response
func (o *RetrieveUserOK) Code() int {
	return 200
}

func (o *RetrieveUserOK) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserOK  %+v", 200, o.Payload)
}

func (o *RetrieveUserOK) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserOK  %+v", 200, o.Payload)
}

func (o *RetrieveUserOK) GetPayload() *models.APIUserMetadataResponse {
	return o.Payload
}

func (o *RetrieveUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIUserMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserBadRequest creates a RetrieveUserBadRequest with default headers values
func NewRetrieveUserBadRequest() *RetrieveUserBadRequest {
	return &RetrieveUserBadRequest{}
}

/*
RetrieveUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrieveUserBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this retrieve user bad request response has a 2xx status code
func (o *RetrieveUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user bad request response has a 3xx status code
func (o *RetrieveUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user bad request response has a 4xx status code
func (o *RetrieveUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user bad request response has a 5xx status code
func (o *RetrieveUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user bad request response a status code equal to that given
func (o *RetrieveUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the retrieve user bad request response
func (o *RetrieveUserBadRequest) Code() int {
	return 400
}

func (o *RetrieveUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveUserBadRequest) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveUserBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *RetrieveUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserForbidden creates a RetrieveUserForbidden with default headers values
func NewRetrieveUserForbidden() *RetrieveUserForbidden {
	return &RetrieveUserForbidden{}
}

/*
RetrieveUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RetrieveUserForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this retrieve user forbidden response has a 2xx status code
func (o *RetrieveUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user forbidden response has a 3xx status code
func (o *RetrieveUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user forbidden response has a 4xx status code
func (o *RetrieveUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user forbidden response has a 5xx status code
func (o *RetrieveUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user forbidden response a status code equal to that given
func (o *RetrieveUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the retrieve user forbidden response
func (o *RetrieveUserForbidden) Code() int {
	return 403
}

func (o *RetrieveUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserForbidden) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *RetrieveUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserNotFound creates a RetrieveUserNotFound with default headers values
func NewRetrieveUserNotFound() *RetrieveUserNotFound {
	return &RetrieveUserNotFound{}
}

/*
RetrieveUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RetrieveUserNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this retrieve user not found response has a 2xx status code
func (o *RetrieveUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user not found response has a 3xx status code
func (o *RetrieveUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user not found response has a 4xx status code
func (o *RetrieveUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user not found response has a 5xx status code
func (o *RetrieveUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user not found response a status code equal to that given
func (o *RetrieveUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the retrieve user not found response
func (o *RetrieveUserNotFound) Code() int {
	return 404
}

func (o *RetrieveUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveUserNotFound) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveUserNotFound) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *RetrieveUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserTooManyRequests creates a RetrieveUserTooManyRequests with default headers values
func NewRetrieveUserTooManyRequests() *RetrieveUserTooManyRequests {
	return &RetrieveUserTooManyRequests{}
}

/*
RetrieveUserTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RetrieveUserTooManyRequests struct {

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

// IsSuccess returns true when this retrieve user too many requests response has a 2xx status code
func (o *RetrieveUserTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user too many requests response has a 3xx status code
func (o *RetrieveUserTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user too many requests response has a 4xx status code
func (o *RetrieveUserTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user too many requests response has a 5xx status code
func (o *RetrieveUserTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user too many requests response a status code equal to that given
func (o *RetrieveUserTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the retrieve user too many requests response
func (o *RetrieveUserTooManyRequests) Code() int {
	return 429
}

func (o *RetrieveUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *RetrieveUserTooManyRequests) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *RetrieveUserTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RetrieveUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRetrieveUserDefault creates a RetrieveUserDefault with default headers values
func NewRetrieveUserDefault(code int) *RetrieveUserDefault {
	return &RetrieveUserDefault{
		_statusCode: code,
	}
}

/*
RetrieveUserDefault describes a response with status code -1, with default header values.

OK
*/
type RetrieveUserDefault struct {
	_statusCode int

	Payload *models.APIUserMetadataResponse
}

// IsSuccess returns true when this retrieve user default response has a 2xx status code
func (o *RetrieveUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this retrieve user default response has a 3xx status code
func (o *RetrieveUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this retrieve user default response has a 4xx status code
func (o *RetrieveUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this retrieve user default response has a 5xx status code
func (o *RetrieveUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this retrieve user default response a status code equal to that given
func (o *RetrieveUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the retrieve user default response
func (o *RetrieveUserDefault) Code() int {
	return o._statusCode
}

func (o *RetrieveUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUser default  %+v", o._statusCode, o.Payload)
}

func (o *RetrieveUserDefault) String() string {
	return fmt.Sprintf("[GET /users/entities/users/v1][%d] retrieveUser default  %+v", o._statusCode, o.Payload)
}

func (o *RetrieveUserDefault) GetPayload() *models.APIUserMetadataResponse {
	return o.Payload
}

func (o *RetrieveUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
