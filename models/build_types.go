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

// BuildTypes Represents a paginated list of BuildType entities.
//
// swagger:model buildTypes
type BuildTypes struct {

	// build type
	BuildType []*BuildType `json:"buildType"`

	// count
	Count int32 `json:"count,omitempty" xml:"count,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref,attr,omitempty"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref,attr,omitempty"`
}

// Validate validates this build types
func (m *BuildTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildTypes) validateBuildType(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildType) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildType); i++ {
		if swag.IsZero(m.BuildType[i]) { // not required
			continue
		}

		if m.BuildType[i] != nil {
			if err := m.BuildType[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buildType" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buildType" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this build types based on the context it is used
func (m *BuildTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildTypes) contextValidateBuildType(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BuildType); i++ {

		if m.BuildType[i] != nil {
			if err := m.BuildType[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buildType" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buildType" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildTypes) UnmarshalBinary(b []byte) error {
	var res BuildTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}