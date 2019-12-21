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

// NewImage new image
// swagger:model NewImage
type NewImage struct {

	// build context
	// Format: byte
	BuildContext strfmt.Base64 `json:"buildContext,omitempty"`

	// build context image type
	// Enum: [standard osv qemu]
	BuildContextImageType string `json:"buildContextImageType,omitempty"`

	// build context type
	// Enum: [promethium docker tar raw qcow2 capstan]
	BuildContextType string `json:"buildContextType,omitempty"`

	// clone from
	CloneFrom string `json:"cloneFrom,omitempty"`

	// operation
	// Required: true
	// Enum: [clone build]
	Operation *string `json:"operation"`
}

// Validate validates this new image
func (m *NewImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildContextImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildContextType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewImage) validateBuildContext(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildContext) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

var newImageTypeBuildContextImageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["standard","osv","qemu"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newImageTypeBuildContextImageTypePropEnum = append(newImageTypeBuildContextImageTypePropEnum, v)
	}
}

const (

	// NewImageBuildContextImageTypeStandard captures enum value "standard"
	NewImageBuildContextImageTypeStandard string = "standard"

	// NewImageBuildContextImageTypeOsv captures enum value "osv"
	NewImageBuildContextImageTypeOsv string = "osv"

	// NewImageBuildContextImageTypeQemu captures enum value "qemu"
	NewImageBuildContextImageTypeQemu string = "qemu"
)

// prop value enum
func (m *NewImage) validateBuildContextImageTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newImageTypeBuildContextImageTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewImage) validateBuildContextImageType(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildContextImageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBuildContextImageTypeEnum("buildContextImageType", "body", m.BuildContextImageType); err != nil {
		return err
	}

	return nil
}

var newImageTypeBuildContextTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["promethium","docker","tar","raw","qcow2","capstan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newImageTypeBuildContextTypePropEnum = append(newImageTypeBuildContextTypePropEnum, v)
	}
}

const (

	// NewImageBuildContextTypePromethium captures enum value "promethium"
	NewImageBuildContextTypePromethium string = "promethium"

	// NewImageBuildContextTypeDocker captures enum value "docker"
	NewImageBuildContextTypeDocker string = "docker"

	// NewImageBuildContextTypeTar captures enum value "tar"
	NewImageBuildContextTypeTar string = "tar"

	// NewImageBuildContextTypeRaw captures enum value "raw"
	NewImageBuildContextTypeRaw string = "raw"

	// NewImageBuildContextTypeQcow2 captures enum value "qcow2"
	NewImageBuildContextTypeQcow2 string = "qcow2"

	// NewImageBuildContextTypeCapstan captures enum value "capstan"
	NewImageBuildContextTypeCapstan string = "capstan"
)

// prop value enum
func (m *NewImage) validateBuildContextTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newImageTypeBuildContextTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewImage) validateBuildContextType(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildContextType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBuildContextTypeEnum("buildContextType", "body", m.BuildContextType); err != nil {
		return err
	}

	return nil
}

var newImageTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clone","build"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newImageTypeOperationPropEnum = append(newImageTypeOperationPropEnum, v)
	}
}

const (

	// NewImageOperationClone captures enum value "clone"
	NewImageOperationClone string = "clone"

	// NewImageOperationBuild captures enum value "build"
	NewImageOperationBuild string = "build"
)

// prop value enum
func (m *NewImage) validateOperationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newImageTypeOperationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewImage) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", *m.Operation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewImage) UnmarshalBinary(b []byte) error {
	var res NewImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}