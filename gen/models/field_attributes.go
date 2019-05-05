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

// FieldAttributes field attributes
// swagger:model fieldAttributes
type FieldAttributes struct {

	// api key
	// Required: true
	// Max Length: 32
	// Min Length: 1
	APIKey string `json:"apiKey"`

	// appearance
	// Required: true
	Appearance FieldAttributesAppearance `json:"appearance"`

	// field type
	// Required: true
	FieldType string `json:"fieldType"`

	// hint
	Hint string `json:"hint,omitempty"`

	// label
	// Required: true
	Label string `json:"label"`

	// localized
	Localized bool `json:"localized,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// validators
	// Required: true
	Validators FieldAttributesValidators `json:"validators"`
}

// Validate validates this field attributes
func (m *FieldAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAppearance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFieldType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValidators(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FieldAttributes) validateAPIKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("apiKey", "body", string(m.APIKey)); err != nil {
		return err
	}

	if err := validate.MinLength("apiKey", "body", string(m.APIKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("apiKey", "body", string(m.APIKey), 32); err != nil {
		return err
	}

	return nil
}

func (m *FieldAttributes) validateAppearance(formats strfmt.Registry) error {

	if err := m.Appearance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appearance")
		}
		return err
	}

	return nil
}

var fieldAttributesTypeFieldTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldAttributesTypeFieldTypePropEnum = append(fieldAttributesTypeFieldTypePropEnum, v)
	}
}

const (

	// FieldAttributesFieldTypeString captures enum value "string"
	FieldAttributesFieldTypeString string = "string"
)

// prop value enum
func (m *FieldAttributes) validateFieldTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldAttributesTypeFieldTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FieldAttributes) validateFieldType(formats strfmt.Registry) error {

	if err := validate.RequiredString("fieldType", "body", string(m.FieldType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateFieldTypeEnum("fieldType", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

func (m *FieldAttributes) validateLabel(formats strfmt.Registry) error {

	if err := validate.RequiredString("label", "body", string(m.Label)); err != nil {
		return err
	}

	return nil
}

func (m *FieldAttributes) validateValidators(formats strfmt.Registry) error {

	if err := m.Validators.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validators")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FieldAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldAttributes) UnmarshalBinary(b []byte) error {
	var res FieldAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}