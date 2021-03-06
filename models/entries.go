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

// Entries Represents a list of Entry entities.
//
// swagger:model entries
type Entries struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// entry
	Entry []*Entry `json:"entry"`
}

// Validate validates this entries
func (m *Entries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entries) validateEntry(formats strfmt.Registry) error {
	if swag.IsZero(m.Entry) { // not required
		return nil
	}

	for i := 0; i < len(m.Entry); i++ {
		if swag.IsZero(m.Entry[i]) { // not required
			continue
		}

		if m.Entry[i] != nil {
			if err := m.Entry[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entry" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this entries based on the context it is used
func (m *Entries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entries) contextValidateEntry(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entry); i++ {

		if m.Entry[i] != nil {
			if err := m.Entry[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entry" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Entries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entries) UnmarshalBinary(b []byte) error {
	var res Entries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
