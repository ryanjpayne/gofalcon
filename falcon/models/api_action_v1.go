// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIActionV1 api action v1
//
// swagger:model api.ActionV1
type APIActionV1 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// platforms by type
	// Required: true
	PlatformsByType map[string]APIActionV1PlatformsByType `json:"platforms_by_type"`

	// severities
	// Required: true
	Severities []string `json:"severities"`
}

// Validate validates this api action v1
func (m *APIActionV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformsByType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIActionV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIActionV1) validatePlatformsByType(formats strfmt.Registry) error {

	if err := validate.Required("platforms_by_type", "body", m.PlatformsByType); err != nil {
		return err
	}

	for k := range m.PlatformsByType {

		if err := validate.Required("platforms_by_type"+"."+k, "body", m.PlatformsByType[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *APIActionV1) validateSeverities(formats strfmt.Registry) error {

	if err := validate.Required("severities", "body", m.Severities); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this api action v1 based on context it is used
func (m *APIActionV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIActionV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIActionV1) UnmarshalBinary(b []byte) error {
	var res APIActionV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
