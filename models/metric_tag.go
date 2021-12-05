// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricTag Represents a metric tag.
//
// swagger:model metricTag
type MetricTag struct {

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// value
	Value string `json:"value,omitempty" xml:"value,attr,omitempty"`
}

// Validate validates this metric tag
func (m *MetricTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metric tag based on context it is used
func (m *MetricTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricTag) UnmarshalBinary(b []byte) error {
	var res MetricTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
