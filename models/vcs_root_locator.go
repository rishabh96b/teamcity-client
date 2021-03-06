// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VcsRootLocator Represents a locator string for filtering VcsRoot entities.
//
// swagger:model VcsRootLocator
type VcsRootLocator struct {

	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`

	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`

	// Entity ID.
	ID int32 `json:"id,omitempty"`

	// internal Id
	InternalID string `json:"internalId,omitempty"`

	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`

	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Project (direct parent) locator.
	Project string `json:"project,omitempty"`

	// Supported matchType values:
	// - generic: exists/not-exists/equals/does-not-equal/starts-with/contains/does-not-contain/ends-with/any;
	// - regular expressions: matches/does-not-match;
	// - numeric: more-than/no-more-than/less-than/no-less-than;
	// - version-specific: ver-more-than/ver-no-more-than/ver-less-than/ver-no-less-than.
	Property string `json:"property,omitempty"`

	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`

	// Type of VCS (e.g. jetbrains.git).
	Type string `json:"type,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this vcs root locator
func (m *VcsRootLocator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vcs root locator based on context it is used
func (m *VcsRootLocator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VcsRootLocator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsRootLocator) UnmarshalBinary(b []byte) error {
	var res VcsRootLocator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
