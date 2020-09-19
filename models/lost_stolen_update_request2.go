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

// LostStolenUpdateRequest2 lost stolen update request2
//
// swagger:model LostStolenUpdateRequest2
type LostStolenUpdateRequest2 struct {

	//  Max length = 1, Card Block Code: Block code assigned to this card.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	CardBlk *string `json:"cardBlk"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, Country Of Loss: Country code for the area where the card was lost.
	// Required: true
	// Max Length: 3
	// Min Length: 0
	CountryOfLoss *string `json:"countryOfLoss"`

	// Format: YYYYMMDD. Loss Date: Date the card was lost, stolen or the first date of fraudulent activity.  The date must be less than or equal to today's date.  The date field will default to the current system date if no date is entered. Date format is CCYYMMDD.
	// Required: true
	LossDate *string `json:"lossDate"`

	//  Max length = 1, Location Of Loss: This field further defines where the card was lost. Valid values are: Space (default) 1 - Home 2 - Auto 3 - Office 4 - Store 5 - Rural 6 - Bar 7 - Hotel or Motel 8 - Unknown A - Airport B - Port
	// Required: true
	// Max Length: 1
	// Min Length: 0
	LossLocation *string `json:"lossLocation"`

	//  Max length = 1, Type Of Loss: Identifies the reason for the lost or stolen card. Values are: 0 = Lost 1 = Stolen 2 = Not received 3 = Fraudulent application 4 = Counterfeit 5 = Account takeover 6 = Mail order / Internet 7 = First party fraud 8 = Other Default is space.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	LossType *string `json:"lossType"`

	//  Max length = 1, New Card: Field to indicate if a new card is issued to the customer.  Values are: 0 = Yes (Default) 1 = No
	// Required: true
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	NewCard *string `json:"newCard"`

	//  Max length = 1, Pin Compromised Flag: Flag indicating if the customer's PIN was compromised.  Values are: 0 = No 1 = Yes An entry of 1 in this field will automatically generate a new PIN for the card being reported lost or stolen.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	PinCompromised *string `json:"pinCompromised"`

	//  Max length = 33, Police reference Number: field used to record either the Loss Property Number or Police Crime Reference Number if either is assigned by law enforcement.
	// Max Length: 33
	// Min Length: 0
	PoliceRefNbr *string `json:"policeRefNbr,omitempty"`

	//  Max length = 1, Process Type: Field to indicate if Same Day Plastics processing should be used for the new plastic request.  The value displayed in this field is defaulted from the Logo record but may be changed by entering a value. Values are: 0 = No (Default) 1 = Yes
	// Required: true
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ProcessType *string `json:"processType"`
}

// Validate validates this lost stolen update request2
func (m *LostStolenUpdateRequest2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardBlk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryOfLoss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLossType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinCompromised(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoliceRefNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LostStolenUpdateRequest2) validateCardBlk(formats strfmt.Registry) error {

	if err := validate.Required("cardBlk", "body", m.CardBlk); err != nil {
		return err
	}

	if err := validate.MinLength("cardBlk", "body", string(*m.CardBlk), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardBlk", "body", string(*m.CardBlk), 1); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validateCardNbr(formats strfmt.Registry) error {

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

func (m *LostStolenUpdateRequest2) validateCommon(formats strfmt.Registry) error {

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

func (m *LostStolenUpdateRequest2) validateCountryOfLoss(formats strfmt.Registry) error {

	if err := validate.Required("countryOfLoss", "body", m.CountryOfLoss); err != nil {
		return err
	}

	if err := validate.MinLength("countryOfLoss", "body", string(*m.CountryOfLoss), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("countryOfLoss", "body", string(*m.CountryOfLoss), 3); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validateLossDate(formats strfmt.Registry) error {

	if err := validate.Required("lossDate", "body", m.LossDate); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validateLossLocation(formats strfmt.Registry) error {

	if err := validate.Required("lossLocation", "body", m.LossLocation); err != nil {
		return err
	}

	if err := validate.MinLength("lossLocation", "body", string(*m.LossLocation), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lossLocation", "body", string(*m.LossLocation), 1); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validateLossType(formats strfmt.Registry) error {

	if err := validate.Required("lossType", "body", m.LossType); err != nil {
		return err
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

func (m *LostStolenUpdateRequest2) validateNewCard(formats strfmt.Registry) error {

	if err := validate.Required("newCard", "body", m.NewCard); err != nil {
		return err
	}

	if err := validate.MinLength("newCard", "body", string(*m.NewCard), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("newCard", "body", string(*m.NewCard), 1); err != nil {
		return err
	}

	if err := validate.Pattern("newCard", "body", string(*m.NewCard), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validatePinCompromised(formats strfmt.Registry) error {

	if err := validate.Required("pinCompromised", "body", m.PinCompromised); err != nil {
		return err
	}

	if err := validate.MinLength("pinCompromised", "body", string(*m.PinCompromised), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pinCompromised", "body", string(*m.PinCompromised), 1); err != nil {
		return err
	}

	if err := validate.Pattern("pinCompromised", "body", string(*m.PinCompromised), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LostStolenUpdateRequest2) validatePoliceRefNbr(formats strfmt.Registry) error {

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

func (m *LostStolenUpdateRequest2) validateProcessType(formats strfmt.Registry) error {

	if err := validate.Required("processType", "body", m.ProcessType); err != nil {
		return err
	}

	if err := validate.MinLength("processType", "body", string(*m.ProcessType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processType", "body", string(*m.ProcessType), 1); err != nil {
		return err
	}

	if err := validate.Pattern("processType", "body", string(*m.ProcessType), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LostStolenUpdateRequest2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LostStolenUpdateRequest2) UnmarshalBinary(b []byte) error {
	var res LostStolenUpdateRequest2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}