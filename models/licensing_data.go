// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicensingData Represents license state details (available build configurations, agents, etc.).
//
// swagger:model licensingData
type LicensingData struct {

	// agents left
	AgentsLeft int32 `json:"agentsLeft,omitempty" xml:"agentsLeft,attr,omitempty"`

	// build types left
	BuildTypesLeft int32 `json:"buildTypesLeft,omitempty" xml:"buildTypesLeft,attr,omitempty"`

	// license keys
	LicenseKeys *LicenseKeys `json:"licenseKeys,omitempty"`

	// license use exceeded
	LicenseUseExceeded bool `json:"licenseUseExceeded,omitempty" xml:"licenseUseExceeded,attr,omitempty"`

	// max agents
	MaxAgents int32 `json:"maxAgents,omitempty" xml:"maxAgents,attr,omitempty"`

	// max build types
	MaxBuildTypes int32 `json:"maxBuildTypes,omitempty" xml:"maxBuildTypes,attr,omitempty"`

	// server effective release date
	ServerEffectiveReleaseDate string `json:"serverEffectiveReleaseDate,omitempty" xml:"serverEffectiveReleaseDate,attr,omitempty"`

	// server license type
	ServerLicenseType string `json:"serverLicenseType,omitempty" xml:"serverLicenseType,attr,omitempty"`

	// unlimited agents
	UnlimitedAgents bool `json:"unlimitedAgents,omitempty" xml:"unlimitedAgents,attr,omitempty"`

	// unlimited build types
	UnlimitedBuildTypes bool `json:"unlimitedBuildTypes,omitempty" xml:"unlimitedBuildTypes,attr,omitempty"`
}

// Validate validates this licensing data
func (m *LicensingData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicensingData) validateLicenseKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseKeys) { // not required
		return nil
	}

	if m.LicenseKeys != nil {
		if err := m.LicenseKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeys")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this licensing data based on the context it is used
func (m *LicensingData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicenseKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicensingData) contextValidateLicenseKeys(ctx context.Context, formats strfmt.Registry) error {

	if m.LicenseKeys != nil {
		if err := m.LicenseKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicensingData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicensingData) UnmarshalBinary(b []byte) error {
	var res LicensingData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
