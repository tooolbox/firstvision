// Code generated by go-swagger; DO NOT EDIT.

package card_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// PinBlockChangeV1FvEmeaReader is a Reader for the PinBlockChangeV1FvEmea structure.
type PinBlockChangeV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PinBlockChangeV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPinBlockChangeV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPinBlockChangeV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPinBlockChangeV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPinBlockChangeV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPinBlockChangeV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPinBlockChangeV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewPinBlockChangeV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewPinBlockChangeV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewPinBlockChangeV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewPinBlockChangeV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPinBlockChangeV1FvEmeaOK creates a PinBlockChangeV1FvEmeaOK with default headers values
func NewPinBlockChangeV1FvEmeaOK() *PinBlockChangeV1FvEmeaOK {
	return &PinBlockChangeV1FvEmeaOK{}
}

/*PinBlockChangeV1FvEmeaOK handles this case with default header values.

successful operation
*/
type PinBlockChangeV1FvEmeaOK struct {
	Payload *models.PinBlockChangeResponse
}

func (o *PinBlockChangeV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaOK) GetPayload() *models.PinBlockChangeResponse {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PinBlockChangeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaBadRequest creates a PinBlockChangeV1FvEmeaBadRequest with default headers values
func NewPinBlockChangeV1FvEmeaBadRequest() *PinBlockChangeV1FvEmeaBadRequest {
	return &PinBlockChangeV1FvEmeaBadRequest{}
}

/*PinBlockChangeV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type PinBlockChangeV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaUnauthorized creates a PinBlockChangeV1FvEmeaUnauthorized with default headers values
func NewPinBlockChangeV1FvEmeaUnauthorized() *PinBlockChangeV1FvEmeaUnauthorized {
	return &PinBlockChangeV1FvEmeaUnauthorized{}
}

/*PinBlockChangeV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type PinBlockChangeV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaForbidden creates a PinBlockChangeV1FvEmeaForbidden with default headers values
func NewPinBlockChangeV1FvEmeaForbidden() *PinBlockChangeV1FvEmeaForbidden {
	return &PinBlockChangeV1FvEmeaForbidden{}
}

/*PinBlockChangeV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type PinBlockChangeV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaNotFound creates a PinBlockChangeV1FvEmeaNotFound with default headers values
func NewPinBlockChangeV1FvEmeaNotFound() *PinBlockChangeV1FvEmeaNotFound {
	return &PinBlockChangeV1FvEmeaNotFound{}
}

/*PinBlockChangeV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type PinBlockChangeV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaTooManyRequests creates a PinBlockChangeV1FvEmeaTooManyRequests with default headers values
func NewPinBlockChangeV1FvEmeaTooManyRequests() *PinBlockChangeV1FvEmeaTooManyRequests {
	return &PinBlockChangeV1FvEmeaTooManyRequests{}
}

/*PinBlockChangeV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type PinBlockChangeV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinBlockChangeV1FvEmeaStatus452 creates a PinBlockChangeV1FvEmeaStatus452 with default headers values
func NewPinBlockChangeV1FvEmeaStatus452() *PinBlockChangeV1FvEmeaStatus452 {
	return &PinBlockChangeV1FvEmeaStatus452{}
}

/*PinBlockChangeV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type PinBlockChangeV1FvEmeaStatus452 struct {
}

func (o *PinBlockChangeV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaStatus452 ", 452)
}

func (o *PinBlockChangeV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinBlockChangeV1FvEmeaStatus453 creates a PinBlockChangeV1FvEmeaStatus453 with default headers values
func NewPinBlockChangeV1FvEmeaStatus453() *PinBlockChangeV1FvEmeaStatus453 {
	return &PinBlockChangeV1FvEmeaStatus453{}
}

/*PinBlockChangeV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type PinBlockChangeV1FvEmeaStatus453 struct {
}

func (o *PinBlockChangeV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaStatus453 ", 453)
}

func (o *PinBlockChangeV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinBlockChangeV1FvEmeaStatus455 creates a PinBlockChangeV1FvEmeaStatus455 with default headers values
func NewPinBlockChangeV1FvEmeaStatus455() *PinBlockChangeV1FvEmeaStatus455 {
	return &PinBlockChangeV1FvEmeaStatus455{}
}

/*PinBlockChangeV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type PinBlockChangeV1FvEmeaStatus455 struct {
}

func (o *PinBlockChangeV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaStatus455 ", 455)
}

func (o *PinBlockChangeV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinBlockChangeV1FvEmeaStatus465 creates a PinBlockChangeV1FvEmeaStatus465 with default headers values
func NewPinBlockChangeV1FvEmeaStatus465() *PinBlockChangeV1FvEmeaStatus465 {
	return &PinBlockChangeV1FvEmeaStatus465{}
}

/*PinBlockChangeV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPE5SCP02S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPE5SCP09S : INVALID ACCOUNT STATUS PIN CHANGE NOT ALLOWED<BR/>VPE5SCP10S : INVALID ACCOUNT TYPE PIN CHANGE NOT ALLOWED<BR/>VPE5SCP11S : ACCOUNT BLOCKED PIN CHANGE NOT ALLOWED<BR/>VPE5SCP13S : INVALID CARD STATUS PIN CHANGE NOT ALLOWED<BR/>VPE5SCP14S : CARD BLOCKED PIN CHANGE NOT ALLOWED<BR/>VPE5SCP15S : CARD IS NOT ACTIVATED PIN CHANGE NOT ALLOWED<BR/>VPE5SCP16S : PIN SUPPRESSION IS ON PIN CHANGE NOT ALLOWED<BR/>VPE5SCP17S : CARD IS EXPIRED PIN CHANGE NOT ALLOWED<BR/>VPE5SCP25S : INVALID PRODUCT TYPE PIN CHANGE NOT ALLOWED<BR/>VPE5SCP26S : INVALID CURRENT PIN BLOCK<BR/>VPE5SCP27S : PIN BLOCK ERRROR<BR/>VPE5SCP28S : PIN BLOCK KEY ERRRO
*/
type PinBlockChangeV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *PinBlockChangeV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pinBlockChange][%d] pinBlockChangeV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *PinBlockChangeV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PinBlockChangeV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}