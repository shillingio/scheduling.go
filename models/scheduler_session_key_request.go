// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulerSessionKeyRequest scheduler session key request
//
// swagger:model schedulerSessionKeyRequest
type SchedulerSessionKeyRequest struct {

	// appointment id
	AppointmentID string `json:"appointment_id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// patient
	Patient *SchedulingPatient `json:"patient,omitempty"`

	// patient id
	PatientID string `json:"patient_id,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// provider id
	ProviderID string `json:"provider_id,omitempty"`

	// range from
	RangeFrom string `json:"range_from,omitempty"`

	// range to
	RangeTo string `json:"range_to,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this scheduler session key request
func (m *SchedulerSessionKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerSessionKeyRequest) validatePatient(formats strfmt.Registry) error {
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

// ContextValidate validate this scheduler session key request based on the context it is used
func (m *SchedulerSessionKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerSessionKeyRequest) contextValidatePatient(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SchedulerSessionKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerSessionKeyRequest) UnmarshalBinary(b []byte) error {
	var res SchedulerSessionKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
