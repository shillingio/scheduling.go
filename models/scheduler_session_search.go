// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulerSessionSearch scheduler session search
//
// swagger:model schedulerSessionSearch
type SchedulerSessionSearch struct {

	// duration
	Duration int32 `json:"duration,omitempty"`

	// range from
	RangeFrom string `json:"range_from,omitempty"`

	// range to
	RangeTo string `json:"range_to,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// zipcode
	Zipcode string `json:"zipcode,omitempty"`
}

// Validate validates this scheduler session search
func (m *SchedulerSessionSearch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scheduler session search based on context it is used
func (m *SchedulerSessionSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerSessionSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerSessionSearch) UnmarshalBinary(b []byte) error {
	var res SchedulerSessionSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
