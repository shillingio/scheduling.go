// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulingStats scheduling stats
//
// swagger:model schedulingStats
type SchedulingStats struct {

	// total active providers
	TotalActiveProviders int32 `json:"total_active_providers,omitempty"`

	// total appts future
	TotalApptsFuture int32 `json:"total_appts_future,omitempty"`

	// total appts last 30 days
	TotalApptsLast30Days int32 `json:"total_appts_last_30_days,omitempty"`

	// total appts last 60 days
	TotalApptsLast60Days int32 `json:"total_appts_last_60_days,omitempty"`

	// total appts last 90 days
	TotalApptsLast90Days int32 `json:"total_appts_last_90_days,omitempty"`

	// total appts last day
	TotalApptsLastDay int32 `json:"total_appts_last_day,omitempty"`

	// total appts last week
	TotalApptsLastWeek int32 `json:"total_appts_last_week,omitempty"`

	// total appts past
	TotalApptsPast int32 `json:"total_appts_past,omitempty"`

	// total inactive providers
	TotalInactiveProviders int32 `json:"total_inactive_providers,omitempty"`

	// total patients
	TotalPatients int32 `json:"total_patients,omitempty"`

	// total providers
	TotalProviders int32 `json:"total_providers,omitempty"`

	// total schedules
	TotalSchedules int32 `json:"total_schedules,omitempty"`

	// total services
	TotalServices int32 `json:"total_services,omitempty"`
}

// Validate validates this scheduling stats
func (m *SchedulingStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scheduling stats based on context it is used
func (m *SchedulingStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingStats) UnmarshalBinary(b []byte) error {
	var res SchedulingStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
