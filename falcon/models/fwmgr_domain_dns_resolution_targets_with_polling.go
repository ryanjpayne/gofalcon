// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FwmgrDomainDNSResolutionTargetsWithPolling fwmgr domain DNS resolution targets with polling
//
// swagger:model fwmgr.domain.DNSResolutionTargetsWithPolling
type FwmgrDomainDNSResolutionTargetsWithPolling struct {

	// polling interval
	// Required: true
	PollingInterval *int32 `json:"polling_interval"`

	// targets
	Targets []*FwmgrDomainDNSTarget `json:"targets"`
}

// Validate validates this fwmgr domain DNS resolution targets with polling
func (m *FwmgrDomainDNSResolutionTargetsWithPolling) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePollingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainDNSResolutionTargetsWithPolling) validatePollingInterval(formats strfmt.Registry) error {

	if err := validate.Required("polling_interval", "body", m.PollingInterval); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrDomainDNSResolutionTargetsWithPolling) validateTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fwmgr domain DNS resolution targets with polling based on the context it is used
func (m *FwmgrDomainDNSResolutionTargetsWithPolling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainDNSResolutionTargetsWithPolling) contextValidateTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Targets); i++ {

		if m.Targets[i] != nil {
			if err := m.Targets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrDomainDNSResolutionTargetsWithPolling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrDomainDNSResolutionTargetsWithPolling) UnmarshalBinary(b []byte) error {
	var res FwmgrDomainDNSResolutionTargetsWithPolling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
