// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AboutUserMsgVpn about user msg vpn
// swagger:model AboutUserMsgVpn
type AboutUserMsgVpn struct {

	// Message VPN access level of the Current User. The allowed values and their meaning are:
	//
	// <pre>
	// "none" - No access allowed.
	// "read-only" - Read only.
	// "read-write" - Read and Write.
	// </pre>
	//
	// Enum: [none read-only read-write]
	AccessLevel string `json:"accessLevel,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
}

// Validate validates this about user msg vpn
func (m *AboutUserMsgVpn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aboutUserMsgVpnTypeAccessLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","read-only","read-write"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aboutUserMsgVpnTypeAccessLevelPropEnum = append(aboutUserMsgVpnTypeAccessLevelPropEnum, v)
	}
}

const (

	// AboutUserMsgVpnAccessLevelNone captures enum value "none"
	AboutUserMsgVpnAccessLevelNone string = "none"

	// AboutUserMsgVpnAccessLevelReadOnly captures enum value "read-only"
	AboutUserMsgVpnAccessLevelReadOnly string = "read-only"

	// AboutUserMsgVpnAccessLevelReadWrite captures enum value "read-write"
	AboutUserMsgVpnAccessLevelReadWrite string = "read-write"
)

// prop value enum
func (m *AboutUserMsgVpn) validateAccessLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, aboutUserMsgVpnTypeAccessLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AboutUserMsgVpn) validateAccessLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessLevelEnum("accessLevel", "body", m.AccessLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AboutUserMsgVpn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AboutUserMsgVpn) UnmarshalBinary(b []byte) error {
	var res AboutUserMsgVpn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
