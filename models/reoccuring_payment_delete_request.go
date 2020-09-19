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

// ReoccuringPaymentDeleteRequest reoccuring payment delete request
//
// swagger:model reoccuringPaymentDeleteRequest
type ReoccuringPaymentDeleteRequest struct {

	// common
	Common *Header `json:"common,omitempty"`

	// Populate with the CPA id of the corresponding recurring payment plan.
	// Required: true
	CpaID *string `json:"cpaId"`

	// Populate with 5 or 6. 5 = Cancelled by client, 6 = Cancelled by customer.
	// Required: true
	StatusID *string `json:"statusId"`

	// Populate with the FirstPay userID to target a specific region. This is a mandatory field.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this reoccuring payment delete request
func (m *ReoccuringPaymentDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCpaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReoccuringPaymentDeleteRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *ReoccuringPaymentDeleteRequest) validateCpaID(formats strfmt.Registry) error {

	if err := validate.Required("cpaId", "body", m.CpaID); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentDeleteRequest) validateStatusID(formats strfmt.Registry) error {

	if err := validate.Required("statusId", "body", m.StatusID); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentDeleteRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReoccuringPaymentDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReoccuringPaymentDeleteRequest) UnmarshalBinary(b []byte) error {
	var res ReoccuringPaymentDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}