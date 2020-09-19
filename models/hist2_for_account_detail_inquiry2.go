// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Hist2ForAccountDetailInquiry2 hist2 for account detail inquiry2
//
// swagger:model Hist2ForAccountDetailInquiry2
type Hist2ForAccountDetailInquiry2 struct {

	//  Max length = 1, Payment history counter that indicates the payment performance or balance condition of the account for the previous billing cycle.
	GetDollarHBal string `json:"get$HBal,omitempty"`
}

// Validate validates this hist2 for account detail inquiry2
func (m *Hist2ForAccountDetailInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Hist2ForAccountDetailInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hist2ForAccountDetailInquiry2) UnmarshalBinary(b []byte) error {
	var res Hist2ForAccountDetailInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}