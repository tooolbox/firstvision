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

// AddressHistoryInquiryResponse address history inquiry response
//
// swagger:model AddressHistoryInquiryResponse
type AddressHistoryInquiryResponse struct {

	//  Max length = 19, Account Number: Number of Customer's account.
	AcctNbr string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Customer number: Customer number that identifies the Customer Name/Address record to which this account is assigned.
	CustNbr string `json:"custNbr,omitempty"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	// name address tbl
	NameAddressTbl []*NameAddressTblForAddressHistoryInquiry1 `json:"nameAddressTbl"`

	//  Max length = 2, Number of occurrences for prior mailer fields
	NbrOfOccurs string `json:"nbrOfOccurs,omitempty"`

	//  Max length = 2, Number of occurrences for prior mailer fields
	PmNbrOccur string `json:"pmNbrOccur,omitempty"`

	// Prior Mailing Name and Address Group variable which occurs 10 times.
	PmTbl []*PmTblForAddressHistoryInquiry1 `json:"pmTbl"`

	//  Max length = 19, Relationship number that indicates the number of the relationship within which this account resides.
	RelNbr string `json:"relNbr,omitempty"`

	//  Max length = 1, Flag to indicate the VIP customer.
	VipIndc string `json:"vipIndc,omitempty"`
}

// Validate validates this address history inquiry response
func (m *AddressHistoryInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameAddressTbl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmTbl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressHistoryInquiryResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *AddressHistoryInquiryResponse) validateNameAddressTbl(formats strfmt.Registry) error {

	if swag.IsZero(m.NameAddressTbl) { // not required
		return nil
	}

	for i := 0; i < len(m.NameAddressTbl); i++ {
		if swag.IsZero(m.NameAddressTbl[i]) { // not required
			continue
		}

		if m.NameAddressTbl[i] != nil {
			if err := m.NameAddressTbl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nameAddressTbl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressHistoryInquiryResponse) validatePmTbl(formats strfmt.Registry) error {

	if swag.IsZero(m.PmTbl) { // not required
		return nil
	}

	for i := 0; i < len(m.PmTbl); i++ {
		if swag.IsZero(m.PmTbl[i]) { // not required
			continue
		}

		if m.PmTbl[i] != nil {
			if err := m.PmTbl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pmTbl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressHistoryInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressHistoryInquiryResponse) UnmarshalBinary(b []byte) error {
	var res AddressHistoryInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}