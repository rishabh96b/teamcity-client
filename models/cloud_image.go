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

// CloudImage Represents a cloud instance image saved with a profile.
//
// swagger:model cloudImage
type CloudImage struct {

	// agent pool Id
	AgentPoolID int32 `json:"agentPoolId,omitempty"`

	// agent type Id
	AgentTypeID int32 `json:"agentTypeId,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// instances
	Instances *CloudInstances `json:"instances,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// operating system name
	OperatingSystemName string `json:"operatingSystemName,omitempty"`

	// profile
	Profile *CloudProfile `json:"profile,omitempty"`
}

// Validate validates this cloud image
func (m *CloudImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudImage) validateInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	if m.Instances != nil {
		if err := m.Instances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instances")
			}
			return err
		}
	}

	return nil
}

func (m *CloudImage) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud image based on the context it is used
func (m *CloudImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudImage) contextValidateInstances(ctx context.Context, formats strfmt.Registry) error {

	if m.Instances != nil {
		if err := m.Instances.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instances")
			}
			return err
		}
	}

	return nil
}

func (m *CloudImage) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.Profile != nil {
		if err := m.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudImage) UnmarshalBinary(b []byte) error {
	var res CloudImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
