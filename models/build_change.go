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

// BuildChange Represents links to the next or previous build.
//
// swagger:model buildChange
type BuildChange struct {

	// next build
	NextBuild *Build `json:"nextBuild,omitempty"`

	// prev build
	PrevBuild *Build `json:"prevBuild,omitempty"`
}

// Validate validates this build change
func (m *BuildChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildChange) validateNextBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.NextBuild) { // not required
		return nil
	}

	if m.NextBuild != nil {
		if err := m.NextBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextBuild")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextBuild")
			}
			return err
		}
	}

	return nil
}

func (m *BuildChange) validatePrevBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.PrevBuild) { // not required
		return nil
	}

	if m.PrevBuild != nil {
		if err := m.PrevBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prevBuild")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prevBuild")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this build change based on the context it is used
func (m *BuildChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrevBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildChange) contextValidateNextBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.NextBuild != nil {
		if err := m.NextBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextBuild")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextBuild")
			}
			return err
		}
	}

	return nil
}

func (m *BuildChange) contextValidatePrevBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.PrevBuild != nil {
		if err := m.PrevBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prevBuild")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prevBuild")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildChange) UnmarshalBinary(b []byte) error {
	var res BuildChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}