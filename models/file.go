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

// File Represents a file.
//
// swagger:model file
type File struct {

	// children
	Children *Files `json:"children,omitempty"`

	// content
	Content *Href `json:"content,omitempty"`

	// full name
	FullName string `json:"fullName,omitempty" xml:"fullName,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// modification time
	ModificationTime string `json:"modificationTime,omitempty" xml:"modificationTime,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// parent
	Parent *File `json:"parent,omitempty"`

	// size
	Size int64 `json:"size,omitempty" xml:"size,attr,omitempty"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) validateChildren(formats strfmt.Registry) error {
	if swag.IsZero(m.Children) { // not required
		return nil
	}

	if m.Children != nil {
		if err := m.Children.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("children")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("children")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file based on the context it is used
func (m *File) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildren(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) contextValidateChildren(ctx context.Context, formats strfmt.Registry) error {

	if m.Children != nil {
		if err := m.Children.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("children")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("children")
			}
			return err
		}
	}

	return nil
}

func (m *File) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *File) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *File) UnmarshalBinary(b []byte) error {
	var res File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
