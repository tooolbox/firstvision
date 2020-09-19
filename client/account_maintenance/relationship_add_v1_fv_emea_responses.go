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

// RelationshipAddV1FvEmeaReader is a Reader for the RelationshipAddV1FvEmea structure.
type RelationshipAddV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelationshipAddV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelationshipAddV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRelationshipAddV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRelationshipAddV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRelationshipAddV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRelationshipAddV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRelationshipAddV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewRelationshipAddV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewRelationshipAddV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewRelationshipAddV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewRelationshipAddV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelationshipAddV1FvEmeaOK creates a RelationshipAddV1FvEmeaOK with default headers values
func NewRelationshipAddV1FvEmeaOK() *RelationshipAddV1FvEmeaOK {
	return &RelationshipAddV1FvEmeaOK{}
}

/*RelationshipAddV1FvEmeaOK handles this case with default header values.

successful operation
*/
type RelationshipAddV1FvEmeaOK struct {
	Payload *models.RelationshipAddResponse
}

func (o *RelationshipAddV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *RelationshipAddV1FvEmeaOK) GetPayload() *models.RelationshipAddResponse {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAddResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaBadRequest creates a RelationshipAddV1FvEmeaBadRequest with default headers values
func NewRelationshipAddV1FvEmeaBadRequest() *RelationshipAddV1FvEmeaBadRequest {
	return &RelationshipAddV1FvEmeaBadRequest{}
}

/*RelationshipAddV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type RelationshipAddV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *RelationshipAddV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaUnauthorized creates a RelationshipAddV1FvEmeaUnauthorized with default headers values
func NewRelationshipAddV1FvEmeaUnauthorized() *RelationshipAddV1FvEmeaUnauthorized {
	return &RelationshipAddV1FvEmeaUnauthorized{}
}

/*RelationshipAddV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type RelationshipAddV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *RelationshipAddV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaForbidden creates a RelationshipAddV1FvEmeaForbidden with default headers values
func NewRelationshipAddV1FvEmeaForbidden() *RelationshipAddV1FvEmeaForbidden {
	return &RelationshipAddV1FvEmeaForbidden{}
}

/*RelationshipAddV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type RelationshipAddV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *RelationshipAddV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaNotFound creates a RelationshipAddV1FvEmeaNotFound with default headers values
func NewRelationshipAddV1FvEmeaNotFound() *RelationshipAddV1FvEmeaNotFound {
	return &RelationshipAddV1FvEmeaNotFound{}
}

/*RelationshipAddV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type RelationshipAddV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *RelationshipAddV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaTooManyRequests creates a RelationshipAddV1FvEmeaTooManyRequests with default headers values
func NewRelationshipAddV1FvEmeaTooManyRequests() *RelationshipAddV1FvEmeaTooManyRequests {
	return &RelationshipAddV1FvEmeaTooManyRequests{}
}

/*RelationshipAddV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type RelationshipAddV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *RelationshipAddV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRelationshipAddV1FvEmeaStatus452 creates a RelationshipAddV1FvEmeaStatus452 with default headers values
func NewRelationshipAddV1FvEmeaStatus452() *RelationshipAddV1FvEmeaStatus452 {
	return &RelationshipAddV1FvEmeaStatus452{}
}

/*RelationshipAddV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type RelationshipAddV1FvEmeaStatus452 struct {
}

func (o *RelationshipAddV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaStatus452 ", 452)
}

func (o *RelationshipAddV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRelationshipAddV1FvEmeaStatus453 creates a RelationshipAddV1FvEmeaStatus453 with default headers values
func NewRelationshipAddV1FvEmeaStatus453() *RelationshipAddV1FvEmeaStatus453 {
	return &RelationshipAddV1FvEmeaStatus453{}
}

/*RelationshipAddV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type RelationshipAddV1FvEmeaStatus453 struct {
}

func (o *RelationshipAddV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaStatus453 ", 453)
}

func (o *RelationshipAddV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRelationshipAddV1FvEmeaStatus455 creates a RelationshipAddV1FvEmeaStatus455 with default headers values
func NewRelationshipAddV1FvEmeaStatus455() *RelationshipAddV1FvEmeaStatus455 {
	return &RelationshipAddV1FvEmeaStatus455{}
}

/*RelationshipAddV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type RelationshipAddV1FvEmeaStatus455 struct {
}

func (o *RelationshipAddV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaStatus455 ", 455)
}

func (o *RelationshipAddV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRelationshipAddV1FvEmeaStatus465 creates a RelationshipAddV1FvEmeaStatus465 with default headers values
func NewRelationshipAddV1FvEmeaStatus465() *RelationshipAddV1FvEmeaStatus465 {
	return &RelationshipAddV1FvEmeaStatus465{}
}

/*RelationshipAddV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SAR01S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SAR02S : SERVICE INPUT TO AB SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SAR03S : CUSTOMER NUMBER IS INVALID<BR/>VPL5SAR04S : FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPL5SAR05S : CMS IS NO-PROCESSING STATE RE-TRY REQUEST IN A FEW MINUTES<BR/>VPL5SAR06S : ORGANIZATION NUMBER MUST BE PROVIDED<BR/>VPL5SAR07S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SAR08S : SYSTEM RECORD NOT FOUND<BR/>VPL5SAR09S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SAR10S : RELATIONSHIPS NOT ACTIVATED IN THIS ORGANIZATION<BR/>VPL5SAR11S : CUSTOMER NUMBER MUST BE PROVIDED<BR/>VPL5SAR12S : CREDIT LIMIT MUST BE PROVIDED<BR/>VPL5SAR13S : BILLING MODIFY MUST BE 1 WHEN BILLING LEVEL IS 0 OR 1<BR/>VPL5SAR14S : COMMERCIAL CARD FLAG MUST BE 0 NOT A COMMERCIAL RELATIONSHIP<BR/>VPL5SAR15S : COMMERCIAL FLAG MUST BE 3 WHEN HCS IS ACTIVE<BR/>VPL5SAR16S : COMMERCIAL FLAG CANNOT BE 3 IF HCS IS NOT ACTIVE<BR/>VPL5SAR17S : CORPORATION CUSTOMER NUMBER IS REQUIRED<BR/>VPL5SAR18S : CUSTOMER GROUP CODE IS REQUIRED<BR/>VPL5SAR19S : CUSTOMER NUMBER NOT FOUND<BR/>VPL5SAR20S : ACCOUNT CONTROL OVERRIDE IS INVALID<BR/>VPL5SAR21S : BILLING CURRENCY NOT FOUND<BR/>VPL5SAR22S : TRANSACTION CODE MUST POINT TO LOGIC MODULE 99<BR/>VPL5SAR23S : REVERSAL TRANSACTION CODE MUST POINT TO LOGIC MODULE 99<BR/>VPL5SAR24S : INVALID BILLING CYCLE<BR/>VPL5SAR25S : MEMO PMT and REV TRANS CAN'T BE 0 WHEN BILL LVL IS 0 OR 2<BR/>VPL5SAR26S : MEMO PMT TRANS and MEMO PMT REV TRANS CANNOT BE THE SAME<BR/>VPL5SAR27S : THERE MUST BE A VALID ACCT CTRL TABLE WHEN BILLING LVL IS 0 OR 2<BR/>VPL5SAR28S : ACCOUNT CONTROL TABLE MUST BE ZERO WHEN BILLING LEVEL IS 1<BR/>VPL5SAR29S : SUBORDINATE CR LIMIT CAN'T BE GREATER THAN RELATIONSHIP CR LIMIT<BR/>VPL5SAR30S : CORPORATE CUSTOMER NUMBER IS NOT FOUND<BR/>VPL5SAR31S : INVALID AUTHORIZATION CIRTERIA TABLE NUMBER<BR/>VPL5SAR32S : SERVICE COULD NOT OBTAIN STORAGE FOR INPUT AREA<BR/>VPL5SAR33S : SERVICE COULD NOT OBTAIN STORAGE FOR OUTPUT AREA<BR/>VPL5SAR34S : NUMBER GENERATION SERVICE FAILED<BR/>VPL5SAR35S : MISMATCHED GENERATED FOREIGN ORG<BR/>VPL5SAR36S : RELATIONSHIP NUMBER ALREADY EXISTS<BR/>VPL5SAR37S : RELATIONSHIP NUMBER MUST NOT BE EQUAL TO EXISTING CUSTOMER<BR/>VPL5SAR38S : RELATIONSHIP NUMBER MUST NOT BE EQUAL TO EXISTING ACCOUNT<BR/>VPL5SAR39S : DUAL RELATIONSHIP ORG IS NOT FOUND<BR/>VPL5SAR40S : DUAL RELATIONSHIP NUMBER ALREADY EXISTS<BR/>VPL5SAR41S : DUAL RELATIONSHIP NUMBER MUST NOT BE EQUAL TO EXISTING CUSTOMER<BR/>VPL5SAR42S : DUAL RELATIONSHIP NUMBER MUST NOT BE EQUAL TO EXISTING ACCOUNT<BR/>VPL5SAR01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 000-998<BR/>VPL5SAR02E : INVALID RELATIONSHIP NUMBER<BR/>VPL5SAR03E : CUSTOMER NUMBER MUST BE NUMERIC<BR/>VPL5SAR04E : CREDIT LIMIT MUST BE NUMERIC AND GREATER THAN ZERO<BR/>VPL5SAR05E : INVALID BLOCK CODE<BR/>VPL5SAR06E : VALID MAIL CODE VALUES ARE 0 AND 1 ONLY<BR/>VPL5SAR07E : VALID AMTorPCT FLAG VALUES ARE 0 AND 1 ONLY<BR/>VPL5SAR08E : INVALID PERCENTAGE AMOUNT<BR/>VPL5SAR09E : AVAILABLE RESERVE AMTorPCT MUST BE NUMERIC<BR/>VPL5SAR10E : VALID STATEMENT TYPES ARE 0 AND 1<BR/>VPL5SAR11E : CNSLDTD IS NOT ALLOWED| ACCT CTRL TABLE OVERRIDE CAN'T BE INPUT<BR/>VPL5SAR12E : ACCOUNT CONTROL OVERRIDE MUST BE NUMERIC<BR/>VPL5SAR13E : BILLING CURRENCY MUST BE NUMERIC<BR/>VPL5SAR14E : CNSLDTD IS NOT ALLOWED| MEMO PAYMENT TXN NUMB CAN'T BE INPUT<BR/>VPL5SAR15E : MEMO PAYMENT TRANSACTION CODE MUST BE NUMERIC<BR/>VPL5SAR16E : CNSLDTD IS NOT ALLOWED| MEMO PMT REVERSAL TXN NUMB CAN'T BE INPUT<BR/>VPL5SAR17E : MEMO PAYMENT REVERSAL TRANSACTION CODE MUST BE NUMERIC<BR/>VPL5SAR18E : VALID MEMO BALANCE FLAG ARE 0 AND 1<BR/>VPL5SAR19E : VALID BILLING CYCLE VALUES ARE 01 THRU 31<BR/>VPL5SAR20E : BILLING CYCLE DAY IS REQUIRED<BR/>VPL5SAR21E : CNSLDTD IS NOT ALLOWED| BILLING LVL CAN'T BE INPUT<BR/>VPL5SAR22E : VALID BILLING LEVEL VALUES ARE 0 1 AND 2<BR/>VPL5SAR23E : CNSLDTD IS NOT ALLOWED| BILLING LVL MODIFY CAN'T BE INPUT<BR/>VPL5SAR24E : VALID BILLING LVL MODIFY VALUES ARE 0 AND 1<BR/>VPL5SAR25E : VALID ANNUAL FEE VALUES ARE 0 1 AND 2<BR/>VPL5SAR26E : VALID ANNUAL FEE MODIFY VALUES ARE 0 AND 1<BR/>VPL5SAR27E : VALID LATE FEE VALUES ARE 0 1 AND 2<BR/>VPL5SAR28E : VALID LATE FEE MODIFY VALUES ARE 0 AND 1<BR/>VPL5SAR29E : VALID NSF FEE VALUES ARE 0 1 AND 2<BR/>VPL5SAR30E : VALID NSF FEE MODIFY VALUES ARE 0 AND 1<BR/>VPL5SAR31E : VALID CREDIT LIMIT DEFAULT MUST BE NUMERIC<BR/>VPL5SAR32E : VALID CREDIT LIMIT DEFAULT MODIFY VALUES ARE 0 AND 1<BR/>VPL5SAR33E : VALID VALUES FOR COMMERICAL CARD ARE 0 THRU 3<BR/>VPL5SAR34E : NOT A COMMERCIAL CARD RELATIONSHIP| CO CUST NUMB CAN'T BE INPUT<BR/>VPL5SAR35E : NOT A COMMERCIAL CARD RELATIONSHIP|CUST GROUP CODE CAN'T BE INPUT<BR/>VPL5SAR36E : NOT A COMMERCIAL CARD RELATIONSHIP| AUTH CRITERIA TAB CAN'T INPUT<BR/>VPL5SAR37E : NOT A COMMERCIAL CARD RELATIONSHIP| COST CENTER NO CAN'T BE INPUT<BR/>VPL5SAR38E : NOT A COMMERCIAL CARD RELATIONSHIP| FISCAL YR END CAN'T BE INPUT<BR/>VPL5SAR39E : NOT A COMMERCIAL CARD RELATIONSHIP|CURR CONTRACT EXPD CAN'T INPUT<BR/>VPL5SAR40E : NOT A COMMERCIAL CARD RELATIONSHIP| CO ID NUMB CAN'T BE INPUT<BR/>VPL5SAR41E : NOT A COMMERCIAL CARD RELATIONSHIP| SOURCE CODE CAN'T BE INPUT<BR/>VPL5SAR42E : NOT A COMMERCIAL CARD RELATIONSHIP| CONTACT NAME CAN'T BE INPUT<BR/>VPL5SAR43E : NOT A COMMERCIAL CARD RELATIONSHIP| CONTACT PHONE CAN'T BE INPUT<BR/>VPL5SAR44E : NOT A COMMERCIAL CARD RELATIONSHIP| OFFICERS NAME CAN'T BE INPUT<BR/>VPL5SAR45E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 1 IS NOT OPEN FOR INPUT<BR/>VPL5SAR46E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 2 IS NOT OPEN FOR INPUT<BR/>VPL5SAR47E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 3 IS NOT OPEN FOR INPUT<BR/>VPL5SAR48E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 4 IS NOT OPEN FOR INPUT<BR/>VPL5SAR49E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 5 IS NOT OPEN FOR INPUT<BR/>VPL5SAR50E : NOT A COMMERCIAL CARD RELATIONSHIP| USER 6 IS NOT OPEN FOR INPUT<BR/>VPL5SAR51E : CORPORATION CUSTOMER NUMBER MUST BE NUMERIC<BR/>VPL5SAR52E : THE VALID VALUES FOR AUTH CRITERIA TAB NUMB ARE THRU AAA TO 999<BR/>VPL5SAR53E : COST CENTER REPORTING NUMBER IS REQUIRED<BR/>VPL5SAR54E : THE VALID VALUE OF FISCAL YEAR END IS THROUGH 01 TO 12<BR/>VPL5SAR55E : FISCAL YEAR END IS REQUIRED<BR/>VPL5SAR56E : NOT A COMMERCIAL ORG AND CURRENT CONTRACT EXPIRE DATE CAN'T INPUT<BR/>VPL5SAR57E : CURRENT CONTRACT EXPIRE DATE MUST BE NUMERIC AND GREATER THAN 0<BR/>VPL5SAR58E : INVALID DATE IN CURRENT CONTRACT EXPIRE DATE<BR/>VPL5SAR59E : CURRENT CONTRACT EXPIRE DATE MUST BE GREATER THAN TODAY<BR/>VPL5SAR60E : CORPORATE ID NUMBER MUST BE NUMERIC<BR/>VPL5SAR61E : USER DATE 1 MUST BE NUMERIC<BR/>VPL5SAR62E : USER DATE 1 - INVALID DATE<BR/>VPL5SAR63E : USER DATE 2 MUST BE NUMERIC<BR/>VPL5SAR64E : USER DATE 2 - INVALID DATE<BR/>VPL5SAR65E : USER AMOUNT 1 MUST BE NUMERIC<BR/>VPL5SAR66E : USER AMOUNT 2 MUST BE NUMERIC<BR/>VPL5SARADE : VALID VISA FLAG VALUES ARE 0 - 3<BR/>VPL5SARAEE : VALID MC FLAG VALUES ARE 0 - 3<BR/>VPL5SARAFE : VALID VIM FLAG VALUES ARE 0 AND 1<BR/>VPL5SARAGE : VALID SDOL FLAG VALUES ARE 0 AND 1<BR/>VPL5SARAHE : VALID VGIS FLAG VALUES ARE 0 AND 1<BR/>VPL5SARAIE : VALID DCAL FLAG VALUES ARE 0 AND 1<BR/>VPL5SARAJE : VALID EXT RPT FREQ VALUES ARE 00 01-31 41-47 70 AND 99<BR/>VPL5SARAKS : EXT RPT FREQ MUST BE 00 WHEN ALL EXT RPT FLAGS ARE 0<BR/>VPL5SARALS : VIM FLAG VALID ONLY WHEN VISA FLAG IS GREATER THAN 0<BR/>VPL5SARAMS : SDOL FLAG VALID ONLY WHEN MC FLAG IS GREATER THAN 0<BR/>VPL5SARANS : VGIS VALID ONLY FOR VISA PURCHASING CARDS<BR/>VPL5SARAOS : DCAL FLAG VALID ONLY WHEN MC OR VISA FLAG ARE GREATER THAN 0<BR/>VPL5SARAPS : DCAL FLAG VALID ONLY WHEN EXT RPT FREQ IS 99<BR/>VPL5SARAQS : DCAL FLAG VALID ONLY WHEN COMMERCIAL FLAG IS 4<BR/>VPL5SARARE : VALID DEF PUR AUTH VALUES ARE 0 - 2<BR/>VPL5SARASS : DEF PUR AUTH MUST BE 0 WHEN REL AUTH IS NOT ALLOWED<BR/>VPL5SARATE : VALID DEF CASH AUTH VALUES ARE 0 - 2<BR/>VPL5SARAUS : DEF CASH AUTH MUST BE 0 WHEN REL AUTH IS NOT ALLOWE
*/
type RelationshipAddV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *RelationshipAddV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/relationshipAdd][%d] relationshipAddV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *RelationshipAddV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *RelationshipAddV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}