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

// NewNetwork new network
// swagger:model NewNetwork
type NewNetwork struct {

	// physical interface
	// Required: true
	PhysicalInterface *string `json:"physicalInterface"`

	// type
	// Required: true
	// Enum: [linux ovs]
	Type *string `json:"type"`
}

// Validate validates this new network
func (m *NewNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicalInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewNetwork) validatePhysicalInterface(formats strfmt.Registry) error {

	if err := validate.Required("physicalInterface", "body", m.PhysicalInterface); err != nil {
		return err
	}

	return nil
}

var newNetworkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["linux","ovs"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newNetworkTypeTypePropEnum = append(newNetworkTypeTypePropEnum, v)
	}
}

const (

	// NewNetworkTypeLinux captures enum value "linux"
	NewNetworkTypeLinux string = "linux"

	// NewNetworkTypeOvs captures enum value "ovs"
	NewNetworkTypeOvs string = "ovs"
)

// prop value enum
func (m *NewNetwork) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newNetworkTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewNetwork) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewNetwork) UnmarshalBinary(b []byte) error {
	var res NewNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
