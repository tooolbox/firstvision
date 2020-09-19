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

// StatementDateInquiryRequest statement date inquiry request
//
// swagger:model StatementDateInquiryRequest
type StatementDateInquiryRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Dual Indicator: Dual Account Flag of account. Valid values are: L = Local account (default) F = Foreign account If dual currency is not being used or if the field is left blank, the value is L
	// Max Length: 1
	// Min Length: 0
	DualInd *string `json:"dualInd,omitempty"`
}

// Validate validates this statement date inquiry request
func (m *StatementDateInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualInd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatementDateInquiryRequest) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	if err := validate.MinLength("account", "body", string(*m.Account), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("account", "body", string(*m.Account), 19); err != nil {
		return err
	}

	return nil
}

func (m *StatementDateInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *StatementDateInquiryRequest) validateDualInd(formats strfmt.Registry) error {

	if swag.IsZero(m.DualInd) { // not required
		return nil
	}

	if err := validate.MinLength("dualInd", "body", string(*m.DualInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("dualInd", "body", string(*m.DualInd), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatementDateInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatementDateInquiryRequest) UnmarshalBinary(b []byte) error {
	var res StatementDateInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}