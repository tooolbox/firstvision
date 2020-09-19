// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// PendingPaymentListV1FvEmeaReader is a Reader for the PendingPaymentListV1FvEmea structure.
type PendingPaymentListV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PendingPaymentListV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPendingPaymentListV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPendingPaymentListV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPendingPaymentListV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewPendingPaymentListV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewPendingPaymentListV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewPendingPaymentListV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 467:
		result := NewPendingPaymentListV1FvEmeaStatus467()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPendingPaymentListV1FvEmeaOK creates a PendingPaymentListV1FvEmeaOK with default headers values
func NewPendingPaymentListV1FvEmeaOK() *PendingPaymentListV1FvEmeaOK {
	return &PendingPaymentListV1FvEmeaOK{}
}

/*PendingPaymentListV1FvEmeaOK handles this case with default header values.

successful operation
*/
type PendingPaymentListV1FvEmeaOK struct {
	Payload *models.PendingPaymentListResponse
}

func (o *PendingPaymentListV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *PendingPaymentListV1FvEmeaOK) GetPayload() *models.PendingPaymentListResponse {
	return o.Payload
}

func (o *PendingPaymentListV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PendingPaymentListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPendingPaymentListV1FvEmeaUnauthorized creates a PendingPaymentListV1FvEmeaUnauthorized with default headers values
func NewPendingPaymentListV1FvEmeaUnauthorized() *PendingPaymentListV1FvEmeaUnauthorized {
	return &PendingPaymentListV1FvEmeaUnauthorized{}
}

/*PendingPaymentListV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized
*/
type PendingPaymentListV1FvEmeaUnauthorized struct {
	Payload *models.Error
}

func (o *PendingPaymentListV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *PendingPaymentListV1FvEmeaUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PendingPaymentListV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPendingPaymentListV1FvEmeaForbidden creates a PendingPaymentListV1FvEmeaForbidden with default headers values
func NewPendingPaymentListV1FvEmeaForbidden() *PendingPaymentListV1FvEmeaForbidden {
	return &PendingPaymentListV1FvEmeaForbidden{}
}

/*PendingPaymentListV1FvEmeaForbidden handles this case with default header values.

Forbidden
*/
type PendingPaymentListV1FvEmeaForbidden struct {
	Payload *models.Error
}

func (o *PendingPaymentListV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *PendingPaymentListV1FvEmeaForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PendingPaymentListV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPendingPaymentListV1FvEmeaStatus452 creates a PendingPaymentListV1FvEmeaStatus452 with default headers values
func NewPendingPaymentListV1FvEmeaStatus452() *PendingPaymentListV1FvEmeaStatus452 {
	return &PendingPaymentListV1FvEmeaStatus452{}
}

/*PendingPaymentListV1FvEmeaStatus452 handles this case with default header values.

System Exception
*/
type PendingPaymentListV1FvEmeaStatus452 struct {
	Payload *models.Error
}

func (o *PendingPaymentListV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaStatus452  %+v", 452, o.Payload)
}

func (o *PendingPaymentListV1FvEmeaStatus452) GetPayload() *models.Error {
	return o.Payload
}

func (o *PendingPaymentListV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPendingPaymentListV1FvEmeaStatus453 creates a PendingPaymentListV1FvEmeaStatus453 with default headers values
func NewPendingPaymentListV1FvEmeaStatus453() *PendingPaymentListV1FvEmeaStatus453 {
	return &PendingPaymentListV1FvEmeaStatus453{}
}

/*PendingPaymentListV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type PendingPaymentListV1FvEmeaStatus453 struct {
}

func (o *PendingPaymentListV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaStatus453 ", 453)
}

func (o *PendingPaymentListV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPendingPaymentListV1FvEmeaStatus455 creates a PendingPaymentListV1FvEmeaStatus455 with default headers values
func NewPendingPaymentListV1FvEmeaStatus455() *PendingPaymentListV1FvEmeaStatus455 {
	return &PendingPaymentListV1FvEmeaStatus455{}
}

/*PendingPaymentListV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type PendingPaymentListV1FvEmeaStatus455 struct {
}

func (o *PendingPaymentListV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaStatus455 ", 455)
}

func (o *PendingPaymentListV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPendingPaymentListV1FvEmeaStatus467 creates a PendingPaymentListV1FvEmeaStatus467 with default headers values
func NewPendingPaymentListV1FvEmeaStatus467() *PendingPaymentListV1FvEmeaStatus467 {
	return &PendingPaymentListV1FvEmeaStatus467{}
}

/*PendingPaymentListV1FvEmeaStatus467 handles this case with default header values.

List of possible error codes that could appear within the field result code:<BR/><BR/> E001-Unsuccessful
*/
type PendingPaymentListV1FvEmeaStatus467 struct {
	Payload *models.Error
}

func (o *PendingPaymentListV1FvEmeaStatus467) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/pendingPaymentList][%d] pendingPaymentListV1FvEmeaStatus467  %+v", 467, o.Payload)
}

func (o *PendingPaymentListV1FvEmeaStatus467) GetPayload() *models.Error {
	return o.Payload
}

func (o *PendingPaymentListV1FvEmeaStatus467) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}