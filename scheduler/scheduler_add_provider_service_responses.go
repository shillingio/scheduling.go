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

// SchedulerAddProviderServiceReader is a Reader for the SchedulerAddProviderService structure.
type SchedulerAddProviderServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerAddProviderServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerAddProviderServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerAddProviderServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerAddProviderServiceOK creates a SchedulerAddProviderServiceOK with default headers values
func NewSchedulerAddProviderServiceOK() *SchedulerAddProviderServiceOK {
	return &SchedulerAddProviderServiceOK{}
}

/* SchedulerAddProviderServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerAddProviderServiceOK struct {
	Payload *models.SchedulerProviderServiceResponse
}

func (o *SchedulerAddProviderServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/providers/{provider.id}/services][%d] schedulerAddProviderServiceOK  %+v", 200, o.Payload)
}
func (o *SchedulerAddProviderServiceOK) GetPayload() *models.SchedulerProviderServiceResponse {
	return o.Payload
}

func (o *SchedulerAddProviderServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerProviderServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerAddProviderServiceDefault creates a SchedulerAddProviderServiceDefault with default headers values
func NewSchedulerAddProviderServiceDefault(code int) *SchedulerAddProviderServiceDefault {
	return &SchedulerAddProviderServiceDefault{
		_statusCode: code,
	}
}

/* SchedulerAddProviderServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerAddProviderServiceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler add provider service default response
func (o *SchedulerAddProviderServiceDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerAddProviderServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/providers/{provider.id}/services][%d] Scheduler_AddProviderService default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerAddProviderServiceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerAddProviderServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
