// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// QuerySensorUpdatePoliciesReader is a Reader for the QuerySensorUpdatePolicies structure.
type QuerySensorUpdatePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySensorUpdatePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySensorUpdatePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuerySensorUpdatePoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuerySensorUpdatePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQuerySensorUpdatePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuerySensorUpdatePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQuerySensorUpdatePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuerySensorUpdatePoliciesOK creates a QuerySensorUpdatePoliciesOK with default headers values
func NewQuerySensorUpdatePoliciesOK() *QuerySensorUpdatePoliciesOK {
	return &QuerySensorUpdatePoliciesOK{}
}

/*
QuerySensorUpdatePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QuerySensorUpdatePoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query sensor update policies o k response has a 2xx status code
func (o *QuerySensorUpdatePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query sensor update policies o k response has a 3xx status code
func (o *QuerySensorUpdatePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policies o k response has a 4xx status code
func (o *QuerySensorUpdatePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensor update policies o k response has a 5xx status code
func (o *QuerySensorUpdatePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policies o k response a status code equal to that given
func (o *QuerySensorUpdatePoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query sensor update policies o k response
func (o *QuerySensorUpdatePoliciesOK) Code() int {
	return 200
}

func (o *QuerySensorUpdatePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesOK  %+v", 200, o.Payload)
}

func (o *QuerySensorUpdatePoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesOK  %+v", 200, o.Payload)
}

func (o *QuerySensorUpdatePoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePoliciesBadRequest creates a QuerySensorUpdatePoliciesBadRequest with default headers values
func NewQuerySensorUpdatePoliciesBadRequest() *QuerySensorUpdatePoliciesBadRequest {
	return &QuerySensorUpdatePoliciesBadRequest{}
}

/*
QuerySensorUpdatePoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QuerySensorUpdatePoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query sensor update policies bad request response has a 2xx status code
func (o *QuerySensorUpdatePoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policies bad request response has a 3xx status code
func (o *QuerySensorUpdatePoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policies bad request response has a 4xx status code
func (o *QuerySensorUpdatePoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policies bad request response has a 5xx status code
func (o *QuerySensorUpdatePoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policies bad request response a status code equal to that given
func (o *QuerySensorUpdatePoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query sensor update policies bad request response
func (o *QuerySensorUpdatePoliciesBadRequest) Code() int {
	return 400
}

func (o *QuerySensorUpdatePoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySensorUpdatePoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySensorUpdatePoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePoliciesForbidden creates a QuerySensorUpdatePoliciesForbidden with default headers values
func NewQuerySensorUpdatePoliciesForbidden() *QuerySensorUpdatePoliciesForbidden {
	return &QuerySensorUpdatePoliciesForbidden{}
}

/*
QuerySensorUpdatePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QuerySensorUpdatePoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query sensor update policies forbidden response has a 2xx status code
func (o *QuerySensorUpdatePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policies forbidden response has a 3xx status code
func (o *QuerySensorUpdatePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policies forbidden response has a 4xx status code
func (o *QuerySensorUpdatePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policies forbidden response has a 5xx status code
func (o *QuerySensorUpdatePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policies forbidden response a status code equal to that given
func (o *QuerySensorUpdatePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query sensor update policies forbidden response
func (o *QuerySensorUpdatePoliciesForbidden) Code() int {
	return 403
}

func (o *QuerySensorUpdatePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorUpdatePoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorUpdatePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePoliciesTooManyRequests creates a QuerySensorUpdatePoliciesTooManyRequests with default headers values
func NewQuerySensorUpdatePoliciesTooManyRequests() *QuerySensorUpdatePoliciesTooManyRequests {
	return &QuerySensorUpdatePoliciesTooManyRequests{}
}

/*
QuerySensorUpdatePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QuerySensorUpdatePoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query sensor update policies too many requests response has a 2xx status code
func (o *QuerySensorUpdatePoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policies too many requests response has a 3xx status code
func (o *QuerySensorUpdatePoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policies too many requests response has a 4xx status code
func (o *QuerySensorUpdatePoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policies too many requests response has a 5xx status code
func (o *QuerySensorUpdatePoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policies too many requests response a status code equal to that given
func (o *QuerySensorUpdatePoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query sensor update policies too many requests response
func (o *QuerySensorUpdatePoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QuerySensorUpdatePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorUpdatePoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorUpdatePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePoliciesInternalServerError creates a QuerySensorUpdatePoliciesInternalServerError with default headers values
func NewQuerySensorUpdatePoliciesInternalServerError() *QuerySensorUpdatePoliciesInternalServerError {
	return &QuerySensorUpdatePoliciesInternalServerError{}
}

/*
QuerySensorUpdatePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QuerySensorUpdatePoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query sensor update policies internal server error response has a 2xx status code
func (o *QuerySensorUpdatePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policies internal server error response has a 3xx status code
func (o *QuerySensorUpdatePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policies internal server error response has a 4xx status code
func (o *QuerySensorUpdatePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensor update policies internal server error response has a 5xx status code
func (o *QuerySensorUpdatePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query sensor update policies internal server error response a status code equal to that given
func (o *QuerySensorUpdatePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query sensor update policies internal server error response
func (o *QuerySensorUpdatePoliciesInternalServerError) Code() int {
	return 500
}

func (o *QuerySensorUpdatePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorUpdatePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorUpdatePoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePoliciesDefault creates a QuerySensorUpdatePoliciesDefault with default headers values
func NewQuerySensorUpdatePoliciesDefault(code int) *QuerySensorUpdatePoliciesDefault {
	return &QuerySensorUpdatePoliciesDefault{
		_statusCode: code,
	}
}

/*
QuerySensorUpdatePoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QuerySensorUpdatePoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query sensor update policies default response has a 2xx status code
func (o *QuerySensorUpdatePoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query sensor update policies default response has a 3xx status code
func (o *QuerySensorUpdatePoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query sensor update policies default response has a 4xx status code
func (o *QuerySensorUpdatePoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query sensor update policies default response has a 5xx status code
func (o *QuerySensorUpdatePoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query sensor update policies default response a status code equal to that given
func (o *QuerySensorUpdatePoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query sensor update policies default response
func (o *QuerySensorUpdatePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *QuerySensorUpdatePoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QuerySensorUpdatePoliciesDefault) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update/v1][%d] querySensorUpdatePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QuerySensorUpdatePoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
