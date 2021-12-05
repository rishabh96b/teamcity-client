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

// TestOccurrences Represents a paginated list of TestOccurrence entities.
//
// swagger:model testOccurrences
type TestOccurrences struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// failed
	Failed int32 `json:"failed,omitempty" xml:"failed,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// ignored
	Ignored int32 `json:"ignored,omitempty" xml:"ignored,attr,omitempty"`

	// muted
	Muted int32 `json:"muted,omitempty" xml:"muted,attr,omitempty"`

	// new failed
	NewFailed int32 `json:"newFailed,omitempty" xml:"newFailed,attr,omitempty"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref,attr,omitempty"`

	// passed
	Passed int32 `json:"passed,omitempty" xml:"passed,attr,omitempty"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref,attr,omitempty"`

	// test counters
	TestCounters *TestCounters `json:"testCounters,omitempty"`

	// test occurrence
	TestOccurrence []*TestOccurrence `json:"testOccurrence"`
}

// Validate validates this test occurrences
func (m *TestOccurrences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTestCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestOccurrence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestOccurrences) validateTestCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.TestCounters) { // not required
		return nil
	}

	if m.TestCounters != nil {
		if err := m.TestCounters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testCounters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("testCounters")
			}
			return err
		}
	}

	return nil
}

func (m *TestOccurrences) validateTestOccurrence(formats strfmt.Registry) error {
	if swag.IsZero(m.TestOccurrence) { // not required
		return nil
	}

	for i := 0; i < len(m.TestOccurrence); i++ {
		if swag.IsZero(m.TestOccurrence[i]) { // not required
			continue
		}

		if m.TestOccurrence[i] != nil {
			if err := m.TestOccurrence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testOccurrence" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("testOccurrence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test occurrences based on the context it is used
func (m *TestOccurrences) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTestCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTestOccurrence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestOccurrences) contextValidateTestCounters(ctx context.Context, formats strfmt.Registry) error {

	if m.TestCounters != nil {
		if err := m.TestCounters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testCounters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("testCounters")
			}
			return err
		}
	}

	return nil
}

func (m *TestOccurrences) contextValidateTestOccurrence(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TestOccurrence); i++ {

		if m.TestOccurrence[i] != nil {
			if err := m.TestOccurrence[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testOccurrence" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("testOccurrence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestOccurrences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestOccurrences) UnmarshalBinary(b []byte) error {
	var res TestOccurrences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
