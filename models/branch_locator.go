// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BranchLocator Represents a locator string for filtering Branch entities.
//
// swagger:model BranchLocator
type BranchLocator struct {

	// Is feature branch.
	Branched bool `json:"branched,omitempty"`

	// Build locator.
	Build string `json:"build,omitempty"`

	// Build type locator.
	BuildType string `json:"buildType,omitempty"`

	// Is default branch.
	Default bool `json:"default,omitempty"`

	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this branch locator
func (m *BranchLocator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this branch locator based on context it is used
func (m *BranchLocator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BranchLocator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchLocator) UnmarshalBinary(b []byte) error {
	var res BranchLocator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
