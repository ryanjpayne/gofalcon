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

// FwmgrAPINetworkLocationModifyRequestV1 fwmgr api network location modify request v1
//
// swagger:model fwmgr.api.NetworkLocationModifyRequestV1
type FwmgrAPINetworkLocationModifyRequestV1 struct {

	// connection types
	// Required: true
	ConnectionTypes *FwmgrDomainConnectionType `json:"connection_types"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created on
	CreatedOn string `json:"created_on,omitempty"`

	// default gateways
	// Required: true
	DefaultGateways []string `json:"default_gateways"`

	// description
	// Required: true
	Description *string `json:"description"`

	// dhcp servers
	// Required: true
	DhcpServers []string `json:"dhcp_servers"`

	// dns resolution targets
	// Required: true
	DNSResolutionTargets *FwmgrDomainDNSResolutionTargets `json:"dns_resolution_targets"`

	// dns servers
	// Required: true
	DNSServers []string `json:"dns_servers"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// host addresses
	// Required: true
	HostAddresses []string `json:"host_addresses"`

	// https reachable hosts
	// Required: true
	HTTPSReachableHosts *FwmgrDomainHTTPSHosts `json:"https_reachable_hosts"`

	// icmp request targets
	// Required: true
	IcmpRequestTargets *FwmgrDomainICMPTargets `json:"icmp_request_targets"`

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
}

// Validate validates this fwmgr api network location modify request v1
func (m *FwmgrAPINetworkLocationModifyRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSResolutionTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPSReachableHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpRequestTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateConnectionTypes(formats strfmt.Registry) error {

	if err := validate.Required("connection_types", "body", m.ConnectionTypes); err != nil {
		return err
	}

	if m.ConnectionTypes != nil {
		if err := m.ConnectionTypes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_types")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection_types")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateDefaultGateways(formats strfmt.Registry) error {

	if err := validate.Required("default_gateways", "body", m.DefaultGateways); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateDhcpServers(formats strfmt.Registry) error {

	if err := validate.Required("dhcp_servers", "body", m.DhcpServers); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateDNSResolutionTargets(formats strfmt.Registry) error {

	if err := validate.Required("dns_resolution_targets", "body", m.DNSResolutionTargets); err != nil {
		return err
	}

	if m.DNSResolutionTargets != nil {
		if err := m.DNSResolutionTargets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_resolution_targets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_resolution_targets")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateDNSServers(formats strfmt.Registry) error {

	if err := validate.Required("dns_servers", "body", m.DNSServers); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateHostAddresses(formats strfmt.Registry) error {

	if err := validate.Required("host_addresses", "body", m.HostAddresses); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateHTTPSReachableHosts(formats strfmt.Registry) error {

	if err := validate.Required("https_reachable_hosts", "body", m.HTTPSReachableHosts); err != nil {
		return err
	}

	if m.HTTPSReachableHosts != nil {
		if err := m.HTTPSReachableHosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("https_reachable_hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("https_reachable_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateIcmpRequestTargets(formats strfmt.Registry) error {

	if err := validate.Required("icmp_request_targets", "body", m.IcmpRequestTargets); err != nil {
		return err
	}

	if m.IcmpRequestTargets != nil {
		if err := m.IcmpRequestTargets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icmp_request_targets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icmp_request_targets")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fwmgr api network location modify request v1 based on the context it is used
func (m *FwmgrAPINetworkLocationModifyRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSResolutionTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPSReachableHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIcmpRequestTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) contextValidateConnectionTypes(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionTypes != nil {
		if err := m.ConnectionTypes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_types")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection_types")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) contextValidateDNSResolutionTargets(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSResolutionTargets != nil {
		if err := m.DNSResolutionTargets.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_resolution_targets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_resolution_targets")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) contextValidateHTTPSReachableHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPSReachableHosts != nil {
		if err := m.HTTPSReachableHosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("https_reachable_hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("https_reachable_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyRequestV1) contextValidateIcmpRequestTargets(ctx context.Context, formats strfmt.Registry) error {

	if m.IcmpRequestTargets != nil {
		if err := m.IcmpRequestTargets.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icmp_request_targets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icmp_request_targets")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationModifyRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationModifyRequestV1) UnmarshalBinary(b []byte) error {
	var res FwmgrAPINetworkLocationModifyRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
