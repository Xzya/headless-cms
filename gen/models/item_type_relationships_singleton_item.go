// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ItemTypeRelationshipsSingletonItem item type relationships singleton item
// swagger:model itemTypeRelationshipsSingletonItem
type ItemTypeRelationshipsSingletonItem struct {

	// data
	Data *Relationship `json:"data,omitempty"`
}

// Validate validates this item type relationships singleton item
func (m *ItemTypeRelationshipsSingletonItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemTypeRelationshipsSingletonItem) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemTypeRelationshipsSingletonItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemTypeRelationshipsSingletonItem) UnmarshalBinary(b []byte) error {
	var res ItemTypeRelationshipsSingletonItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
