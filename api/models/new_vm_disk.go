// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NewVMDisk new VM disk
// swagger:model NewVMDisk
type NewVMDisk struct {

	// backend storage ID
	BackendStorageID string `json:"backendStorageID,omitempty"`

	// is root
	IsRoot bool `json:"isRoot,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this new VM disk
func (m *NewVMDisk) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewVMDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewVMDisk) UnmarshalBinary(b []byte) error {
	var res NewVMDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
