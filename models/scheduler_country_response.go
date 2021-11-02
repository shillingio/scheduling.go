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

// SchedulerCountryResponse scheduler country response
//
// swagger:model schedulerCountryResponse
type SchedulerCountryResponse struct {

	// countries
	Countries []*SchedulingCountry `json:"countries"`

	// country
	Country *SchedulingCountry `json:"country,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this scheduler country response
func (m *SchedulerCountryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerCountryResponse) validateCountries(formats strfmt.Registry) error {
	if swag.IsZero(m.Countries) { // not required
		return nil
	}

	for i := 0; i < len(m.Countries); i++ {
		if swag.IsZero(m.Countries[i]) { // not required
			continue
		}

		if m.Countries[i] != nil {
			if err := m.Countries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SchedulerCountryResponse) validateCountry(formats strfmt.Registry) error {
	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if m.Country != nil {
		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduler country response based on the context it is used
func (m *SchedulerCountryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCountries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerCountryResponse) contextValidateCountries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Countries); i++ {

		if m.Countries[i] != nil {
			if err := m.Countries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SchedulerCountryResponse) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if m.Country != nil {
		if err := m.Country.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerCountryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerCountryResponse) UnmarshalBinary(b []byte) error {
	var res SchedulerCountryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}