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

// Role role
// swagger:model Role
type Role struct {

	// attributes
	// Required: true
	Attributes RoleAttributes `json:"attributes"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// type
	// Required: true
	Type string `json:"type"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateAttributes(formats strfmt.Registry) error {

	if err := m.Attributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes")
		}
		return err
	}

	return nil
}

var roleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleTypeTypePropEnum = append(roleTypeTypePropEnum, v)
	}
}

const (

	// RoleTypeRole captures enum value "role"
	RoleTypeRole string = "role"
)

// prop value enum
func (m *Role) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, roleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Role) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
