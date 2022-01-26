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

// SchedulerActivateScheduleParamsBody scheduler activate schedule params body
//
// swagger:model schedulerActivateScheduleParamsBody
type SchedulerActivateScheduleParamsBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// active from
	ActiveFrom string `json:"active_from,omitempty"`

	// active to
	ActiveTo string `json:"active_to,omitempty"`

	// blocks
	Blocks []*SchedulingScheduleBlock `json:"blocks"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// include blocks
	IncludeBlocks bool `json:"include_blocks,omitempty"`

	// max per
	MaxPer int32 `json:"max_per,omitempty"`

	// max per type
	MaxPerType string `json:"max_per_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provider
	Provider *SchedulingProvider `json:"provider,omitempty"`

	// timezone
	Timezone *SchedulingTimezone `json:"timezone,omitempty"`

	// tz
	Tz string `json:"tz,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduler activate schedule params body
func (m *SchedulerActivateScheduleParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerActivateScheduleParamsBody) validateBlocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Blocks) { // not required
		return nil
	}

	for i := 0; i < len(m.Blocks); i++ {
		if swag.IsZero(m.Blocks[i]) { // not required
			continue
		}

		if m.Blocks[i] != nil {
			if err := m.Blocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SchedulerActivateScheduleParamsBody) validateProvider(formats strfmt.Registry) error {
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

func (m *SchedulerActivateScheduleParamsBody) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	if m.Timezone != nil {
		if err := m.Timezone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduler activate schedule params body based on the context it is used
func (m *SchedulerActivateScheduleParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerActivateScheduleParamsBody) contextValidateBlocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Blocks); i++ {

		if m.Blocks[i] != nil {
			if err := m.Blocks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SchedulerActivateScheduleParamsBody) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchedulerActivateScheduleParamsBody) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if m.Timezone != nil {
		if err := m.Timezone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerActivateScheduleParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerActivateScheduleParamsBody) UnmarshalBinary(b []byte) error {
	var res SchedulerActivateScheduleParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}