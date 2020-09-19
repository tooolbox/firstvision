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

// AccountBlockCodeUpdateRequest account block code update request
//
// swagger:model AccountBlockCodeUpdateRequest
type AccountBlockCodeUpdateRequest struct {

	//  Max length = 1, Account Restructure Flag: Flag that indicates whether the account should be restructured.  Values are: Y = Restructure Requested N = No Account Restructure Requested
	// Max Length: 1
	// Min Length: 0
	AcctRestructure *string `json:"acctRestructure,omitempty"`

	//  Max length = 1, Block Code: Single letter alphanumeric code which decide the block code, part of block matrix. 1. Values are A-Z and space 2. Must be different from space if FUNCTION = B (Block) 3. Must equal space if FUNCTION = U (Unblock)
	// Max Length: 1
	// Min Length: 0
	BlkCd *string `json:"blkCd,omitempty"`

	//  Max length = 1, Block Code Indicator.  Values are: 0 (Default) - The Service will check BLOCK CODE 1 and BLOCK 2 and update either field based on the block code priority defined on the Block Code Matrix 1 - Update BLOCK CODE 1 only based on the block code priority defined on the Block Code Matrix. 2 - Update BLOCK CODE 2 only based on the block code priority defined on the Block Code Matrix. When block code fields are filled, only a block code with higher priority can replace an existing block code.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	BlkCdInd *string `json:"blkCdInd"`

	//  Max length = 2, BLOCK CODE 1 REASON
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	BlockCode1Reason *string `json:"blockCode1Reason,omitempty"`

	//  Max length = 2, BLOCK CODE 2 REASON
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	BlockCode2Reason *string `json:"blockCode2Reason,omitempty"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Foreign use indicator: Value indicates whether the incoming action is applied to the local or foreign account. Values are: 0 = Local account (default) 1 = Foreign account SPACE = defaults to 0
	// Max Length: 1
	// Min Length: 0
	ForeignUse *string `json:"foreignUse,omitempty"`

	//  Max length = 1, Function code: Function to be performed as part of the service request. Values are: B (Block) U (Unblock)
	// Required: true
	// Max Length: 1
	// Min Length: 0
	FunctionCd *string `json:"functionCd"`

	//  Max length = 1, User Defined.
	// Max Length: 1
	// Min Length: 0
	ProcessFlg1 *string `json:"processFlg1,omitempty"`

	//  Max length = 1, User Defined.
	// Max Length: 1
	// Min Length: 0
	ProcessFlg2 *string `json:"processFlg2,omitempty"`

	//  Max length = 1, User Defined.
	// Max Length: 1
	// Min Length: 0
	ProcessFlg3 *string `json:"processFlg3,omitempty"`

	//  Max length = 1, User Defined.
	// Max Length: 1
	// Min Length: 0
	ProcessFlg4 *string `json:"processFlg4,omitempty"`

	//  Max length = 1, User Defined.
	// Max Length: 1
	// Min Length: 0
	ProcessFlg5 *string `json:"processFlg5,omitempty"`

	//  Max length = 2, Account Type/Product Code.  User defined.  Used as a Variable Decision Key in TRIAD.
	// Max Length: 2
	// Min Length: 0
	ProdCodeAccType *string `json:"prodCodeAccType,omitempty"`

	//  Max length = 2, User Defined Sub-Status.  Works in conjunction with Block Code 1 to identify reason for Block code being set.  Variable Decision Key used in TRIAD.
	// Max Length: 2
	// Min Length: 0
	SubStatus *string `json:"subStatus,omitempty"`
}

// Validate validates this account block code update request
func (m *AccountBlockCodeUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctRestructure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkCd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkCdInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockCode1Reason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockCode2Reason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionCd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessFlg1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessFlg2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessFlg3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessFlg4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessFlg5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProdCodeAccType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateAcctRestructure(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctRestructure) { // not required
		return nil
	}

	if err := validate.MinLength("acctRestructure", "body", string(*m.AcctRestructure), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctRestructure", "body", string(*m.AcctRestructure), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateBlkCd(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkCd) { // not required
		return nil
	}

	if err := validate.MinLength("blkCd", "body", string(*m.BlkCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("blkCd", "body", string(*m.BlkCd), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateBlkCdInd(formats strfmt.Registry) error {

	if err := validate.Required("blkCdInd", "body", m.BlkCdInd); err != nil {
		return err
	}

	if err := validate.MinLength("blkCdInd", "body", string(*m.BlkCdInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("blkCdInd", "body", string(*m.BlkCdInd), 1); err != nil {
		return err
	}

	if err := validate.Pattern("blkCdInd", "body", string(*m.BlkCdInd), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateBlockCode1Reason(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockCode1Reason) { // not required
		return nil
	}

	if err := validate.MinLength("blockCode1Reason", "body", string(*m.BlockCode1Reason), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("blockCode1Reason", "body", string(*m.BlockCode1Reason), 2); err != nil {
		return err
	}

	if err := validate.Pattern("blockCode1Reason", "body", string(*m.BlockCode1Reason), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateBlockCode2Reason(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockCode2Reason) { // not required
		return nil
	}

	if err := validate.MinLength("blockCode2Reason", "body", string(*m.BlockCode2Reason), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("blockCode2Reason", "body", string(*m.BlockCode2Reason), 2); err != nil {
		return err
	}

	if err := validate.Pattern("blockCode2Reason", "body", string(*m.BlockCode2Reason), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateCardNbr(formats strfmt.Registry) error {

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

func (m *AccountBlockCodeUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *AccountBlockCodeUpdateRequest) validateForeignUse(formats strfmt.Registry) error {

	if swag.IsZero(m.ForeignUse) { // not required
		return nil
	}

	if err := validate.MinLength("foreignUse", "body", string(*m.ForeignUse), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("foreignUse", "body", string(*m.ForeignUse), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateFunctionCd(formats strfmt.Registry) error {

	if err := validate.Required("functionCd", "body", m.FunctionCd); err != nil {
		return err
	}

	if err := validate.MinLength("functionCd", "body", string(*m.FunctionCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("functionCd", "body", string(*m.FunctionCd), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProcessFlg1(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessFlg1) { // not required
		return nil
	}

	if err := validate.MinLength("processFlg1", "body", string(*m.ProcessFlg1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processFlg1", "body", string(*m.ProcessFlg1), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProcessFlg2(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessFlg2) { // not required
		return nil
	}

	if err := validate.MinLength("processFlg2", "body", string(*m.ProcessFlg2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processFlg2", "body", string(*m.ProcessFlg2), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProcessFlg3(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessFlg3) { // not required
		return nil
	}

	if err := validate.MinLength("processFlg3", "body", string(*m.ProcessFlg3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processFlg3", "body", string(*m.ProcessFlg3), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProcessFlg4(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessFlg4) { // not required
		return nil
	}

	if err := validate.MinLength("processFlg4", "body", string(*m.ProcessFlg4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processFlg4", "body", string(*m.ProcessFlg4), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProcessFlg5(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessFlg5) { // not required
		return nil
	}

	if err := validate.MinLength("processFlg5", "body", string(*m.ProcessFlg5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("processFlg5", "body", string(*m.ProcessFlg5), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateProdCodeAccType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProdCodeAccType) { // not required
		return nil
	}

	if err := validate.MinLength("prodCodeAccType", "body", string(*m.ProdCodeAccType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prodCodeAccType", "body", string(*m.ProdCodeAccType), 2); err != nil {
		return err
	}

	return nil
}

func (m *AccountBlockCodeUpdateRequest) validateSubStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SubStatus) { // not required
		return nil
	}

	if err := validate.MinLength("subStatus", "body", string(*m.SubStatus), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("subStatus", "body", string(*m.SubStatus), 2); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountBlockCodeUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountBlockCodeUpdateRequest) UnmarshalBinary(b []byte) error {
	var res AccountBlockCodeUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}