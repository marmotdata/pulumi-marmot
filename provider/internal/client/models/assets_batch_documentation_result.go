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

// AssetsBatchDocumentationResult assets batch documentation result
//
// swagger:model assets.BatchDocumentationResult
type AssetsBatchDocumentationResult struct {

	// documentation
	Documentation *AssetdocsDocumentation `json:"documentation,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this assets batch documentation result
func (m *AssetsBatchDocumentationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetsBatchDocumentationResult) validateDocumentation(formats strfmt.Registry) error {
	if swag.IsZero(m.Documentation) { // not required
		return nil
	}

	if m.Documentation != nil {
		if err := m.Documentation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("documentation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this assets batch documentation result based on the context it is used
func (m *AssetsBatchDocumentationResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetsBatchDocumentationResult) contextValidateDocumentation(ctx context.Context, formats strfmt.Registry) error {

	if m.Documentation != nil {

		if swag.IsZero(m.Documentation) { // not required
			return nil
		}

		if err := m.Documentation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("documentation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetsBatchDocumentationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetsBatchDocumentationResult) UnmarshalBinary(b []byte) error {
	var res AssetsBatchDocumentationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
