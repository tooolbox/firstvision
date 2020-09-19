// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CardLostAndStolenServiceRequest card lost and stolen service request
//
// swagger:model CardLostAndStolenServiceRequest
type CardLostAndStolenServiceRequest struct {

	//  Max length = 1, Block code to apply to the transfer card. The values are blank (or space) and A to Z.  Edit checks: The block code must be defined on the logo record and be designated lost/stolen.  When the value is blank, then the  embosser records is cross checked to see if a block code is already present on the card. If a block code is not present or is not defined as lost/stolen/fraud then an error message is returned.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	BlockCode *string `json:"blockCode"`

	//  Max length = 1, Action to be take on the card. Valid values: T = Transfer the card(lost, Stolen, or Fraud) R = Reverse a card transfer  Cannot be space or Low-Values
	// Required: true
	// Max Length: 1
	// Min Length: 0
	CardActn *string `json:"cardActn"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr"`

	//  Max length = 1, CARD PROCESS TYPE: Code that indicates whether the cards is processed during night batch or  in the same day embossing processing cycle.   Valid values:  0 - Same-day plastic processing is not required (default) 1 - Same-day plastic processing is required. Space - Default from logo SDP settings.
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardProcType *string `json:"cardProcType,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Fraud Transfer Effective Date: Date the fraud card transfer is effective on the card. The date value must be a less or equal to today's date. Default is today's date and cannot accept any future dates. Valid value has to be entered in this field. 1.Must be entered, cannot be zero.  2. Format is CCYY-MM-DD
	// Required: true
	EffectiveDate *string `json:"effectiveDate"`

	//  Max length = 3, Loss country: Country code for the area where the card was lost. Must be valid ISO alphanumeric country code.
	// Max Length: 3
	// Min Length: 0
	LossCountry *string `json:"lossCountry,omitempty"`

	// Format: YYYYMMDD. Date the card was lost, stolen or the first date of fraudulent activity. The date must be less than or equal to today's date. If the date is not entered in this field,this field defaults to the current system date. Format is CCYY-MM-DD
	LossDate string `json:"lossDate,omitempty"`

	//  Max length = 1, Loss location: Code that indicates where the cards was lost. Valid values: Space = Not used (Default) 1 = Home 2 = Auto 3 = Office 4 = Store 5 = Rural 6 = Bar 7 = Hotel or Motel 8 = Unknown A = Airport B = Port.
	// Max Length: 1
	// Min Length: 0
	LossLoc *string `json:"lossLoc,omitempty"`

	//  Max length = 1, Loss type: Code that identifies the reason for lost or stolen card. Valid values: 0 to 9
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	LossType *string `json:"lossType,omitempty"`

	//  Max length = 1, Pin Compromised:  Code that indicates if the customer's PIN  was compromised. Valid values: 0 - No (Default) 1 - Yes When the values is set to 1 a new PIN is automatically generated for the card reported lost or stolen.
	// Max Length: 1
	// Min Length: 0
	PinComp *string `json:"pinComp,omitempty"`

	//  Max length = 33, Police reference number: Loss property number or police crime reference number assigned to the card when this card is reported lost or stolen.
	// Max Length: 33
	// Min Length: 0
	PoliceRefNbr *string `json:"policeRefNbr,omitempty"`

	//  Max length = 1, Transfer replacement indicator: To indicate whether to replace the transfer card  by generating  a new embosser record. Valid values: 0 - Replace lost/stolen/fraud card (default) 1 - Do not replace lost/stolen/fraud card
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	XfrReplacementInd *string `json:"xfrReplacementInd,omitempty"`

	//  Max length = 19, Transfer to account number.
	// Max Length: 19
	// Min Length: 0
	XfrToAcct *string `json:"xfrToAcct,omitempty"`

	//  Max length = 19, Transfer to customer number
	// Max Length: 19
	// Min Length: 0
	XfrToCust *string `json:"xfrToCust,omitempty"`
}

// Validate validates this card lost and stolen service request
func (m *CardLostAndStolenServiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardActn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossLoc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinComp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoliceRefNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXfrReplacementInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXfrToAcct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXfrToCust(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardLostAndStolenServiceRequest) validateBlockCode(formats strfmt.Registry) error {

	if err := validate.Required("blockCode", "body", m.BlockCode); err != nil {
		return err
	}

	if err := validate.MinLength("blockCode", "body", string(*m.BlockCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("blockCode", "body", string(*m.BlockCode), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateCardActn(formats strfmt.Registry) error {

	if err := validate.Required("cardActn", "body", m.CardActn); err != nil {
		return err
	}

	if err := validate.MinLength("cardActn", "body", string(*m.CardActn), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardActn", "body", string(*m.CardActn), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateCardNbr(formats strfmt.Registry) error {

	if err := validate.Required("cardNbr", "body", m.CardNbr); err != nil {
		return err
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateCardProcType(formats strfmt.Registry) error {

	if swag.IsZero(m.CardProcType) { // not required
		return nil
	}

	if err := validate.MinLength("cardProcType", "body", string(*m.CardProcType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardProcType", "body", string(*m.CardProcType), 1); err != nil {
		return err
	}

	if err := validate.Pattern("cardProcType", "body", string(*m.CardProcType), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.Common) { // not required
		return nil
	}

	if m.Common != nil {
		if err := m.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common")
			}
			return err
		}
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateEffectiveDate(formats strfmt.Registry) error {

	if err := validate.Required("effectiveDate", "body", m.EffectiveDate); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateLossCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.LossCountry) { // not required
		return nil
	}

	if err := validate.MinLength("lossCountry", "body", string(*m.LossCountry), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lossCountry", "body", string(*m.LossCountry), 3); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateLossLoc(formats strfmt.Registry) error {

	if swag.IsZero(m.LossLoc) { // not required
		return nil
	}

	if err := validate.MinLength("lossLoc", "body", string(*m.LossLoc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lossLoc", "body", string(*m.LossLoc), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateLossType(formats strfmt.Registry) error {

	if swag.IsZero(m.LossType) { // not required
		return nil
	}

	if err := validate.MinLength("lossType", "body", string(*m.LossType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lossType", "body", string(*m.LossType), 1); err != nil {
		return err
	}

	if err := validate.Pattern("lossType", "body", string(*m.LossType), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validatePinComp(formats strfmt.Registry) error {

	if swag.IsZero(m.PinComp) { // not required
		return nil
	}

	if err := validate.MinLength("pinComp", "body", string(*m.PinComp), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pinComp", "body", string(*m.PinComp), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validatePoliceRefNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.PoliceRefNbr) { // not required
		return nil
	}

	if err := validate.MinLength("policeRefNbr", "body", string(*m.PoliceRefNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("policeRefNbr", "body", string(*m.PoliceRefNbr), 33); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateXfrReplacementInd(formats strfmt.Registry) error {

	if swag.IsZero(m.XfrReplacementInd) { // not required
		return nil
	}

	if err := validate.MinLength("xfrReplacementInd", "body", string(*m.XfrReplacementInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("xfrReplacementInd", "body", string(*m.XfrReplacementInd), 1); err != nil {
		return err
	}

	if err := validate.Pattern("xfrReplacementInd", "body", string(*m.XfrReplacementInd), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateXfrToAcct(formats strfmt.Registry) error {

	if swag.IsZero(m.XfrToAcct) { // not required
		return nil
	}

	if err := validate.MinLength("xfrToAcct", "body", string(*m.XfrToAcct), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("xfrToAcct", "body", string(*m.XfrToAcct), 19); err != nil {
		return err
	}

	return nil
}

func (m *CardLostAndStolenServiceRequest) validateXfrToCust(formats strfmt.Registry) error {

	if swag.IsZero(m.XfrToCust) { // not required
		return nil
	}

	if err := validate.MinLength("xfrToCust", "body", string(*m.XfrToCust), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("xfrToCust", "body", string(*m.XfrToCust), 19); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardLostAndStolenServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardLostAndStolenServiceRequest) UnmarshalBinary(b []byte) error {
	var res CardLostAndStolenServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}