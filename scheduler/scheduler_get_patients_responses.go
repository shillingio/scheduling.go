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

// SchedulerGetPatientsReader is a Reader for the SchedulerGetPatients structure.
type SchedulerGetPatientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerGetPatientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerGetPatientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerGetPatientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerGetPatientsOK creates a SchedulerGetPatientsOK with default headers values
func NewSchedulerGetPatientsOK() *SchedulerGetPatientsOK {
	return &SchedulerGetPatientsOK{}
}

/* SchedulerGetPatientsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerGetPatientsOK struct {
	Payload *models.SchedulerPatientResponse
}

func (o *SchedulerGetPatientsOK) Error() string {
	return fmt.Sprintf("[GET /v1/patients][%d] schedulerGetPatientsOK  %+v", 200, o.Payload)
}
func (o *SchedulerGetPatientsOK) GetPayload() *models.SchedulerPatientResponse {
	return o.Payload
}

func (o *SchedulerGetPatientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerPatientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerGetPatientsDefault creates a SchedulerGetPatientsDefault with default headers values
func NewSchedulerGetPatientsDefault(code int) *SchedulerGetPatientsDefault {
	return &SchedulerGetPatientsDefault{
		_statusCode: code,
	}
}

/* SchedulerGetPatientsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerGetPatientsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler get patients default response
func (o *SchedulerGetPatientsDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerGetPatientsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/patients][%d] Scheduler_GetPatients default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerGetPatientsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerGetPatientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
