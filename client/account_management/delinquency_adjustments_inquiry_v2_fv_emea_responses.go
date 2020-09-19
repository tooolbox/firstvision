// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// DelinquencyAdjustmentsInquiryV2FvEmeaReader is a Reader for the DelinquencyAdjustmentsInquiryV2FvEmea structure.
type DelinquencyAdjustmentsInquiryV2FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DelinquencyAdjustmentsInquiryV2FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaOK creates a DelinquencyAdjustmentsInquiryV2FvEmeaOK with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaOK() *DelinquencyAdjustmentsInquiryV2FvEmeaOK {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaOK{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaOK handles this case with default header values.

successful operation
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaOK struct {
	Payload *models.DelinquencyAdjustmentsInquiryResponse2
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaOK  %+v", 200, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaOK) GetPayload() *models.DelinquencyAdjustmentsInquiryResponse2 {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DelinquencyAdjustmentsInquiryResponse2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaBadRequest creates a DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaBadRequest() *DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized creates a DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized() *DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaForbidden creates a DelinquencyAdjustmentsInquiryV2FvEmeaForbidden with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaForbidden() *DelinquencyAdjustmentsInquiryV2FvEmeaForbidden {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaForbidden{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaNotFound creates a DelinquencyAdjustmentsInquiryV2FvEmeaNotFound with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaNotFound() *DelinquencyAdjustmentsInquiryV2FvEmeaNotFound {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaNotFound{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests creates a DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests() *DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus452 creates a DelinquencyAdjustmentsInquiryV2FvEmeaStatus452 with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus452() *DelinquencyAdjustmentsInquiryV2FvEmeaStatus452 {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaStatus452{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaStatus452 struct {
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaStatus452 ", 452)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus453 creates a DelinquencyAdjustmentsInquiryV2FvEmeaStatus453 with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus453() *DelinquencyAdjustmentsInquiryV2FvEmeaStatus453 {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaStatus453{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaStatus453 struct {
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaStatus453 ", 453)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus455 creates a DelinquencyAdjustmentsInquiryV2FvEmeaStatus455 with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus455() *DelinquencyAdjustmentsInquiryV2FvEmeaStatus455 {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaStatus455{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaStatus455 struct {
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaStatus455 ", 455)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus465 creates a DelinquencyAdjustmentsInquiryV2FvEmeaStatus465 with default headers values
func NewDelinquencyAdjustmentsInquiryV2FvEmeaStatus465() *DelinquencyAdjustmentsInquiryV2FvEmeaStatus465 {
	return &DelinquencyAdjustmentsInquiryV2FvEmeaStatus465{}
}

/*DelinquencyAdjustmentsInquiryV2FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SDJZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SDJZ2S : SERVICE INPUT TO DELQ ADJ SERVICE IS AN INCORRECT<BR/>VPL5SDJZ4S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SDJZ5S : REQUESTED ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SDJZ6S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SDJZ7S : FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPL5SDJZ8S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUT<BR/>VPL5SDJ01E : ORGANIZATION IS NOT NUMERIC OR MUST BE VALUE BETWEEN 000 AND 998<BR/>VPL5SDJ02S : ACCOUNT NUMBER IS NOT NUMERIC OR EQUAL SPACES<BR/>VPL5SDJ03S : ORGANIZATION NOT FOUND ON FILE OR IN ADD PENDING<BR/>VPL5SDJ04S : ACCOUNT NOT FOUND ON FILE<BR/>VPL5SDJ05S : DELINQUENCY ADJ NOT ALLOWED FOR THIS ACCOUNT<BR/>VPL5SDJ06S : LOGO RECORD NOT FOUND OR IN ADD PENDING<BR/>VPL5SDJ07S : NO PLAN SEGMENTS ON FIL
*/
type DelinquencyAdjustmentsInquiryV2FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/delinquencyAdjustmentsInquiry][%d] delinquencyAdjustmentsInquiryV2FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DelinquencyAdjustmentsInquiryV2FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}