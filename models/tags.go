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

// Tags Represents a list of Tag entities.
//
// swagger:model tags
type Tags struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// tag
	Tag []*Tag `json:"tag"`
}

// Validate validates this tags
func (m *Tags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tags) validateTag(formats strfmt.Registry) error {
	if swag.IsZero(m.Tag) { // not required
		return nil
	}

	for i := 0; i < len(m.Tag); i++ {
		if swag.IsZero(m.Tag[i]) { // not required
			continue
		}

		if m.Tag[i] != nil {
			if err := m.Tag[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tag" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tags based on the context it is used
func (m *Tags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tags) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tag); i++ {

		if m.Tag[i] != nil {
			if err := m.Tag[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tag" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tags) UnmarshalBinary(b []byte) error {
	var res Tags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
