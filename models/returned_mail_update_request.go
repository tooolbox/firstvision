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

// ReturnedMailUpdateRequest returned mail update request
//
// swagger:model ReturnedMailUpdateRequest
type ReturnedMailUpdateRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 2, Number of times mail was returned.
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	RtnMailCtr *string `json:"rtnMailCtr,omitempty"`

	// Format: YYYYMMDD. Field that indicates the date of the previously returned mail.
	RtnMailDate string `json:"rtnMailDate,omitempty"`

	//  Max length = 3, Field that indicates a user-defined reason code that identifies the type of mail previously returned.
	// Max Length: 3
	// Min Length: 0
	RtnMailUser *string `json:"rtnMailUser,omitempty"`
}

// Validate validates this returned mail update request
func (m *ReturnedMailUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtnMailCtr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtnMailUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnedMailUpdateRequest) validateAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *ReturnedMailUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *ReturnedMailUpdateRequest) validateRtnMailCtr(formats strfmt.Registry) error {

	if swag.IsZero(m.RtnMailCtr) { // not required
		return nil
	}

	if err := validate.MinLength("rtnMailCtr", "body", string(*m.RtnMailCtr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rtnMailCtr", "body", string(*m.RtnMailCtr), 2); err != nil {
		return err
	}

	if err := validate.Pattern("rtnMailCtr", "body", string(*m.RtnMailCtr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnedMailUpdateRequest) validateRtnMailUser(formats strfmt.Registry) error {

	if swag.IsZero(m.RtnMailUser) { // not required
		return nil
	}

	if err := validate.MinLength("rtnMailUser", "body", string(*m.RtnMailUser), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rtnMailUser", "body", string(*m.RtnMailUser), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnedMailUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnedMailUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ReturnedMailUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}