// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StateField Represents a project state field (as of now, limited to the read-only state of project).
//
// swagger:model StateField
type StateField struct {

	// inherited
	Inherited bool `json:"inherited,omitempty" xml:"inherited,attr,omitempty"`

	// value
	Value bool `json:"value,omitempty" xml:"value,attr,omitempty"`
}

// Validate validates this state field
func (m *StateField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this state field based on context it is used
func (m *StateField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateField) UnmarshalBinary(b []byte) error {
	var res StateField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
