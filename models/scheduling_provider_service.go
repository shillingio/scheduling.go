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

// SchedulingProviderService scheduling provider service
//
// swagger:model schedulingProviderService
type SchedulingProviderService struct {

	// active
	Active bool `json:"active,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// durations
	Durations []int32 `json:"durations"`

	// id
	ID int32 `json:"id,omitempty"`

	// in person
	InPerson bool `json:"in_person,omitempty"`

	// languages
	Languages []string `json:"languages"`

	// license data
	LicenseData map[string]ProtobufAny `json:"license_data,omitempty"`

	// provider
	Provider *SchedulingProvider `json:"provider,omitempty"`

	// provider id
	ProviderID string `json:"provider_id,omitempty"`

	// service
	Service *SchedulingService `json:"service,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// states
	States []string `json:"states"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduling provider service
func (m *SchedulingProviderService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingProviderService) validateLicenseData(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseData) { // not required
		return nil
	}

	for k := range m.LicenseData {

		if err := validate.Required("license_data"+"."+k, "body", m.LicenseData[k]); err != nil {
			return err
		}
		if val, ok := m.LicenseData[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SchedulingProviderService) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulingProviderService) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduling provider service based on the context it is used
func (m *SchedulingProviderService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicenseData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingProviderService) contextValidateLicenseData(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.LicenseData {

		if val, ok := m.LicenseData[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SchedulingProviderService) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.Provider != nil {
		if err := m.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulingProviderService) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {
		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingProviderService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingProviderService) UnmarshalBinary(b []byte) error {
	var res SchedulingProviderService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}