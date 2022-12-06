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

// FwmgrAPINetworkLocationSummaryV1 fwmgr api network location summary v1
//
// swagger:model fwmgr.api.NetworkLocationSummaryV1
type FwmgrAPINetworkLocationSummaryV1 struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created on
	CreatedOn string `json:"created_on,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Required: true
	ID *string `json:"id"`

	// modified by
	ModifiedBy string `json:"modified_by,omitempty"`

	// modified on
	ModifiedOn string `json:"modified_on,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// rule count
	// Required: true
	RuleCount *int32 `json:"rule_count"`
}

// Validate validates this fwmgr api network location summary v1
func (m *FwmgrAPINetworkLocationSummaryV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationSummaryV1) validateRuleCount(formats strfmt.Registry) error {

	if err := validate.Required("rule_count", "body", m.RuleCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr api network location summary v1 based on context it is used
func (m *FwmgrAPINetworkLocationSummaryV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationSummaryV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationSummaryV1) UnmarshalBinary(b []byte) error {
	var res FwmgrAPINetworkLocationSummaryV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
