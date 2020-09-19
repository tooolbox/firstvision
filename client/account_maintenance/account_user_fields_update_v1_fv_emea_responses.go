// Code generated by go-swagger; DO NOT EDIT.

package account_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// AccountUserFieldsUpdateV1FvEmeaReader is a Reader for the AccountUserFieldsUpdateV1FvEmea structure.
type AccountUserFieldsUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUserFieldsUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountUserFieldsUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountUserFieldsUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountUserFieldsUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountUserFieldsUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountUserFieldsUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccountUserFieldsUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAccountUserFieldsUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAccountUserFieldsUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAccountUserFieldsUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewAccountUserFieldsUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountUserFieldsUpdateV1FvEmeaOK creates a AccountUserFieldsUpdateV1FvEmeaOK with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaOK() *AccountUserFieldsUpdateV1FvEmeaOK {
	return &AccountUserFieldsUpdateV1FvEmeaOK{}
}

/*AccountUserFieldsUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type AccountUserFieldsUpdateV1FvEmeaOK struct {
	Payload *models.AccountUserFieldsUpdateResponse
}

func (o *AccountUserFieldsUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaOK) GetPayload() *models.AccountUserFieldsUpdateResponse {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountUserFieldsUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaBadRequest creates a AccountUserFieldsUpdateV1FvEmeaBadRequest with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaBadRequest() *AccountUserFieldsUpdateV1FvEmeaBadRequest {
	return &AccountUserFieldsUpdateV1FvEmeaBadRequest{}
}

/*AccountUserFieldsUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AccountUserFieldsUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaUnauthorized creates a AccountUserFieldsUpdateV1FvEmeaUnauthorized with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaUnauthorized() *AccountUserFieldsUpdateV1FvEmeaUnauthorized {
	return &AccountUserFieldsUpdateV1FvEmeaUnauthorized{}
}

/*AccountUserFieldsUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AccountUserFieldsUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaForbidden creates a AccountUserFieldsUpdateV1FvEmeaForbidden with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaForbidden() *AccountUserFieldsUpdateV1FvEmeaForbidden {
	return &AccountUserFieldsUpdateV1FvEmeaForbidden{}
}

/*AccountUserFieldsUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AccountUserFieldsUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaNotFound creates a AccountUserFieldsUpdateV1FvEmeaNotFound with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaNotFound() *AccountUserFieldsUpdateV1FvEmeaNotFound {
	return &AccountUserFieldsUpdateV1FvEmeaNotFound{}
}

/*AccountUserFieldsUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AccountUserFieldsUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaTooManyRequests creates a AccountUserFieldsUpdateV1FvEmeaTooManyRequests with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaTooManyRequests() *AccountUserFieldsUpdateV1FvEmeaTooManyRequests {
	return &AccountUserFieldsUpdateV1FvEmeaTooManyRequests{}
}

/*AccountUserFieldsUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AccountUserFieldsUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaStatus452 creates a AccountUserFieldsUpdateV1FvEmeaStatus452 with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaStatus452() *AccountUserFieldsUpdateV1FvEmeaStatus452 {
	return &AccountUserFieldsUpdateV1FvEmeaStatus452{}
}

/*AccountUserFieldsUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AccountUserFieldsUpdateV1FvEmeaStatus452 struct {
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaStatus452 ", 452)
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaStatus453 creates a AccountUserFieldsUpdateV1FvEmeaStatus453 with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaStatus453() *AccountUserFieldsUpdateV1FvEmeaStatus453 {
	return &AccountUserFieldsUpdateV1FvEmeaStatus453{}
}

/*AccountUserFieldsUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AccountUserFieldsUpdateV1FvEmeaStatus453 struct {
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaStatus453 ", 453)
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaStatus455 creates a AccountUserFieldsUpdateV1FvEmeaStatus455 with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaStatus455() *AccountUserFieldsUpdateV1FvEmeaStatus455 {
	return &AccountUserFieldsUpdateV1FvEmeaStatus455{}
}

/*AccountUserFieldsUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AccountUserFieldsUpdateV1FvEmeaStatus455 struct {
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaStatus455 ", 455)
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV1FvEmeaStatus465 creates a AccountUserFieldsUpdateV1FvEmeaStatus465 with default headers values
func NewAccountUserFieldsUpdateV1FvEmeaStatus465() *AccountUserFieldsUpdateV1FvEmeaStatus465 {
	return &AccountUserFieldsUpdateV1FvEmeaStatus465{}
}

/*AccountUserFieldsUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SUD02S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SUD03E : ORGANIZATION IS NOT NUMERIC OR MUST BE VALUE BETWEEN 000 AND 998<BR/>VPL5SUD04E : REQUESTED CARDorACCT NUMBER IS NOT NUMERIC OR EQUAL SPACES<BR/>VPL5SUD05E : INVALID FOREIGN USE IND SPECIFIED - VALID VALUES SPACE 0 1 OR 2<BR/>VPL5SUD06S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SUD07S : ORGANIZATION NOT FOUND ON FILE<BR/>VPL5SUD08S : FOREIGN ORGANIZATION NOT FOUND ON FILE<BR/>VPL5SUD09S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SUD10S : ACCOUNT NUMBER NOT FOUND ON FILE<BR/>VPL5SUD11S : LOGO NOT FOUND ON FILE<BR/>VPL5SUD13I : FOREIGN USE ORG UNAVAILABLE PROCESSED LOCAL ORG ACCOUNT DATA<BR/>VPL5SUD18S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SUD19S : SERVICE INPUT TO USER DATES AMTS AND CODES IS INCORRECT LENGTH<BR/>VPL5SUD20S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SUD21S : REQUESTED ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SUD22S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPL5SUD23S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SUD50E : USER DATE 1 IS NOT NUMERIC<BR/>VPL5SUD51E : USER DATE 1 IS INVALID<BR/>VPL5SUD52E : USER DATE 2 IS NOT NUMERIC<BR/>VPL5SUD53E : USER DATE 2 IS INVALID<BR/>VPL5SUD54E : USER DATE 3 IS NOT NUMERIC<BR/>VPL5SUD55E : USER DATE 3 IS INVALID<BR/>VPL5SUD56E : USER DATE 4 IS NOT NUMERIC<BR/>VPL5SUD57E : USER DATE 4 IS INVALID<BR/>VPL5SUD58E : USER DATE 5 IS NOT NUMERIC<BR/>VPL5SUD59E : USER DATE 5 IS INVALID<BR/>VPL5SUD60E : USER DATE 6 IS NOT NUMERIC<BR/>VPL5SUD61E : USER DATE 6 IS INVALID<BR/>VPL5SUD62E : USER DATE 7 IS NOT NUMERIC<BR/>VPL5SUD63E : USER DATE 7 IS INVALID<BR/>VPL5SUD64E : USER DATE 8 IS NOT NUMERIC<BR/>VPL5SUD65E : USER DATE 8 IS INVALID<BR/>VPL5SUD66E : USER DATE 9 IS NOT NUMERIC<BR/>VPL5SUD67E : USER DATE 9 IS INVALID<BR/>VPL5SUD68E : USER DATE 10 IS NOT NUMERIC<BR/>VPL5SUD69E : USER DATE 10 IS INVALID<BR/>VPL5SUD70E : USER DATE 11 IS NOT NUMERIC<BR/>VPL5SUD71E : USER DATE 11 IS INVALID<BR/>VPL5SUD72E : USER DATE 12 IS NOT NUMERIC<BR/>VPL5SUD73E : USER DATE 12 IS INVALID<BR/>VPL5SUD74E : USER DATE 13 IS NOT NUMERIC<BR/>VPL5SUD75E : USER DATE 13 IS INVALID<BR/>VPL5SUD76E : USER DATE 14 IS NOT NUMERIC<BR/>VPL5SUD77E : USER DATE 14 IS INVALID<BR/>VPL5SUD81E : USER AMOUNT 1 IS NOT NUMERIC<BR/>VPL5SUD82E : USER AMOUNT 2 IS NOT NUMERIC<BR/>VPL5SUD83E : USER AMOUNT 3 IS NOT NUMERIC<BR/>VPL5SUD84E : USER AMOUNT 4 IS NOT NUMERIC<BR/>VPL5SUD85E : USER AMOUNT 5 IS NOT NUMERIC<BR/>VPL5SUD86E : USER AMOUNT 6 IS NOT NUMERIC<BR/>VPL5SUD87E : USER AMOUNT 7 IS NOT NUMERIC<BR/>VPL5SUD88E : USER AMOUNT 8 IS NOT NUMERIC<BR/>VPL5SUD89E : USER AMOUNT 9 IS NOT NUMERIC<BR/>VPL5SUD90E : USER AMOUNT 10 IS NOT NUMERIC<BR/>VPL5SUD91E : USER AMOUNT 11 IS NOT NUMERIC<BR/>VPL5SUD92E : USER AMOUNT 12 IS NOT NUMERIC<BR/>VPL5SUD93E : USER AMOUNT 13 IS NOT NUMERIC<BR/>VPL5SUD94E : USER AMOUNT 14 IS NOT NUMERIC<BR/>VPL5SUD95E : FI SUB STATUS UDPATE FLAG IS NOT NUMERIC OR NOT EQUAL TO 0 AND
*/
type AccountUserFieldsUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}