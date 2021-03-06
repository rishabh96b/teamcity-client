// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricValue Represents a metric value.
//
// swagger:model metricValue
type MetricValue struct {

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// value
	Value float64 `json:"value,omitempty" xml:"value,attr,omitempty"`
}

// Validate validates this metric value
func (m *MetricValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metric value based on context it is used
func (m *MetricValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricValue) UnmarshalBinary(b []byte) error {
	var res MetricValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
