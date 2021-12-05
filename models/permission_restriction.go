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

// PermissionRestriction Represents permission restrictions of an authentication token.
//
// swagger:model permissionRestriction
type PermissionRestriction struct {

	// is global scope
	IsGlobalScope bool `json:"isGlobalScope,omitempty" xml:"isGlobalScope,attr,omitempty"`

	// permission
	Permission *Permission `json:"permission,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`
}

// Validate validates this permission restriction
func (m *PermissionRestriction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionRestriction) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if m.Permission != nil {
		if err := m.Permission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permission")
			}
			return err
		}
	}

	return nil
}

func (m *PermissionRestriction) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this permission restriction based on the context it is used
func (m *PermissionRestriction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionRestriction) contextValidatePermission(ctx context.Context, formats strfmt.Registry) error {

	if m.Permission != nil {
		if err := m.Permission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permission")
			}
			return err
		}
	}

	return nil
}

func (m *PermissionRestriction) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PermissionRestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionRestriction) UnmarshalBinary(b []byte) error {
	var res PermissionRestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}