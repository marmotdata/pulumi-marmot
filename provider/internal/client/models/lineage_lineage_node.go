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

// LineageLineageNode lineage lineage node
//
// swagger:model lineage.LineageNode
type LineageLineageNode struct {

	// asset
	Asset *AssetAsset `json:"asset,omitempty"`

	// depth
	Depth int64 `json:"depth,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lineage lineage node
func (m *LineageLineageNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageLineageNode) validateAsset(formats strfmt.Registry) error {
	if swag.IsZero(m.Asset) { // not required
		return nil
	}

	if m.Asset != nil {
		if err := m.Asset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asset")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lineage lineage node based on the context it is used
func (m *LineageLineageNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageLineageNode) contextValidateAsset(ctx context.Context, formats strfmt.Registry) error {

	if m.Asset != nil {

		if swag.IsZero(m.Asset) { // not required
			return nil
		}

		if err := m.Asset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LineageLineageNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LineageLineageNode) UnmarshalBinary(b []byte) error {
	var res LineageLineageNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
