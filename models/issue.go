// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Issue Represents an issue related to the specific change.
//
// swagger:model Issue
type Issue struct {

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// url
	URL string `json:"url,omitempty" xml:"url,attr,omitempty"`
}

// Validate validates this issue
func (m *Issue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this issue based on context it is used
func (m *Issue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Issue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Issue) UnmarshalBinary(b []byte) error {
	var res Issue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
