// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLocator Represents a locator string for filtering AuditEvent entities.
//
// swagger:model AuditLocator
type AuditLocator struct {

	// Use `$help` to get the full list of supported actions.
	Action string `json:"action,omitempty"`

	// Related project locator.
	AffectedProject string `json:"affectedProject,omitempty"`

	// Related build type or template locator.
	BuildType string `json:"buildType,omitempty"`

	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`

	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`

	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`

	// system action
	SystemAction bool `json:"systemAction,omitempty"`

	// Locator of user who caused the audit event.
	User string `json:"user,omitempty"`
}

// Validate validates this audit locator
func (m *AuditLocator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit locator based on context it is used
func (m *AuditLocator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLocator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLocator) UnmarshalBinary(b []byte) error {
	var res AuditLocator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
