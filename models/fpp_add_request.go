// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FppAddRequest fpp add request
//
// swagger:model FppAddRequest
type FppAddRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	// Max Length: 19
	// Min Length: 0
	Acct *string `json:"acct,omitempty"`

	//  Max length = 17, Installment Amount. This is the fixed instalment amount that customer will have to pay.
	// Pattern: ^(-)?[0-9]{1,17}$
	AgrdPmtAmt string `json:"agrdPmtAmt,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Sum of  current balance of all flexible payment  processing plans on the account.
	// Pattern: ^(-)?[0-9]{1,17}$
	FppBal string `json:"fppBal,omitempty"`

	//  Max length = 1, Valid values are space, P, or N Field is used for scrolling. Front-end interface system calling service will populate this field from the previous output message when more data that meet the search criteria is available.
	// Max Length: 1
	// Min Length: 0
	Function *string `json:"function,omitempty"`

	//  Max length = 5, Credit Plan Number: Identifies the plan number of the Credit Plan Master record associated with the Credit Plan Segment record.
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Plan *string `json:"plan,omitempty"`

	// Format: YYYYMMDD. Start Date.  1. Must be a numeric value and a valid julian date/Cannot be 0.  2. Cannot be > the next processing date and  cannot be < AMBS open date.  3. Must be a valid date. For LOC plan types: - if date is not given, Then the Today's Julian date is populated.
	PlanDate string `json:"planDate,omitempty"`

	//  Max length = 40, Credit Plan Description: Description of the FPP plan selected by the cardholder. You cannot edit this field once you enter the description. If you do not enter a description, the description available in the Credit Plan Segment record will be updated for this FPP plan.
	// Max Length: 40
	// Min Length: 0
	PlanDesc *string `json:"planDesc,omitempty"`

	// plan group
	PlanGroup []*PlanGroupForFppAdd1 `json:"planGroup"`

	//  Max length = 3, Credit Plan Data Record Number: Sequence number of Credit Plan. This field value may be changed by CMS during the daily batch process if a plan segment is purged. The system will re-sequence all the plan segments on file for the account.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	RecNbr *string `json:"recNbr,omitempty"`

	//  Max length = 3, Term: Original term of the FPP plan
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Term *string `json:"term,omitempty"`
}

// Validate validates this fpp add request
func (m *FppAddRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgrdPmtAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppBal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FppAddRequest) validateAcct(formats strfmt.Registry) error {

	if swag.IsZero(m.Acct) { // not required
		return nil
	}

	if err := validate.MinLength("acct", "body", string(*m.Acct), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acct", "body", string(*m.Acct), 19); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validateAgrdPmtAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.AgrdPmtAmt) { // not required
		return nil
	}

	if err := validate.Pattern("agrdPmtAmt", "body", string(m.AgrdPmtAmt), `^(-)?[0-9]{1,17}$`); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *FppAddRequest) validateFppBal(formats strfmt.Registry) error {

	if swag.IsZero(m.FppBal) { // not required
		return nil
	}

	if err := validate.Pattern("fppBal", "body", string(m.FppBal), `^(-)?[0-9]{1,17}$`); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validateFunction(formats strfmt.Registry) error {

	if swag.IsZero(m.Function) { // not required
		return nil
	}

	if err := validate.MinLength("function", "body", string(*m.Function), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("function", "body", string(*m.Function), 1); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if err := validate.MinLength("plan", "body", string(*m.Plan), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("plan", "body", string(*m.Plan), 5); err != nil {
		return err
	}

	if err := validate.Pattern("plan", "body", string(*m.Plan), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validatePlanDesc(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanDesc) { // not required
		return nil
	}

	if err := validate.MinLength("planDesc", "body", string(*m.PlanDesc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("planDesc", "body", string(*m.PlanDesc), 40); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validatePlanGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.PlanGroup); i++ {
		if swag.IsZero(m.PlanGroup[i]) { // not required
			continue
		}

		if m.PlanGroup[i] != nil {
			if err := m.PlanGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("planGroup" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FppAddRequest) validateRecNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.RecNbr) { // not required
		return nil
	}

	if err := validate.MinLength("recNbr", "body", string(*m.RecNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("recNbr", "body", string(*m.RecNbr), 3); err != nil {
		return err
	}

	if err := validate.Pattern("recNbr", "body", string(*m.RecNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *FppAddRequest) validateTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.Term) { // not required
		return nil
	}

	if err := validate.MinLength("term", "body", string(*m.Term), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("term", "body", string(*m.Term), 3); err != nil {
		return err
	}

	if err := validate.Pattern("term", "body", string(*m.Term), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FppAddRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FppAddRequest) UnmarshalBinary(b []byte) error {
	var res FppAddRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}