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

// SettlementQuoteUpdateResponse settlement quote update response
//
// swagger:model SettlementQuoteUpdateResponse
type SettlementQuoteUpdateResponse struct {

	// Format: YYYYMMDD. Account level quote expiry Date
	AcctExpDate string `json:"acctExpDate,omitempty"`

	//  Max length = 17, Loan Settlement Amount
	AcctLvlPayoff string `json:"acctLvlPayoff,omitempty"`

	//  Max length = 1, Type of account level quote requested.
	AcctLvlQt string `json:"acctLvlQt,omitempty"`

	//  Max length = 19, Account Number: Identification Number of Customer's account for which the settlement quote update requested.
	AcctNbr string `json:"acctNbr,omitempty"`

	//  Max length = 1, Code that indicates whether more records are on file. Values are: N = No more records on file Y = More records on file
	AqMoreInd string `json:"aqMoreInd,omitempty"`

	//  Max length = 19, Card number for which settlement quote  is required.
	CardNbr string `json:"cardNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	CurrCode string `json:"currCode,omitempty"`

	//  Max length = 1, Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	CurrNod string `json:"currNod,omitempty"`

	//  Max length = 3, Foreign Organization: Three digit Identification number of the Foreign organization.
	FrgnOrg string `json:"frgnOrg,omitempty"`

	//  Max length = 1, Foreign Use value received from input.
	FrgnUse string `json:"frgnUse,omitempty"`

	// loan data
	LoanData []*LoanDataForSettlementQuoteUpdate1 `json:"loanData"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	// Format: YYYYMMDD. Pay-off date/loan settlement date used for account level generation quote.
	PayoffByDt string `json:"payoffByDt,omitempty"`

	//  Max length = 2, Number of plans present for that particular account
	PlanNbrOccurs string `json:"planNbrOccurs,omitempty"`

	//  Max length = 3, Start plan sequence number associated with the scroll.
	SvcStartPlanSeq string `json:"svcStartPlanSeq,omitempty"`
}

// Validate validates this settlement quote update response
func (m *SettlementQuoteUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoanData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettlementQuoteUpdateResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *SettlementQuoteUpdateResponse) validateLoanData(formats strfmt.Registry) error {

	if swag.IsZero(m.LoanData) { // not required
		return nil
	}

	for i := 0; i < len(m.LoanData); i++ {
		if swag.IsZero(m.LoanData[i]) { // not required
			continue
		}

		if m.LoanData[i] != nil {
			if err := m.LoanData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("loanData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettlementQuoteUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettlementQuoteUpdateResponse) UnmarshalBinary(b []byte) error {
	var res SettlementQuoteUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}