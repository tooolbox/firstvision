// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FsErrorDetails fs error details
//
// swagger:model FsErrorDetails
type FsErrorDetails struct {

	// This field contains Error Code assigned to the error Scenario. Example FS000001, FS000002
	ErrorCode string `json:"errorCode,omitempty"`

	// This field contains description of the Error occurred
	Message string `json:"message,omitempty"`

	// This field contains list of backend Errors that has occurred
	OdsMessages []interface{} `json:"odsMessages"`

	// This field contains the HTTP response code for the response
	ResponseCode int32 `json:"responseCode,omitempty"`

	// This field contains result Code assigned for this result
	ResultCode string `json:"resultCode,omitempty"`
}

// Validate validates this fs error details
func (m *FsErrorDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FsErrorDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FsErrorDetails) UnmarshalBinary(b []byte) error {
	var res FsErrorDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}