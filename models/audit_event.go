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

// AuditEvent Represents an audit event including a user and affected entities.
//
// swagger:model auditEvent
type AuditEvent struct {

	// action
	Action *AuditAction `json:"action,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// related entities
	RelatedEntities *RelatedEntities `json:"relatedEntities,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this audit event
func (m *AuditEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditEvent) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *AuditEvent) validateRelatedEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedEntities) { // not required
		return nil
	}

	if m.RelatedEntities != nil {
		if err := m.RelatedEntities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relatedEntities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relatedEntities")
			}
			return err
		}
	}

	return nil
}

func (m *AuditEvent) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this audit event based on the context it is used
func (m *AuditEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditEvent) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *AuditEvent) contextValidateRelatedEntities(ctx context.Context, formats strfmt.Registry) error {

	if m.RelatedEntities != nil {
		if err := m.RelatedEntities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relatedEntities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relatedEntities")
			}
			return err
		}
	}

	return nil
}

func (m *AuditEvent) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEvent) UnmarshalBinary(b []byte) error {
	var res AuditEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
