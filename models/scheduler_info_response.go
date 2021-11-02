// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulerInfoResponse scheduler info response
//
// swagger:model schedulerInfoResponse
type SchedulerInfoResponse struct {

	// build
	Build string `json:"build,omitempty"`

	// built at
	BuiltAt string `json:"built_at,omitempty"`

	// git hash
	GitHash string `json:"git_hash,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this scheduler info response
func (m *SchedulerInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scheduler info response based on context it is used
func (m *SchedulerInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerInfoResponse) UnmarshalBinary(b []byte) error {
	var res SchedulerInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
