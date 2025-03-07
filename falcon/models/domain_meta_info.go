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

// DomainMetaInfo domain meta info
//
// swagger:model domain.MetaInfo
type DomainMetaInfo struct {

	// msa meta info
	// Required: true
	MsaMetaInfo *MsaMetaInfo `json:"MsaMetaInfo"`

	// quota
	Quota *DomainQuota `json:"quota,omitempty"`
}

// Validate validates this domain meta info
func (m *DomainMetaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMsaMetaInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMetaInfo) validateMsaMetaInfo(formats strfmt.Registry) error {

	if err := validate.Required("MsaMetaInfo", "body", m.MsaMetaInfo); err != nil {
		return err
	}

	if m.MsaMetaInfo != nil {
		if err := m.MsaMetaInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MsaMetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MsaMetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainMetaInfo) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain meta info based on the context it is used
func (m *DomainMetaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMsaMetaInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMetaInfo) contextValidateMsaMetaInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MsaMetaInfo != nil {
		if err := m.MsaMetaInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MsaMetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MsaMetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainMetaInfo) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {
		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainMetaInfo) UnmarshalBinary(b []byte) error {
	var res DomainMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
