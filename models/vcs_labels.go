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

// VcsLabels Represents a list of VcsLabel entities.
//
// swagger:model vcsLabels
type VcsLabels struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// vcs label
	VcsLabel []*VcsLabel `json:"vcsLabel"`
}

// Validate validates this vcs labels
func (m *VcsLabels) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcsLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsLabels) validateVcsLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsLabel) { // not required
		return nil
	}

	for i := 0; i < len(m.VcsLabel); i++ {
		if swag.IsZero(m.VcsLabel[i]) { // not required
			continue
		}

		if m.VcsLabel[i] != nil {
			if err := m.VcsLabel[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcsLabel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcsLabel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vcs labels based on the context it is used
func (m *VcsLabels) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsLabels) contextValidateVcsLabel(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VcsLabel); i++ {

		if m.VcsLabel[i] != nil {
			if err := m.VcsLabel[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcsLabel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcsLabel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsLabels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsLabels) UnmarshalBinary(b []byte) error {
	var res VcsLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
