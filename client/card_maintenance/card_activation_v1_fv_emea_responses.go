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

// CardActivationV1FvEmeaReader is a Reader for the CardActivationV1FvEmea structure.
type CardActivationV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CardActivationV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCardActivationV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCardActivationV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCardActivationV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCardActivationV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCardActivationV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCardActivationV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewCardActivationV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewCardActivationV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewCardActivationV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewCardActivationV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCardActivationV1FvEmeaOK creates a CardActivationV1FvEmeaOK with default headers values
func NewCardActivationV1FvEmeaOK() *CardActivationV1FvEmeaOK {
	return &CardActivationV1FvEmeaOK{}
}

/*CardActivationV1FvEmeaOK handles this case with default header values.

successful operation
*/
type CardActivationV1FvEmeaOK struct {
	Payload *models.CardActivationResponse
}

func (o *CardActivationV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *CardActivationV1FvEmeaOK) GetPayload() *models.CardActivationResponse {
	return o.Payload
}

func (o *CardActivationV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardActivationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaBadRequest creates a CardActivationV1FvEmeaBadRequest with default headers values
func NewCardActivationV1FvEmeaBadRequest() *CardActivationV1FvEmeaBadRequest {
	return &CardActivationV1FvEmeaBadRequest{}
}

/*CardActivationV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type CardActivationV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *CardActivationV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaUnauthorized creates a CardActivationV1FvEmeaUnauthorized with default headers values
func NewCardActivationV1FvEmeaUnauthorized() *CardActivationV1FvEmeaUnauthorized {
	return &CardActivationV1FvEmeaUnauthorized{}
}

/*CardActivationV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type CardActivationV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *CardActivationV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaForbidden creates a CardActivationV1FvEmeaForbidden with default headers values
func NewCardActivationV1FvEmeaForbidden() *CardActivationV1FvEmeaForbidden {
	return &CardActivationV1FvEmeaForbidden{}
}

/*CardActivationV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type CardActivationV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *CardActivationV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaNotFound creates a CardActivationV1FvEmeaNotFound with default headers values
func NewCardActivationV1FvEmeaNotFound() *CardActivationV1FvEmeaNotFound {
	return &CardActivationV1FvEmeaNotFound{}
}

/*CardActivationV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type CardActivationV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *CardActivationV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaTooManyRequests creates a CardActivationV1FvEmeaTooManyRequests with default headers values
func NewCardActivationV1FvEmeaTooManyRequests() *CardActivationV1FvEmeaTooManyRequests {
	return &CardActivationV1FvEmeaTooManyRequests{}
}

/*CardActivationV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type CardActivationV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *CardActivationV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardActivationV1FvEmeaStatus452 creates a CardActivationV1FvEmeaStatus452 with default headers values
func NewCardActivationV1FvEmeaStatus452() *CardActivationV1FvEmeaStatus452 {
	return &CardActivationV1FvEmeaStatus452{}
}

/*CardActivationV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type CardActivationV1FvEmeaStatus452 struct {
}

func (o *CardActivationV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaStatus452 ", 452)
}

func (o *CardActivationV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardActivationV1FvEmeaStatus453 creates a CardActivationV1FvEmeaStatus453 with default headers values
func NewCardActivationV1FvEmeaStatus453() *CardActivationV1FvEmeaStatus453 {
	return &CardActivationV1FvEmeaStatus453{}
}

/*CardActivationV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type CardActivationV1FvEmeaStatus453 struct {
}

func (o *CardActivationV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaStatus453 ", 453)
}

func (o *CardActivationV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardActivationV1FvEmeaStatus455 creates a CardActivationV1FvEmeaStatus455 with default headers values
func NewCardActivationV1FvEmeaStatus455() *CardActivationV1FvEmeaStatus455 {
	return &CardActivationV1FvEmeaStatus455{}
}

/*CardActivationV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type CardActivationV1FvEmeaStatus455 struct {
}

func (o *CardActivationV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaStatus455 ", 455)
}

func (o *CardActivationV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardActivationV1FvEmeaStatus465 creates a CardActivationV1FvEmeaStatus465 with default headers values
func NewCardActivationV1FvEmeaStatus465() *CardActivationV1FvEmeaStatus465 {
	return &CardActivationV1FvEmeaStatus465{}
}

/*CardActivationV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SEA01E : ORGANIZATION IS NOT NUMERIC OR MUST BE VALUE BETWEEN 000 AND 998<BR/>VPL5SEA01S : REQUESTED CARDorACCOUNT NUMBER IS NOT FOUND<BR/>VPL5SEA02E : ACCOUNT NUMBER IS NOT NUMERIC OR EQUALS SPACES<BR/>VPL5SEA02S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SEA03E : FUNCTION TYPE IS INVALID VALID VALUES A OR S<BR/>VPL5SEA03S : REQUEST FOR LOCAL ORG BUT PASSED FOREIGN ORG<BR/>VPL5SEA04S : FUNCTION TYPE MUST BE 'A' WHEN ACCOUNT NUMBER IS PROVIDED<BR/>VPL5SEA05S : CARD NUMBER NOT FOUND<BR/>VPL5SEA06S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SEA07S : CARD CROSS REFERENCE NOT FOUND<BR/>VPL5SEA08S : CARD NUMBER NOT FOUND<BR/>VPL5SEA09S : LOGO RECORD NOT FOUND<BR/>VPL5SEA10S : FUNCTION TYPE MUST BE 'A' WHEN ACCOUNT NUMBER IS PROVIDED<BR/>VPL5SEA11S : FOREIGN CARD NUMBER NOT FOUND<BR/>VPL5SEA12S : FOREIGN ACCOUNT NUMBER NOT FOUND<BR/>VPL5SEA13S : FOREIGN CARD CROSS REFERENCE NOT FOUND<BR/>VPL5SEA14S : FOREIGN CARD NUMBER NOT FOUND<BR/>VPL5SEAZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SEAZ2S : SERVICE INPUT TO CARD ACTIVATED SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SEAZ4S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SEAZ5S : INVALID ORG<BR/>VPL5SEAZ6S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SEAZ7S : ORGANIZATION FILE UNAVAILABLE<BR/>VPL5SEAZ8S : CARD FILE UNAVAILABLE<BR/>VPL5SEAZ9S : ACCOUNT FILE UNAVAILABLE<BR/>VPL5SEAZAS : ACCOUNTorCARD CROSS REFERENCE FILE UNAVAILABL
*/
type CardActivationV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *CardActivationV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardActivation][%d] cardActivationV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *CardActivationV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardActivationV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}