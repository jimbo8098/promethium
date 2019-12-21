// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkIP4Config network IP 4 config
// swagger:model NetworkIP4Config
type NetworkIP4Config struct {

	// address
	// Format: ipv4
	Address strfmt.IPv4 `json:"address,omitempty"`

	// dhcp
	Dhcp bool `json:"dhcp,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// gateway
	// Format: ipv4
	Gateway strfmt.IPv4 `json:"gateway,omitempty"`

	// vlan
	Vlan int32 `json:"vlan,omitempty"`
}

// Validate validates this network IP 4 config
func (m *NetworkIP4Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkIP4Config) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.FormatOf("address", "body", "ipv4", m.Address.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIP4Config) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if err := validate.FormatOf("gateway", "body", "ipv4", m.Gateway.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIP4Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIP4Config) UnmarshalBinary(b []byte) error {
	var res NetworkIP4Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
