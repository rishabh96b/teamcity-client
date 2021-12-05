// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangeLocator Represents a locator string for filtering Change entities.
//
// swagger:model ChangeLocator
type ChangeLocator struct {

	// Build locator.
	Build string `json:"build,omitempty"`

	// Build type locator.
	BuildType string `json:"buildType,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`

	// file
	File string `json:"file,omitempty"`

	// Entity ID.
	ID int32 `json:"id,omitempty"`

	// internal version
	InternalVersion string `json:"internalVersion,omitempty"`

	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`

	// Is pending.
	Pending bool `json:"pending,omitempty"`

	// Project locator.
	Project string `json:"project,omitempty"`

	// Commit SHA since which the changes should be returned.
	SinceChange string `json:"sinceChange,omitempty"`

	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`

	// User locator.
	User string `json:"user,omitempty"`

	// VCS side username.
	Username string `json:"username,omitempty"`

	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`

	// VCS instance locator.
	VcsRootInstance string `json:"vcsRootInstance,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this change locator
func (m *ChangeLocator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change locator based on context it is used
func (m *ChangeLocator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeLocator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeLocator) UnmarshalBinary(b []byte) error {
	var res ChangeLocator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}