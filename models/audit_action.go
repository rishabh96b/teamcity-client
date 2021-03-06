// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditAction Represents an audit action.
//
// swagger:model auditAction
type AuditAction struct {

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`
}

// Validate validates this audit action
func (m *AuditAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit action based on context it is used
func (m *AuditAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditAction) UnmarshalBinary(b []byte) error {
	var res AuditAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
