// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ItemTypeRelationships item type relationships
// swagger:model ItemTypeRelationships
type ItemTypeRelationships struct {

	// fields
	Fields *ItemTypeRelationshipsFields `json:"fields,omitempty"`

	// ordering field
	OrderingField *ItemTypeRelationshipsOrderingField `json:"orderingField,omitempty"`

	// singleton item
	SingletonItem *ItemTypeRelationshipsSingletonItem `json:"singletonItem,omitempty"`
}

// Validate validates this item type relationships
func (m *ItemTypeRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderingField(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSingletonItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemTypeRelationships) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if m.Fields != nil {

		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			}
			return err
		}

	}

	return nil
}

func (m *ItemTypeRelationships) validateOrderingField(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderingField) { // not required
		return nil
	}

	if m.OrderingField != nil {

		if err := m.OrderingField.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderingField")
			}
			return err
		}

	}

	return nil
}

func (m *ItemTypeRelationships) validateSingletonItem(formats strfmt.Registry) error {

	if swag.IsZero(m.SingletonItem) { // not required
		return nil
	}

	if m.SingletonItem != nil {

		if err := m.SingletonItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("singletonItem")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemTypeRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemTypeRelationships) UnmarshalBinary(b []byte) error {
	var res ItemTypeRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}