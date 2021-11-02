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

// SchedulingAppointment scheduling appointment
// Example: {}
//
// swagger:model schedulingAppointment
type SchedulingAppointment struct {

	// active
	Active bool `json:"active,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// The duration of the appointment in minutes
	Duration int32 `json:"duration,omitempty"`

	// The end time of the appointment (RFC3339)
	EndAt string `json:"end_at,omitempty"`

	// UUID of the appointment
	ID string `json:"id,omitempty"`

	// If the appointment is in-person and not online
	InPerson bool `json:"in_person,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// Patient the appointment is with
	Patient *SchedulingPatient `json:"patient,omitempty"`

	// Provider the appointment is with
	Provider *SchedulingProvider `json:"provider,omitempty"`

	// ref id
	RefID string `json:"ref_id,omitempty"`

	// schedule block id
	ScheduleBlockID string `json:"schedule_block_id,omitempty"`

	// schedule id
	ScheduleID string `json:"schedule_id,omitempty"`

	// service
	Service *SchedulingService `json:"service,omitempty"`

	// The service(s) that is/are scheduled with the appointment
	Services []*SchedulingService `json:"services"`

	// The start time of the appointment (RFC3339)
	StartAt string `json:"start_at,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduling appointment
func (m *SchedulingAppointment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingAppointment) validatePatient(formats strfmt.Registry) error {
	if swag.IsZero(m.Patient) { // not required
		return nil
	}

	if m.Patient != nil {
		if err := m.Patient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patient")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulingAppointment) validateProvider(formats strfmt.Registry) error {
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

func (m *SchedulingAppointment) validateService(formats strfmt.Registry) error {
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

func (m *SchedulingAppointment) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scheduling appointment based on the context it is used
func (m *SchedulingAppointment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingAppointment) contextValidatePatient(ctx context.Context, formats strfmt.Registry) error {

	if m.Patient != nil {
		if err := m.Patient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patient")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulingAppointment) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulingAppointment) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulingAppointment) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {
			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingAppointment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingAppointment) UnmarshalBinary(b []byte) error {
	var res SchedulingAppointment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
