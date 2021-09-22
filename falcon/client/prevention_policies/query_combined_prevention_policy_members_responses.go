// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// QueryCombinedPreventionPolicyMembersReader is a Reader for the QueryCombinedPreventionPolicyMembers structure.
type QueryCombinedPreventionPolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedPreventionPolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedPreventionPolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedPreventionPolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedPreventionPolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedPreventionPolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedPreventionPolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedPreventionPolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedPreventionPolicyMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedPreventionPolicyMembersOK creates a QueryCombinedPreventionPolicyMembersOK with default headers values
func NewQueryCombinedPreventionPolicyMembersOK() *QueryCombinedPreventionPolicyMembersOK {
	return &QueryCombinedPreventionPolicyMembersOK{}
}

/* QueryCombinedPreventionPolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedPreventionPolicyMembersOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedPreventionPolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersOK  %+v", 200, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersOK) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPolicyMembersBadRequest creates a QueryCombinedPreventionPolicyMembersBadRequest with default headers values
func NewQueryCombinedPreventionPolicyMembersBadRequest() *QueryCombinedPreventionPolicyMembersBadRequest {
	return &QueryCombinedPreventionPolicyMembersBadRequest{}
}

/* QueryCombinedPreventionPolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedPreventionPolicyMembersBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedPreventionPolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersBadRequest) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPolicyMembersForbidden creates a QueryCombinedPreventionPolicyMembersForbidden with default headers values
func NewQueryCombinedPreventionPolicyMembersForbidden() *QueryCombinedPreventionPolicyMembersForbidden {
	return &QueryCombinedPreventionPolicyMembersForbidden{}
}

/* QueryCombinedPreventionPolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedPreventionPolicyMembersForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryCombinedPreventionPolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersForbidden  %+v", 403, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPolicyMembersNotFound creates a QueryCombinedPreventionPolicyMembersNotFound with default headers values
func NewQueryCombinedPreventionPolicyMembersNotFound() *QueryCombinedPreventionPolicyMembersNotFound {
	return &QueryCombinedPreventionPolicyMembersNotFound{}
}

/* QueryCombinedPreventionPolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedPreventionPolicyMembersNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedPreventionPolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersNotFound  %+v", 404, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersNotFound) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPolicyMembersTooManyRequests creates a QueryCombinedPreventionPolicyMembersTooManyRequests with default headers values
func NewQueryCombinedPreventionPolicyMembersTooManyRequests() *QueryCombinedPreventionPolicyMembersTooManyRequests {
	return &QueryCombinedPreventionPolicyMembersTooManyRequests{}
}

/* QueryCombinedPreventionPolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedPreventionPolicyMembersTooManyRequests struct {

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

func (o *QueryCombinedPreventionPolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedPreventionPolicyMembersInternalServerError creates a QueryCombinedPreventionPolicyMembersInternalServerError with default headers values
func NewQueryCombinedPreventionPolicyMembersInternalServerError() *QueryCombinedPreventionPolicyMembersInternalServerError {
	return &QueryCombinedPreventionPolicyMembersInternalServerError{}
}

/* QueryCombinedPreventionPolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedPreventionPolicyMembersInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedPreventionPolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembersInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersInternalServerError) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPolicyMembersDefault creates a QueryCombinedPreventionPolicyMembersDefault with default headers values
func NewQueryCombinedPreventionPolicyMembersDefault(code int) *QueryCombinedPreventionPolicyMembersDefault {
	return &QueryCombinedPreventionPolicyMembersDefault{
		_statusCode: code,
	}
}

/* QueryCombinedPreventionPolicyMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedPreventionPolicyMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesPolicyMembersRespV1
}

// Code gets the status code for the query combined prevention policy members default response
func (o *QueryCombinedPreventionPolicyMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedPreventionPolicyMembersDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention-members/v1][%d] queryCombinedPreventionPolicyMembers default  %+v", o._statusCode, o.Payload)
}
func (o *QueryCombinedPreventionPolicyMembersDefault) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPolicyMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
