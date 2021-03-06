// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/shillingio/scheduling.go/models"
)

// SchedulerDeleteAppointmentReader is a Reader for the SchedulerDeleteAppointment structure.
type SchedulerDeleteAppointmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerDeleteAppointmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerDeleteAppointmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerDeleteAppointmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerDeleteAppointmentOK creates a SchedulerDeleteAppointmentOK with default headers values
func NewSchedulerDeleteAppointmentOK() *SchedulerDeleteAppointmentOK {
	return &SchedulerDeleteAppointmentOK{}
}

/* SchedulerDeleteAppointmentOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerDeleteAppointmentOK struct {
	Payload interface{}
}

func (o *SchedulerDeleteAppointmentOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/appointments/{id}][%d] schedulerDeleteAppointmentOK  %+v", 200, o.Payload)
}
func (o *SchedulerDeleteAppointmentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SchedulerDeleteAppointmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerDeleteAppointmentDefault creates a SchedulerDeleteAppointmentDefault with default headers values
func NewSchedulerDeleteAppointmentDefault(code int) *SchedulerDeleteAppointmentDefault {
	return &SchedulerDeleteAppointmentDefault{
		_statusCode: code,
	}
}

/* SchedulerDeleteAppointmentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerDeleteAppointmentDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler delete appointment default response
func (o *SchedulerDeleteAppointmentDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerDeleteAppointmentDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/appointments/{id}][%d] Scheduler_DeleteAppointment default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerDeleteAppointmentDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerDeleteAppointmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
