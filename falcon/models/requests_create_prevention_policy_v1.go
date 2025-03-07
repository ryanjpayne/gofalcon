// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestsCreatePreventionPolicyV1 requests create prevention policy v1
//
// swagger:model requests.CreatePreventionPolicyV1
type RequestsCreatePreventionPolicyV1 struct {

	// If specified the settings of the prevention policy identified by the id will be used
	CloneID string `json:"clone_id,omitempty"`

	// The description to use when creating the policy
	Description string `json:"description,omitempty"`

	// The name to use when creating the policy
	// Required: true
	Name *string `json:"name"`

	// The name of the platform
	// Required: true
	// Enum: [Windows Mac Linux]
	PlatformName *string `json:"platform_name"`

	// The settings to create the policy with
	Settings []*RequestsPreventionSettingV1 `json:"settings"`
}

// Validate validates this requests create prevention policy v1
func (m *RequestsCreatePreventionPolicyV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsCreatePreventionPolicyV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var requestsCreatePreventionPolicyV1TypePlatformNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Windows","Mac","Linux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestsCreatePreventionPolicyV1TypePlatformNamePropEnum = append(requestsCreatePreventionPolicyV1TypePlatformNamePropEnum, v)
	}
}

const (

	// RequestsCreatePreventionPolicyV1PlatformNameWindows captures enum value "Windows"
	RequestsCreatePreventionPolicyV1PlatformNameWindows string = "Windows"

	// RequestsCreatePreventionPolicyV1PlatformNameMac captures enum value "Mac"
	RequestsCreatePreventionPolicyV1PlatformNameMac string = "Mac"

	// RequestsCreatePreventionPolicyV1PlatformNameLinux captures enum value "Linux"
	RequestsCreatePreventionPolicyV1PlatformNameLinux string = "Linux"
)

// prop value enum
func (m *RequestsCreatePreventionPolicyV1) validatePlatformNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, requestsCreatePreventionPolicyV1TypePlatformNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RequestsCreatePreventionPolicyV1) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platform_name", "body", m.PlatformName); err != nil {
		return err
	}

	// value enum
	if err := m.validatePlatformNameEnum("platform_name", "body", *m.PlatformName); err != nil {
		return err
	}

	return nil
}

func (m *RequestsCreatePreventionPolicyV1) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	for i := 0; i < len(m.Settings); i++ {
		if swag.IsZero(m.Settings[i]) { // not required
			continue
		}

		if m.Settings[i] != nil {
			if err := m.Settings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this requests create prevention policy v1 based on the context it is used
func (m *RequestsCreatePreventionPolicyV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsCreatePreventionPolicyV1) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Settings); i++ {

		if m.Settings[i] != nil {
			if err := m.Settings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestsCreatePreventionPolicyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestsCreatePreventionPolicyV1) UnmarshalBinary(b []byte) error {
	var res RequestsCreatePreventionPolicyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
