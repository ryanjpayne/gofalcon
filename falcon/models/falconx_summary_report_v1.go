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
)

// FalconxSummaryReportV1 falconx summary report v1
//
// swagger:model falconx.SummaryReportV1
type FalconxSummaryReportV1 struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// intel
	Intel []*FalconxIntelSummaryReportV1 `json:"intel"`

	// ioc report broad csv artifact id
	IocReportBroadCsvArtifactID string `json:"ioc_report_broad_csv_artifact_id,omitempty"`

	// ioc report broad json artifact id
	IocReportBroadJSONArtifactID string `json:"ioc_report_broad_json_artifact_id,omitempty"`

	// ioc report broad maec artifact id
	IocReportBroadMaecArtifactID string `json:"ioc_report_broad_maec_artifact_id,omitempty"`

	// ioc report broad stix artifact id
	IocReportBroadStixArtifactID string `json:"ioc_report_broad_stix_artifact_id,omitempty"`

	// ioc report strict csv artifact id
	IocReportStrictCsvArtifactID string `json:"ioc_report_strict_csv_artifact_id,omitempty"`

	// ioc report strict json artifact id
	IocReportStrictJSONArtifactID string `json:"ioc_report_strict_json_artifact_id,omitempty"`

	// ioc report strict maec artifact id
	IocReportStrictMaecArtifactID string `json:"ioc_report_strict_maec_artifact_id,omitempty"`

	// ioc report strict stix artifact id
	IocReportStrictStixArtifactID string `json:"ioc_report_strict_stix_artifact_id,omitempty"`

	// origin
	Origin string `json:"origin,omitempty"`

	// sandbox
	Sandbox []*FalconxSandboxSummaryReportV1 `json:"sandbox"`

	// tags
	Tags []string `json:"tags"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`

	// user tags
	UserTags []string `json:"user_tags"`

	// verdict
	Verdict string `json:"verdict,omitempty"`
}

// Validate validates this falconx summary report v1
func (m *FalconxSummaryReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandbox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSummaryReportV1) validateIntel(formats strfmt.Registry) error {
	if swag.IsZero(m.Intel) { // not required
		return nil
	}

	for i := 0; i < len(m.Intel); i++ {
		if swag.IsZero(m.Intel[i]) { // not required
			continue
		}

		if m.Intel[i] != nil {
			if err := m.Intel[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSummaryReportV1) validateSandbox(formats strfmt.Registry) error {
	if swag.IsZero(m.Sandbox) { // not required
		return nil
	}

	for i := 0; i < len(m.Sandbox); i++ {
		if swag.IsZero(m.Sandbox[i]) { // not required
			continue
		}

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx summary report v1 based on the context it is used
func (m *FalconxSummaryReportV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSandbox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSummaryReportV1) contextValidateIntel(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Intel); i++ {

		if m.Intel[i] != nil {
			if err := m.Intel[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intel" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intel" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSummaryReportV1) contextValidateSandbox(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sandbox); i++ {

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSummaryReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSummaryReportV1) UnmarshalBinary(b []byte) error {
	var res FalconxSummaryReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
