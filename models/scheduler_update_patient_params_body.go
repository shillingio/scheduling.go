// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulerUpdatePatientParamsBody scheduler update patient params body
// Example: {}
//
// swagger:model schedulerUpdatePatientParamsBody
type SchedulerUpdatePatientParamsBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// given name
	GivenName string `json:"given_name,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// ref id
	RefID string `json:"ref_id,omitempty"`

	// suffix
	Suffix string `json:"suffix,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduler update patient params body
func (m *SchedulerUpdatePatientParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scheduler update patient params body based on context it is used
func (m *SchedulerUpdatePatientParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerUpdatePatientParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerUpdatePatientParamsBody) UnmarshalBinary(b []byte) error {
	var res SchedulerUpdatePatientParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}