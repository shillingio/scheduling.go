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

// SchedulerScheduleResponse scheduler schedule response
//
// swagger:model schedulerScheduleResponse
type SchedulerScheduleResponse struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// provider
	Provider *SchedulingProvider `json:"provider,omitempty"`

	// range from
	RangeFrom string `json:"range_from,omitempty"`

	// range to
	RangeTo string `json:"range_to,omitempty"`

	// schedule
	Schedule *SchedulingSchedule `json:"schedule,omitempty"`

	// schedules
	Schedules []*SchedulingSchedule `json:"schedules"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this scheduler schedule response
func (m *SchedulerScheduleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerScheduleResponse) validateProvider(formats strfmt.Registry) error {
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

func (m *SchedulerScheduleResponse) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerScheduleResponse) validateSchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedules) { // not required
		return nil
	}

	for i := 0; i < len(m.Schedules); i++ {
		if swag.IsZero(m.Schedules[i]) { // not required
			continue
		}

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scheduler schedule response based on the context it is used
func (m *SchedulerScheduleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerScheduleResponse) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulerScheduleResponse) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulerScheduleResponse) contextValidateSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Schedules); i++ {

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerScheduleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerScheduleResponse) UnmarshalBinary(b []byte) error {
	var res SchedulerScheduleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
