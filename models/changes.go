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

// Changes Represents a paginated list of Change entities.
//
// swagger:model changes
type Changes struct {

	// change
	Change []*Change `json:"change"`

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref,attr,omitempty"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref,attr,omitempty"`
}

// Validate validates this changes
func (m *Changes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Changes) validateChange(formats strfmt.Registry) error {
	if swag.IsZero(m.Change) { // not required
		return nil
	}

	for i := 0; i < len(m.Change); i++ {
		if swag.IsZero(m.Change[i]) { // not required
			continue
		}

		if m.Change[i] != nil {
			if err := m.Change[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("change" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("change" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this changes based on the context it is used
func (m *Changes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Changes) contextValidateChange(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Change); i++ {

		if m.Change[i] != nil {
			if err := m.Change[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("change" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("change" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Changes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Changes) UnmarshalBinary(b []byte) error {
	var res Changes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
