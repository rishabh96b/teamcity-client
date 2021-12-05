// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagLocator Represents a locator string for filtering Tag entities.
//
// swagger:model TagLocator
type TagLocator struct {

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// private
	Private string `json:"private,omitempty"`
}

// Validate validates this tag locator
func (m *TagLocator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag locator based on context it is used
func (m *TagLocator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagLocator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagLocator) UnmarshalBinary(b []byte) error {
	var res TagLocator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
