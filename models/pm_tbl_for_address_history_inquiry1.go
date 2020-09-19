// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PmTblForAddressHistoryInquiry1 pm tbl for address history inquiry1
//
// swagger:model PmTblForAddressHistoryInquiry1
type PmTblForAddressHistoryInquiry1 struct {

	//  Max length = 40, Address Line 1 of the customer as defined on the Customer Name/Address record.
	PmAddr1 string `json:"pmAddr1,omitempty"`

	//  Max length = 40, Address Line 2 of the customer as defined on the Customer Name/Address record.
	PmAddr2 string `json:"pmAddr2,omitempty"`

	//  Max length = 40, Mailing address line 3 for the customer Group variable of below 2 fields  All the below two fields need to be passed to update ADDR-3.
	PmAddr3 string `json:"pmAddr3,omitempty"`

	//  Max length = 40, Mailing address line 4 for the customer
	PmAddr4 string `json:"pmAddr4,omitempty"`

	//  Max length = 30, City: City portion of the mailing address.
	PmCity string `json:"pmCity,omitempty"`

	//  Max length = 3, Country portion of the mailing address. The country code must be a valid value as defined in the security subsystem.
	PmCountryCd string `json:"pmCountryCd,omitempty"`

	// Format: YYYYMMDD. Date.
	PmDate string `json:"pmDate,omitempty"`

	//  Max length = 60, Email address of the customer.
	PmEmailAddress string `json:"pmEmailAddress,omitempty"`

	//  Max length = 1, Email Flag: code that indicates whether you have permission from the cardholder to Email. Values are:  0 - Do not send  1 - May be sent.  2 - Prefer to be sent.
	PmEmailFlag string `json:"pmEmailFlag,omitempty"`

	//  Max length = 20, Fax number of the customer
	PmFaxPhone string `json:"pmFaxPhone,omitempty"`

	//  Max length = 1, Fax Flag: Code that specifies whether the customer grants permission to use this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	PmFaxPhoneFlag string `json:"pmFaxPhoneFlag,omitempty"`

	//  Max length = 20, Home phone number of the customer
	PmHomePhone string `json:"pmHomePhone,omitempty"`

	//  Max length = 1, Home Phone Flag: Code that specifies whether the customer grants permission to call this number.  Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method.
	PmHomePhoneFlag string `json:"pmHomePhoneFlag,omitempty"`

	//  Max length = 40, House name of the customer.
	PmHouseName string `json:"pmHouseName,omitempty"`

	//  Max length = 20, House Number of the customer.
	PmHouseNbr string `json:"pmHouseNbr,omitempty"`

	//  Max length = 3, Language Indicator: User-defined code that indicates the language spoken by the customer.
	PmLanguageIndic string `json:"pmLanguageIndic,omitempty"`

	//  Max length = 20, Mobile number of the accountholder.
	PmMobilePhone string `json:"pmMobilePhone,omitempty"`

	//  Max length = 1, Mobile Phone Flag: Code that specifies whether the customer grants permission to call this number.  Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	PmMobilePhoneFlag string `json:"pmMobilePhoneFlag,omitempty"`

	//  Max length = 40, Name line 1 of the customer as defined on the Customer Name/Address record.
	PmNameLine1 string `json:"pmNameLine1,omitempty"`

	//  Max length = 40, Name line 2 of the customer as defined on the Customer Name/Address record.
	PmNameLine2 string `json:"pmNameLine2,omitempty"`

	//  Max length = 40, Name line 3 of the customer as defined on the Customer Name/Address record.
	PmNameLine3 string `json:"pmNameLine3,omitempty"`

	//  Max length = 1, Name Type Indicator 1: Indicates the type of name. Values are: 0 = Personal name (Default) 1 = Business name 2 = Store name 3 = Generic name.
	PmNameTypeInd1 string `json:"pmNameTypeInd1,omitempty"`

	//  Max length = 1, Indicates the type of name. Valid values are: 0 = Personal name (Default) 1 = Business name 2 = Store name 3 = Generic name.
	PmNameTypeInd2 string `json:"pmNameTypeInd2,omitempty"`

	//  Max length = 1, Indicates the type of name. Valid values are: 0 = Personal name (Default) 1 = Business name 2 = Store name 3 = Generic name.
	PmNameTypeInd3 string `json:"pmNameTypeInd3,omitempty"`

	//  Max length = 10, Postal Code: Postal code portion of the mailing address.
	PmPstlCd string `json:"pmPstlCd,omitempty"`

	//  Max length = 68, Reserved amounts for future use.
	PmResvAmount string `json:"pmResvAmount,omitempty"`

	//  Max length = 8, Reserved codes for future use.
	PmResvCode string `json:"pmResvCode,omitempty"`

	//  Max length = 32, Reserved dates for future use.
	PmResvDate string `json:"pmResvDate,omitempty"`

	//  Max length = 3, State: State or province portion of the mailing address.
	PmState string `json:"pmState,omitempty"`

	//  Max length = 1, Version number of the Statement Message record to use for the customer . The values are 0-9. The default is 0.
	PmStmtMsgIndic string `json:"pmStmtMsgIndic,omitempty"`

	//  Max length = 20, Employee Phone: Work telephone number of the customer.
	PmWorkPhone string `json:"pmWorkPhone,omitempty"`

	//  Max length = 6, Work phone extension number of the customer.
	PmWorkPhoneExtn string `json:"pmWorkPhoneExtn,omitempty"`

	//  Max length = 1, workPhone Flag: Code that specifies whether the customer grants permission to  call this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method.
	PmWorkPhoneFlag string `json:"pmWorkPhoneFlag,omitempty"`
}

// Validate validates this pm tbl for address history inquiry1
func (m *PmTblForAddressHistoryInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PmTblForAddressHistoryInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PmTblForAddressHistoryInquiry1) UnmarshalBinary(b []byte) error {
	var res PmTblForAddressHistoryInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}