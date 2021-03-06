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

// TestRunMetadata Represents a list of TypedValue entities.
//
// swagger:model testRunMetadata
type TestRunMetadata struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// typed values
	TypedValues []*TypedValue `json:"typedValues"`
}

// Validate validates this test run metadata
func (m *TestRunMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypedValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRunMetadata) validateTypedValues(formats strfmt.Registry) error {
	if swag.IsZero(m.TypedValues) { // not required
		return nil
	}

	for i := 0; i < len(m.TypedValues); i++ {
		if swag.IsZero(m.TypedValues[i]) { // not required
			continue
		}

		if m.TypedValues[i] != nil {
			if err := m.TypedValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("typedValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("typedValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test run metadata based on the context it is used
func (m *TestRunMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTypedValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRunMetadata) contextValidateTypedValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypedValues); i++ {

		if m.TypedValues[i] != nil {
			if err := m.TypedValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("typedValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("typedValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestRunMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRunMetadata) UnmarshalBinary(b []byte) error {
	var res TestRunMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
