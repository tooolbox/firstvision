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

// AcctData2ForListCustomerAccounts2 acct data2 for list customer accounts2
//
// swagger:model AcctData2ForListCustomerAccounts2
type AcctData2ForListCustomerAccounts2 struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	Acct string `json:"acct,omitempty"`

	//  Max length = 1, Account Loan Indicator: Indicator to indicate whether the account is a loan account or the Credit account. Valid Values are : L - Loan Account Space - Credit Account
	AcctLoanInd string `json:"acctLoanInd,omitempty"`

	//  Max length = 2, Account level billing cycle: Billing cycle for the account.
	BillngCycl string `json:"billngCycl,omitempty"`

	//  Max length = 1, Account Block Code 1 - Block codes control specific processing options at the account level.
	BlkCd1 string `json:"blkCd1,omitempty"`

	//  Max length = 2, Block code 1 priority: values are 00-99 with 00 as the lowest priority and 99 as the highest priority.
	BlkCd1Prty string `json:"blkCd1Prty,omitempty"`

	//  Max length = 1, Account Block Code 2 - Block codes control specific processing options at the account level.
	BlkCd2 string `json:"blkCd2,omitempty"`

	//  Max length = 2, Block code 2 priority: values are 00-99 with 00 as the lowest priority and 99 as the highest priority.
	BlkCd2Prty string `json:"blkCd2Prty,omitempty"`

	//  Max length = 1, Block code applied on the card.
	BlockCode string `json:"blockCode,omitempty"`

	//  Max length = 1, Type of the cardholder. Values are 0, 1.
	CardholderType string `json:"cardholderType,omitempty"`

	//  Max length = 7, Annual Percentage Rate of the loan plan.
	CelIntApr string `json:"celIntApr,omitempty"`

	//  Max length = 1, Charge Off Status: Code that identifies the charge-off status of this account.  Values are: 0 = Not charged off (Default) 1 = Potential (within 60 days of automatic charge off) 2 = Pending (within 30 days of automatic charge off) 3 = Manual 4 = Blocked 5 = Automatically completed 6 = Manually completed 9 = Account abandoned.
	ChgoffStatus string `json:"chgoffStatus,omitempty"`

	//  Max length = 17, Account spending/credit limit: current credit limit assigned to the account.
	Crlim string `json:"crlim,omitempty"`

	//  Max length = 17, Current balance: Total current balance of all credit plan segments associated with the account.
	CurrBal string `json:"currBal,omitempty"`

	// Format: YYYYMMDD. Transfer effective date.
	DateXfrEff string `json:"dateXfrEff,omitempty"`

	// Format: YYYYMMDD. Date of last payment.
	DteLstPmt string `json:"dteLstPmt,omitempty"`

	//  Max length = 1, Dual Account Indicator: Value received from Input.  Values are: L = Local account F = Foreign account
	DualAcctInd string `json:"dualAcctInd,omitempty"`

	//  Max length = 3, Foreign Organization: Three digit Identification number of the organization.
	ForeignOrg string `json:"foreignOrg,omitempty"`

	//  Max length = 3, Foreign Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	FrgnCurrCd string `json:"frgnCurrCd,omitempty"`

	//  Max length = 1, Foreign Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	FrgnCurrNod string `json:"frgnCurrNod,omitempty"`

	//  Max length = 1, Foreign Percentage NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in percentage fields.
	FrgnPctNod string `json:"frgnPctNod,omitempty"`

	//  Max length = 1, Current First usage flag: Indicates whether card activation needs to occur for the most recently issued card.  Values are:  'Y' - Card activation needs to occur  'N' - No card activation is required or card activation has already occurred   'D' - Card deactivated
	FrstUsageFlg string `json:"frstUsageFlg,omitempty"`

	//  Max length = 1, Status: Code that identifies the status of the account.  Values are: A = Active B = Conversion fraud C = Conversion transfer D = Dormant F = Fraud transfer H = Closed/conversion I = Inactive J = Transfer in, migrated from inactive K = Transfer out, migrated from inactive M = Migrated N = New P = To be purged in next reload Q = Transfer in today R = Transfer out today T = Transfer V = Conversion X = Charge-off/conversion Z = Charge-off 8 = Closed 9 = To be purged after reload.
	IntStatus string `json:"intStatus,omitempty"`

	//  Max length = 1, Loan Status Indicator: Indicator to indicate the status of the loans related to the account. Valid Values are: A -  Loan Active Space - Loan In-active
	LoanStatInd string `json:"loanStatInd,omitempty"`

	//  Max length = 3, Local Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	LocalCurrCd string `json:"localCurrCd,omitempty"`

	//  Max length = 1, Local Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	LocalCurrNod string `json:"localCurrNod,omitempty"`

	//  Max length = 3, Local Organization: Three digit Identification number of the organization.
	LocalOrg string `json:"localOrg,omitempty"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	//  Max length = 3, Life-To-Date Return purchase Percentage. Calculated based on the LTD number of purchase and return numbers.
	LtdRtnPurPct string `json:"ltdRtnPurPct,omitempty"`

	//  Max length = 5, Number of Life-To-Date purchases.
	NbrLtdPrchs string `json:"nbrLtdPrchs,omitempty"`

	//  Max length = 5, Number of Life-To-Date Returns.
	NbrLtdRtrn string `json:"nbrLtdRtrn,omitempty"`

	//  Max length = 5, Number of Year-to-date purchases.
	NbrYtdPrchs string `json:"nbrYtdPrchs,omitempty"`

	//  Max length = 5, Number of Year-To-Date Returns.
	NbrYtdRtrn string `json:"nbrYtdRtrn,omitempty"`

	//  Max length = 17, Current open-to-buy amount of the account.
	OpnToBuy string `json:"opnToBuy,omitempty"`

	//  Max length = 1, Overlimit Flag: Status of the overlimit account.  Values are:  1 thru 7 - Overlimit  1 3 5 7  - Overlimit Fee Assessed  2 3 6 7  - Overlimit Letter Sent  4 5 6 7  - Overlimit Reported
	OvrlmtFlg string `json:"ovrlmtFlg,omitempty"`

	//  Max length = 1, Local Percentage NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in percentage fields.
	PctNod string `json:"pctNod,omitempty"`

	//  Max length = 2, Payment Cycle Due: Code that identifies the level of contractual delinquency. The contractual delinquency (also called the cycle due code), indicates the number of cycles for which payments are due.  Values are: 0 - No amount is due; account is current  1 - Current amount is due; not past due  2 - Amount is past due 1-29 days (or 'X' days)  3 - 30 days  4 - 60 days  5 - 90 days  6 - 120 days  7 - 150 days  8 - 180 days  9 - 210+ days
	PmtCyclDue string `json:"pmtCyclDue,omitempty"`

	//  Max length = 17, Total amount due of the account. This includes the current amount due and any past due amounts.
	PmtTotAmtDue string `json:"pmtTotAmtDue,omitempty"`

	// WRNG CD X
	WrngCdX []*WrngCdX2ForListCustomerAccounts2 `json:"wrngCdX"`

	//  Max length = 19, Transfer Account: Account Number of the transferred account.
	XfrAcct string `json:"xfrAcct,omitempty"`
}

// Validate validates this acct data2 for list customer accounts2
func (m *AcctData2ForListCustomerAccounts2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWrngCdX(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcctData2ForListCustomerAccounts2) validateWrngCdX(formats strfmt.Registry) error {

	if swag.IsZero(m.WrngCdX) { // not required
		return nil
	}

	for i := 0; i < len(m.WrngCdX); i++ {
		if swag.IsZero(m.WrngCdX[i]) { // not required
			continue
		}

		if m.WrngCdX[i] != nil {
			if err := m.WrngCdX[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wrngCdX" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcctData2ForListCustomerAccounts2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcctData2ForListCustomerAccounts2) UnmarshalBinary(b []byte) error {
	var res AcctData2ForListCustomerAccounts2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}