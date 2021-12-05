// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Steps Represents a list of Step entities.
//
// swagger:model steps
type Steps struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// step
	Step []*Step `json:"step"`
}

// Validate validates this steps
func (m *Steps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Steps) validateStep(formats strfmt.Registry) error {
	if swag.IsZero(m.Step) { // not required
		return nil
	}

	for i := 0; i < len(m.Step); i++ {
		if swag.IsZero(m.Step[i]) { // not required
			continue
		}

		if m.Step[i] != nil {
			if err := m.Step[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("step" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("step" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this steps based on the context it is used
func (m *Steps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Steps) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Step); i++ {

		if m.Step[i] != nil {
			if err := m.Step[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("step" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("step" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Steps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Steps) UnmarshalBinary(b []byte) error {
	var res Steps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
