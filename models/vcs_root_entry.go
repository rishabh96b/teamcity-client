// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VcsRootEntry Represents a VCS root assigned to this build configuration.
//
// swagger:model vcs-root-entry
type VcsRootEntry struct {

	// checkout rules
	CheckoutRules string `json:"checkout-rules,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// inherited
	Inherited bool `json:"inherited,omitempty" xml:"inherited,attr,omitempty"`

	// vcs root
	VcsRoot *VcsRoot `json:"vcs-root,omitempty"`
}

// Validate validates this vcs root entry
func (m *VcsRootEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcsRoot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsRootEntry) validateVcsRoot(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsRoot) { // not required
		return nil
	}

	if m.VcsRoot != nil {
		if err := m.VcsRoot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs-root")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs-root")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vcs root entry based on the context it is used
func (m *VcsRootEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsRoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsRootEntry) contextValidateVcsRoot(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsRoot != nil {
		if err := m.VcsRoot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs-root")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs-root")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsRootEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsRootEntry) UnmarshalBinary(b []byte) error {
	var res VcsRootEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
