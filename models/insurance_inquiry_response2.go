// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InsuranceInquiryResponse2 insurance inquiry response2
//
// swagger:model InsuranceInquiryResponse2
type InsuranceInquiryResponse2 struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	Account string `json:"account,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Unique Card Number embossed on card plastic.
	Crd string `json:"crd,omitempty"`

	//  Max length = 3, Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	CurrCd string `json:"currCd,omitempty"`

	//  Max length = 1, Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	CurrNod string `json:"currNod,omitempty"`

	//  Max length = 1, Dual Indicator: Dual Account Flag of account. Indicator that displays only if this is a dual currency account. If this is not a dual currency account, the indicator does not display. If this is a dual currency account, a duplicate Account Base Segment record resides in the associated dual currency Organisation. Values are: L = Local account (default) F = Foreign account
	DualInd string `json:"dualInd,omitempty"`

	//  Max length = 3, Foreign Organisation Number: Three digit Identification number of the Organisation
	FgnOrg string `json:"fgnOrg,omitempty"`

	// Up to 6 occurrences insurance data.
	InsData []*InsData2ForInsuranceInquiry2 `json:"insData"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	//  Max length = 2, Number of insurance products returned.
	NbrInsPrd string `json:"nbrInsPrd,omitempty"`

	//  Max length = 1, Percentage NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in percentage fields.
	PercentNod string `json:"percentNod,omitempty"`
}

// Validate validates this insurance inquiry response2
func (m *InsuranceInquiryResponse2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInsData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InsuranceInquiryResponse2) validateCommon(formats strfmt.Registry) error {

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

func (m *InsuranceInquiryResponse2) validateInsData(formats strfmt.Registry) error {

	if swag.IsZero(m.InsData) { // not required
		return nil
	}

	for i := 0; i < len(m.InsData); i++ {
		if swag.IsZero(m.InsData[i]) { // not required
			continue
		}

		if m.InsData[i] != nil {
			if err := m.InsData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("insData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InsuranceInquiryResponse2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InsuranceInquiryResponse2) UnmarshalBinary(b []byte) error {
	var res InsuranceInquiryResponse2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}