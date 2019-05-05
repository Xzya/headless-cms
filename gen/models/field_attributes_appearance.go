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

// FieldAttributesAppearance field attributes appearance
// swagger:model fieldAttributesAppearance
type FieldAttributesAppearance struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this field attributes appearance
func (m *FieldAttributesAppearance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fieldAttributesAppearanceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["title","plain","markdown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldAttributesAppearanceTypeTypePropEnum = append(fieldAttributesAppearanceTypeTypePropEnum, v)
	}
}

const (

	// FieldAttributesAppearanceTypeTitle captures enum value "title"
	FieldAttributesAppearanceTypeTitle string = "title"

	// FieldAttributesAppearanceTypePlain captures enum value "plain"
	FieldAttributesAppearanceTypePlain string = "plain"

	// FieldAttributesAppearanceTypeMarkdown captures enum value "markdown"
	FieldAttributesAppearanceTypeMarkdown string = "markdown"
)

// prop value enum
func (m *FieldAttributesAppearance) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldAttributesAppearanceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FieldAttributesAppearance) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FieldAttributesAppearance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldAttributesAppearance) UnmarshalBinary(b []byte) error {
	var res FieldAttributesAppearance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}