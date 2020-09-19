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

// CardslistbyAccountV1FvEmeaReader is a Reader for the CardslistbyAccountV1FvEmea structure.
type CardslistbyAccountV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CardslistbyAccountV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCardslistbyAccountV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCardslistbyAccountV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCardslistbyAccountV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCardslistbyAccountV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCardslistbyAccountV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCardslistbyAccountV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewCardslistbyAccountV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewCardslistbyAccountV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewCardslistbyAccountV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewCardslistbyAccountV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCardslistbyAccountV1FvEmeaOK creates a CardslistbyAccountV1FvEmeaOK with default headers values
func NewCardslistbyAccountV1FvEmeaOK() *CardslistbyAccountV1FvEmeaOK {
	return &CardslistbyAccountV1FvEmeaOK{}
}

/*CardslistbyAccountV1FvEmeaOK handles this case with default header values.

successful operation
*/
type CardslistbyAccountV1FvEmeaOK struct {
	Payload *models.CardslistbyAccountResponse
}

func (o *CardslistbyAccountV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaOK) GetPayload() *models.CardslistbyAccountResponse {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardslistbyAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaBadRequest creates a CardslistbyAccountV1FvEmeaBadRequest with default headers values
func NewCardslistbyAccountV1FvEmeaBadRequest() *CardslistbyAccountV1FvEmeaBadRequest {
	return &CardslistbyAccountV1FvEmeaBadRequest{}
}

/*CardslistbyAccountV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type CardslistbyAccountV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaUnauthorized creates a CardslistbyAccountV1FvEmeaUnauthorized with default headers values
func NewCardslistbyAccountV1FvEmeaUnauthorized() *CardslistbyAccountV1FvEmeaUnauthorized {
	return &CardslistbyAccountV1FvEmeaUnauthorized{}
}

/*CardslistbyAccountV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type CardslistbyAccountV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaForbidden creates a CardslistbyAccountV1FvEmeaForbidden with default headers values
func NewCardslistbyAccountV1FvEmeaForbidden() *CardslistbyAccountV1FvEmeaForbidden {
	return &CardslistbyAccountV1FvEmeaForbidden{}
}

/*CardslistbyAccountV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type CardslistbyAccountV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaNotFound creates a CardslistbyAccountV1FvEmeaNotFound with default headers values
func NewCardslistbyAccountV1FvEmeaNotFound() *CardslistbyAccountV1FvEmeaNotFound {
	return &CardslistbyAccountV1FvEmeaNotFound{}
}

/*CardslistbyAccountV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type CardslistbyAccountV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaTooManyRequests creates a CardslistbyAccountV1FvEmeaTooManyRequests with default headers values
func NewCardslistbyAccountV1FvEmeaTooManyRequests() *CardslistbyAccountV1FvEmeaTooManyRequests {
	return &CardslistbyAccountV1FvEmeaTooManyRequests{}
}

/*CardslistbyAccountV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type CardslistbyAccountV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardslistbyAccountV1FvEmeaStatus452 creates a CardslistbyAccountV1FvEmeaStatus452 with default headers values
func NewCardslistbyAccountV1FvEmeaStatus452() *CardslistbyAccountV1FvEmeaStatus452 {
	return &CardslistbyAccountV1FvEmeaStatus452{}
}

/*CardslistbyAccountV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type CardslistbyAccountV1FvEmeaStatus452 struct {
}

func (o *CardslistbyAccountV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaStatus452 ", 452)
}

func (o *CardslistbyAccountV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardslistbyAccountV1FvEmeaStatus453 creates a CardslistbyAccountV1FvEmeaStatus453 with default headers values
func NewCardslistbyAccountV1FvEmeaStatus453() *CardslistbyAccountV1FvEmeaStatus453 {
	return &CardslistbyAccountV1FvEmeaStatus453{}
}

/*CardslistbyAccountV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type CardslistbyAccountV1FvEmeaStatus453 struct {
}

func (o *CardslistbyAccountV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaStatus453 ", 453)
}

func (o *CardslistbyAccountV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardslistbyAccountV1FvEmeaStatus455 creates a CardslistbyAccountV1FvEmeaStatus455 with default headers values
func NewCardslistbyAccountV1FvEmeaStatus455() *CardslistbyAccountV1FvEmeaStatus455 {
	return &CardslistbyAccountV1FvEmeaStatus455{}
}

/*CardslistbyAccountV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type CardslistbyAccountV1FvEmeaStatus455 struct {
}

func (o *CardslistbyAccountV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaStatus455 ", 455)
}

func (o *CardslistbyAccountV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardslistbyAccountV1FvEmeaStatus465 creates a CardslistbyAccountV1FvEmeaStatus465 with default headers values
func NewCardslistbyAccountV1FvEmeaStatus465() *CardslistbyAccountV1FvEmeaStatus465 {
	return &CardslistbyAccountV1FvEmeaStatus465{}
}

/*CardslistbyAccountV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SAEZ1S : SERVICE INPUT TO ACCCT TO CARD NAV SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SAEZ3S : FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPL5SAEZ4S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPL5SAEZ4S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPL5SAEZ5S : ORGANIZATION FILE NOT AVAILABLE<BR/>VPL5SAEZ6S : ACCOUNT FILE NOT AVAILABLE<BR/>VPL5SAEZ7S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SAEZ8S : REQUESTED ORG NOT FOUND<BR/>VPL5SAE01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-998 ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE          001-998<BR/>VPL5SAE01S : REQUESTED CARDorACCOUNT NUMBER IS NOT FOUND<BR/>VPL5SAE02E : ACCOUNT NUMBER IS NOT NUMERIC OR EQUALS SPACES<BR/>VPL5SAE02S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SAE03E : SERVICE FUNCTION INVALID - VALID VALUES ARE P N OR SPACE<BR/>VPL5SAE03S : CARD NUMBER NOT FOUND<BR/>VPL5SAE04E : INPUT START CARD ORG NOT NUMERIC<BR/>VPL5SAE04S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SAE05E : INPUT START CARD NUMBER NOT NUMERIC<BR/>VPL5SAE06E : INPUT START CARD SEQ IS NOT VALID<BR/>VPL5SAE07E : NO VALID CARDS FOUND FOR THE ACCOUNT<BR/>VPL5S9900S : RESOURCE IS UNAVAILABLE - CONTACT YOUR ADMINISTRATO
*/
type CardslistbyAccountV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *CardslistbyAccountV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardslistbyAccount][%d] cardslistbyAccountV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *CardslistbyAccountV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardslistbyAccountV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}