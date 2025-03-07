// Code generated by go-swagger; DO NOT EDIT.

package ioc

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

// GetIndicatorsReportReader is a Reader for the GetIndicatorsReport structure.
type GetIndicatorsReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndicatorsReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIndicatorsReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetIndicatorsReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIndicatorsReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIndicatorsReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIndicatorsReportOK creates a GetIndicatorsReportOK with default headers values
func NewGetIndicatorsReportOK() *GetIndicatorsReportOK {
	return &GetIndicatorsReportOK{}
}

/*
GetIndicatorsReportOK describes a response with status code 200, with default header values.

OK
*/
type GetIndicatorsReportOK struct {

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

// IsSuccess returns true when this get indicators report o k response has a 2xx status code
func (o *GetIndicatorsReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get indicators report o k response has a 3xx status code
func (o *GetIndicatorsReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get indicators report o k response has a 4xx status code
func (o *GetIndicatorsReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get indicators report o k response has a 5xx status code
func (o *GetIndicatorsReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get indicators report o k response a status code equal to that given
func (o *GetIndicatorsReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get indicators report o k response
func (o *GetIndicatorsReportOK) Code() int {
	return 200
}

func (o *GetIndicatorsReportOK) Error() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportOK  %+v", 200, o.Payload)
}

func (o *GetIndicatorsReportOK) String() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportOK  %+v", 200, o.Payload)
}

func (o *GetIndicatorsReportOK) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetIndicatorsReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndicatorsReportForbidden creates a GetIndicatorsReportForbidden with default headers values
func NewGetIndicatorsReportForbidden() *GetIndicatorsReportForbidden {
	return &GetIndicatorsReportForbidden{}
}

/*
GetIndicatorsReportForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIndicatorsReportForbidden struct {

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

// IsSuccess returns true when this get indicators report forbidden response has a 2xx status code
func (o *GetIndicatorsReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get indicators report forbidden response has a 3xx status code
func (o *GetIndicatorsReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get indicators report forbidden response has a 4xx status code
func (o *GetIndicatorsReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get indicators report forbidden response has a 5xx status code
func (o *GetIndicatorsReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get indicators report forbidden response a status code equal to that given
func (o *GetIndicatorsReportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get indicators report forbidden response
func (o *GetIndicatorsReportForbidden) Code() int {
	return 403
}

func (o *GetIndicatorsReportForbidden) Error() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportForbidden  %+v", 403, o.Payload)
}

func (o *GetIndicatorsReportForbidden) String() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportForbidden  %+v", 403, o.Payload)
}

func (o *GetIndicatorsReportForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIndicatorsReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndicatorsReportTooManyRequests creates a GetIndicatorsReportTooManyRequests with default headers values
func NewGetIndicatorsReportTooManyRequests() *GetIndicatorsReportTooManyRequests {
	return &GetIndicatorsReportTooManyRequests{}
}

/*
GetIndicatorsReportTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIndicatorsReportTooManyRequests struct {

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

// IsSuccess returns true when this get indicators report too many requests response has a 2xx status code
func (o *GetIndicatorsReportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get indicators report too many requests response has a 3xx status code
func (o *GetIndicatorsReportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get indicators report too many requests response has a 4xx status code
func (o *GetIndicatorsReportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get indicators report too many requests response has a 5xx status code
func (o *GetIndicatorsReportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get indicators report too many requests response a status code equal to that given
func (o *GetIndicatorsReportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get indicators report too many requests response
func (o *GetIndicatorsReportTooManyRequests) Code() int {
	return 429
}

func (o *GetIndicatorsReportTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIndicatorsReportTooManyRequests) String() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] getIndicatorsReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIndicatorsReportTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIndicatorsReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndicatorsReportDefault creates a GetIndicatorsReportDefault with default headers values
func NewGetIndicatorsReportDefault(code int) *GetIndicatorsReportDefault {
	return &GetIndicatorsReportDefault{
		_statusCode: code,
	}
}

/*
GetIndicatorsReportDefault describes a response with status code -1, with default header values.

OK
*/
type GetIndicatorsReportDefault struct {
	_statusCode int

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this get indicators report default response has a 2xx status code
func (o *GetIndicatorsReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get indicators report default response has a 3xx status code
func (o *GetIndicatorsReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get indicators report default response has a 4xx status code
func (o *GetIndicatorsReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get indicators report default response has a 5xx status code
func (o *GetIndicatorsReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get indicators report default response a status code equal to that given
func (o *GetIndicatorsReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get indicators report default response
func (o *GetIndicatorsReportDefault) Code() int {
	return o._statusCode
}

func (o *GetIndicatorsReportDefault) Error() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] GetIndicatorsReport default  %+v", o._statusCode, o.Payload)
}

func (o *GetIndicatorsReportDefault) String() string {
	return fmt.Sprintf("[POST /iocs/entities/indicators-reports/v1][%d] GetIndicatorsReport default  %+v", o._statusCode, o.Payload)
}

func (o *GetIndicatorsReportDefault) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetIndicatorsReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
