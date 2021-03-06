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

// SchedulerSessionResponse scheduler session response
//
// swagger:model schedulerSessionResponse
type SchedulerSessionResponse struct {

	// Current Appointment Info
	Appointment *SchedulingAppointment `json:"appointment,omitempty"`

	// Expiration time of the session.  (Once expired creating a new session is required)
	ExpiresAt string `json:"expires_at,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// Additional Metadata Provided for session
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Current Patient Info (typically only ID and Reference IDs)
	Patient *SchedulingPatient `json:"patient,omitempty"`

	// Current Provider Info
	Provider *SchedulingProvider `json:"provider,omitempty"`

	// Redirect Search Parameters
	Redirect *SchedulerSessionRedirect `json:"redirect,omitempty"`

	// Current Search Parameters
	Search *SchedulerSessionSearch `json:"search,omitempty"`

	// Current Service Info
	Service *SchedulingService `json:"service,omitempty"`

	// Session Settings (Used for Shilling UIs)
	Settings map[string]interface{} `json:"settings,omitempty"`

	// Shortened URL for simplicity and commonality in presenting to end-users
	ShortLink string `json:"short_link,omitempty"`

	// Current Appointment Slot Info
	Slot *SchedulingAppointmentSlot `json:"slot,omitempty"`
}

// Validate validates this scheduler session response
func (m *SchedulerSessionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerSessionResponse) validateAppointment(formats strfmt.Registry) error {
	if swag.IsZero(m.Appointment) { // not required
		return nil
	}

	if m.Appointment != nil {
		if err := m.Appointment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointment")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) validatePatient(formats strfmt.Registry) error {
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

func (m *SchedulerSessionResponse) validateProvider(formats strfmt.Registry) error {
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

func (m *SchedulerSessionResponse) validateRedirect(formats strfmt.Registry) error {
	if swag.IsZero(m.Redirect) { // not required
		return nil
	}

	if m.Redirect != nil {
		if err := m.Redirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redirect")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) validateSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.Search) { // not required
		return nil
	}

	if m.Search != nil {
		if err := m.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) validateService(formats strfmt.Registry) error {
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

func (m *SchedulerSessionResponse) validateSlot(formats strfmt.Registry) error {
	if swag.IsZero(m.Slot) { // not required
		return nil
	}

	if m.Slot != nil {
		if err := m.Slot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduler session response based on the context it is used
func (m *SchedulerSessionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedirect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerSessionResponse) contextValidateAppointment(ctx context.Context, formats strfmt.Registry) error {

	if m.Appointment != nil {
		if err := m.Appointment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointment")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) contextValidatePatient(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulerSessionResponse) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulerSessionResponse) contextValidateRedirect(ctx context.Context, formats strfmt.Registry) error {

	if m.Redirect != nil {
		if err := m.Redirect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redirect")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) contextValidateSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Search != nil {
		if err := m.Search.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerSessionResponse) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulerSessionResponse) contextValidateSlot(ctx context.Context, formats strfmt.Registry) error {

	if m.Slot != nil {
		if err := m.Slot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerSessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerSessionResponse) UnmarshalBinary(b []byte) error {
	var res SchedulerSessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
