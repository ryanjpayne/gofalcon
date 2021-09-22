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

// ResponsesPreventionPolicyV1 A fully formed prevention policy
//
// swagger:model responses.PreventionPolicyV1
type ResponsesPreventionPolicyV1 struct {

	// The customer id associated with the policy
	// Required: true
	Cid *string `json:"cid"`

	// The email of the user which created the policy
	// Required: true
	CreatedBy *string `json:"created_by"`

	// The time at which the policy was created
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// The description of a policy. Use this field to provide a high level summary of what this policy enforces
	// Required: true
	Description *string `json:"description"`

	// If a policy is enabled it will be used during the course of policy evaluation
	// Required: true
	Enabled *bool `json:"enabled"`

	// The groups that are currently attached to the policy
	// Required: true
	Groups []*ResponsesHostGroupV1 `json:"groups"`

	// The unique id of the policy
	// Required: true
	ID *string `json:"id"`

	// The IOA rule groups that are currently attached to the policy
	// Required: true
	IoaRuleGroups []*ResponsesIOARuleGroupV1 `json:"ioa_rule_groups"`

	// The email of the user which last modified the policy
	// Required: true
	ModifiedBy *string `json:"modified_by"`

	// The time at which the policy was last modified
	// Required: true
	// Format: date-time
	ModifiedTimestamp *strfmt.DateTime `json:"modified_timestamp"`

	// The human readable name of the policy
	// Required: true
	Name *string `json:"name"`

	// The name of the platform
	// Required: true
	// Enum: [Windows Mac Linux]
	PlatformName *string `json:"platform_name"`

	// A category of settings in a prevention policy
	// Required: true
	PreventionSettings []*ResponsesPreventionCategoryV1 `json:"prevention_settings"`
}

// Validate validates this responses prevention policy v1
func (m *ResponsesPreventionPolicyV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoaRuleGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreventionSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesPreventionPolicyV1) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateIoaRuleGroups(formats strfmt.Registry) error {

	if err := validate.Required("ioa_rule_groups", "body", m.IoaRuleGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.IoaRuleGroups); i++ {
		if swag.IsZero(m.IoaRuleGroups[i]) { // not required
			continue
		}

		if m.IoaRuleGroups[i] != nil {
			if err := m.IoaRuleGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioa_rule_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("modified_by", "body", m.ModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateModifiedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("modified_timestamp", "body", m.ModifiedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_timestamp", "body", "date-time", m.ModifiedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var responsesPreventionPolicyV1TypePlatformNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Windows","Mac","Linux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesPreventionPolicyV1TypePlatformNamePropEnum = append(responsesPreventionPolicyV1TypePlatformNamePropEnum, v)
	}
}

const (

	// ResponsesPreventionPolicyV1PlatformNameWindows captures enum value "Windows"
	ResponsesPreventionPolicyV1PlatformNameWindows string = "Windows"

	// ResponsesPreventionPolicyV1PlatformNameMac captures enum value "Mac"
	ResponsesPreventionPolicyV1PlatformNameMac string = "Mac"

	// ResponsesPreventionPolicyV1PlatformNameLinux captures enum value "Linux"
	ResponsesPreventionPolicyV1PlatformNameLinux string = "Linux"
)

// prop value enum
func (m *ResponsesPreventionPolicyV1) validatePlatformNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesPreventionPolicyV1TypePlatformNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesPreventionPolicyV1) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platform_name", "body", m.PlatformName); err != nil {
		return err
	}

	// value enum
	if err := m.validatePlatformNameEnum("platform_name", "body", *m.PlatformName); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) validatePreventionSettings(formats strfmt.Registry) error {

	if err := validate.Required("prevention_settings", "body", m.PreventionSettings); err != nil {
		return err
	}

	for i := 0; i < len(m.PreventionSettings); i++ {
		if swag.IsZero(m.PreventionSettings[i]) { // not required
			continue
		}

		if m.PreventionSettings[i] != nil {
			if err := m.PreventionSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prevention_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this responses prevention policy v1 based on the context it is used
func (m *ResponsesPreventionPolicyV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoaRuleGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreventionSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesPreventionPolicyV1) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) contextValidateIoaRuleGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IoaRuleGroups); i++ {

		if m.IoaRuleGroups[i] != nil {
			if err := m.IoaRuleGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioa_rule_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesPreventionPolicyV1) contextValidatePreventionSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreventionSettings); i++ {

		if m.PreventionSettings[i] != nil {
			if err := m.PreventionSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prevention_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesPreventionPolicyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesPreventionPolicyV1) UnmarshalBinary(b []byte) error {
	var res ResponsesPreventionPolicyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
