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

// SchedulerInfoReader is a Reader for the SchedulerInfo structure.
type SchedulerInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerInfoOK creates a SchedulerInfoOK with default headers values
func NewSchedulerInfoOK() *SchedulerInfoOK {
	return &SchedulerInfoOK{}
}

/* SchedulerInfoOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerInfoOK struct {
	Payload *models.SchedulerInfoResponse
}

func (o *SchedulerInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] schedulerInfoOK  %+v", 200, o.Payload)
}
func (o *SchedulerInfoOK) GetPayload() *models.SchedulerInfoResponse {
	return o.Payload
}

func (o *SchedulerInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerInfoDefault creates a SchedulerInfoDefault with default headers values
func NewSchedulerInfoDefault(code int) *SchedulerInfoDefault {
	return &SchedulerInfoDefault{
		_statusCode: code,
	}
}

/* SchedulerInfoDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerInfoDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler info default response
func (o *SchedulerInfoDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerInfoDefault) Error() string {
	return fmt.Sprintf("[GET /info][%d] Scheduler_Info default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerInfoDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}