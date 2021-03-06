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

// Features Represents a list of Feature entities.
//
// swagger:model features
type Features struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// feature
	Feature []*Feature `json:"feature"`
}

// Validate validates this features
func (m *Features) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Features) validateFeature(formats strfmt.Registry) error {
	if swag.IsZero(m.Feature) { // not required
		return nil
	}

	for i := 0; i < len(m.Feature); i++ {
		if swag.IsZero(m.Feature[i]) { // not required
			continue
		}

		if m.Feature[i] != nil {
			if err := m.Feature[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feature" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feature" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this features based on the context it is used
func (m *Features) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Features) contextValidateFeature(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Feature); i++ {

		if m.Feature[i] != nil {
			if err := m.Feature[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feature" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feature" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Features) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Features) UnmarshalBinary(b []byte) error {
	var res Features
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
