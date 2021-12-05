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

// User Represents a user.
//
// swagger:model user
type User struct {

	// email
	Email string `json:"email,omitempty" xml:"email,attr,omitempty"`

	// groups
	Groups *Groups `json:"groups,omitempty"`

	// has password
	HasPassword bool `json:"hasPassword,omitempty" xml:"hasPassword,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID int64 `json:"id,omitempty" xml:"id,attr,omitempty"`

	// last login
	LastLogin string `json:"lastLogin,omitempty" xml:"lastLogin,attr,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// password
	Password string `json:"password,omitempty" xml:"password,attr,omitempty"`

	// properties
	Properties *Properties `json:"properties,omitempty"`

	// realm
	Realm string `json:"realm,omitempty" xml:"realm,attr,omitempty"`

	// roles
	Roles *Roles `json:"roles,omitempty"`

	// username
	Username string `json:"username,omitempty" xml:"username,attr,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	if m.Groups != nil {
		if err := m.Groups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groups")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if m.Roles != nil {
		if err := m.Roles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.Groups != nil {
		if err := m.Groups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groups")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {
		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	if m.Roles != nil {
		if err := m.Roles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}