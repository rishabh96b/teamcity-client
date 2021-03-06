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

// ProblemOccurrence Represents an instance of a failed test in the specific build.
//
// swagger:model problemOccurrence
type ProblemOccurrence struct {

	// additional data
	AdditionalData string `json:"additionalData,omitempty"`

	// build
	Build *Build `json:"build,omitempty"`

	// currently investigated
	CurrentlyInvestigated bool `json:"currentlyInvestigated,omitempty" xml:"currentlyInvestigated,attr,omitempty"`

	// currently muted
	CurrentlyMuted bool `json:"currentlyMuted,omitempty" xml:"currentlyMuted,attr,omitempty"`

	// details
	Details string `json:"details,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// identity
	Identity string `json:"identity,omitempty" xml:"identity,attr,omitempty"`

	// log anchor
	LogAnchor string `json:"logAnchor,omitempty" xml:"logAnchor,attr,omitempty"`

	// mute
	Mute *Mute `json:"mute,omitempty"`

	// muted
	Muted bool `json:"muted,omitempty" xml:"muted,attr,omitempty"`

	// new failure
	NewFailure bool `json:"newFailure,omitempty" xml:"newFailure,attr,omitempty"`

	// problem
	Problem *Problem `json:"problem,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`
}

// Validate validates this problem occurrence
func (m *ProblemOccurrence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProblem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProblemOccurrence) validateBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemOccurrence) validateMute(formats strfmt.Registry) error {
	if swag.IsZero(m.Mute) { // not required
		return nil
	}

	if m.Mute != nil {
		if err := m.Mute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mute")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemOccurrence) validateProblem(formats strfmt.Registry) error {
	if swag.IsZero(m.Problem) { // not required
		return nil
	}

	if m.Problem != nil {
		if err := m.Problem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("problem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("problem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this problem occurrence based on the context it is used
func (m *ProblemOccurrence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProblem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProblemOccurrence) contextValidateBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.Build != nil {
		if err := m.Build.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemOccurrence) contextValidateMute(ctx context.Context, formats strfmt.Registry) error {

	if m.Mute != nil {
		if err := m.Mute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mute")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemOccurrence) contextValidateProblem(ctx context.Context, formats strfmt.Registry) error {

	if m.Problem != nil {
		if err := m.Problem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("problem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("problem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProblemOccurrence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemOccurrence) UnmarshalBinary(b []byte) error {
	var res ProblemOccurrence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
