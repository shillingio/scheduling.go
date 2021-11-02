// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/shillingio/scheduling.go/models"
)

// SchedulerUpdatePatientReader is a Reader for the SchedulerUpdatePatient structure.
type SchedulerUpdatePatientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerUpdatePatientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerUpdatePatientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerUpdatePatientDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerUpdatePatientOK creates a SchedulerUpdatePatientOK with default headers values
func NewSchedulerUpdatePatientOK() *SchedulerUpdatePatientOK {
	return &SchedulerUpdatePatientOK{}
}

/* SchedulerUpdatePatientOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerUpdatePatientOK struct {
	Payload *models.SchedulerPatientResponse
}

func (o *SchedulerUpdatePatientOK) Error() string {
	return fmt.Sprintf("[PUT /v1/patients/{id}][%d] schedulerUpdatePatientOK  %+v", 200, o.Payload)
}
func (o *SchedulerUpdatePatientOK) GetPayload() *models.SchedulerPatientResponse {
	return o.Payload
}

func (o *SchedulerUpdatePatientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerPatientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerUpdatePatientDefault creates a SchedulerUpdatePatientDefault with default headers values
func NewSchedulerUpdatePatientDefault(code int) *SchedulerUpdatePatientDefault {
	return &SchedulerUpdatePatientDefault{
		_statusCode: code,
	}
}

/* SchedulerUpdatePatientDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerUpdatePatientDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler update patient default response
func (o *SchedulerUpdatePatientDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerUpdatePatientDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/patients/{id}][%d] Scheduler_UpdatePatient default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerUpdatePatientDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerUpdatePatientDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SchedulerUpdatePatientBody scheduler update patient body
// Example: {}
swagger:model SchedulerUpdatePatientBody
*/
type SchedulerUpdatePatientBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// data
	Data *models.ProtobufAny `json:"data,omitempty"`

	// given name
	GivenName string `json:"given_name,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// ref id
	RefID string `json:"ref_id,omitempty"`

	// suffix
	Suffix string `json:"suffix,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduler update patient body
func (o *SchedulerUpdatePatientBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerUpdatePatientBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduler update patient body based on the context it is used
func (o *SchedulerUpdatePatientBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerUpdatePatientBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchedulerUpdatePatientBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchedulerUpdatePatientBody) UnmarshalBinary(b []byte) error {
	var res SchedulerUpdatePatientBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
