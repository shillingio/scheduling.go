// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/shillingio/scheduling.go/models"
)

// SchedulerCreateScheduleBlockReader is a Reader for the SchedulerCreateScheduleBlock structure.
type SchedulerCreateScheduleBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulerCreateScheduleBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulerCreateScheduleBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchedulerCreateScheduleBlockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchedulerCreateScheduleBlockOK creates a SchedulerCreateScheduleBlockOK with default headers values
func NewSchedulerCreateScheduleBlockOK() *SchedulerCreateScheduleBlockOK {
	return &SchedulerCreateScheduleBlockOK{}
}

/* SchedulerCreateScheduleBlockOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchedulerCreateScheduleBlockOK struct {
	Payload *models.SchedulerScheduleBlockResponse
}

func (o *SchedulerCreateScheduleBlockOK) Error() string {
	return fmt.Sprintf("[POST /v1/schedules/{schedule.id}/blocks][%d] schedulerCreateScheduleBlockOK  %+v", 200, o.Payload)
}
func (o *SchedulerCreateScheduleBlockOK) GetPayload() *models.SchedulerScheduleBlockResponse {
	return o.Payload
}

func (o *SchedulerCreateScheduleBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerScheduleBlockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchedulerCreateScheduleBlockDefault creates a SchedulerCreateScheduleBlockDefault with default headers values
func NewSchedulerCreateScheduleBlockDefault(code int) *SchedulerCreateScheduleBlockDefault {
	return &SchedulerCreateScheduleBlockDefault{
		_statusCode: code,
	}
}

/* SchedulerCreateScheduleBlockDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchedulerCreateScheduleBlockDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the scheduler create schedule block default response
func (o *SchedulerCreateScheduleBlockDefault) Code() int {
	return o._statusCode
}

func (o *SchedulerCreateScheduleBlockDefault) Error() string {
	return fmt.Sprintf("[POST /v1/schedules/{schedule.id}/blocks][%d] Scheduler_CreateScheduleBlock default  %+v", o._statusCode, o.Payload)
}
func (o *SchedulerCreateScheduleBlockDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SchedulerCreateScheduleBlockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SchedulerCreateScheduleBlockBody scheduler create schedule block body
swagger:model SchedulerCreateScheduleBlockBody
*/
type SchedulerCreateScheduleBlockBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// day of week
	DayOfWeek int32 `json:"day_of_week,omitempty"`

	// end at
	EndAt string `json:"end_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// minimum padding
	MinimumPadding int32 `json:"minimum_padding,omitempty"`

	// parent id
	ParentID string `json:"parent_id,omitempty"`

	// provider id
	ProviderID string `json:"provider_id,omitempty"`

	// repeat
	Repeat bool `json:"repeat,omitempty"`

	// repeat limit
	RepeatLimit int32 `json:"repeat_limit,omitempty"`

	// repeat step
	RepeatStep string `json:"repeat_step,omitempty"`

	// repeat until
	RepeatUntil string `json:"repeat_until,omitempty"`

	// repeat week days
	RepeatWeekDays []int32 `json:"repeat_week_days"`

	// schedule
	Schedule *SchedulerCreateScheduleBlockParamsBodySchedule `json:"schedule,omitempty"`

	// services
	Services []*models.SchedulingProviderService `json:"services"`

	// start at
	StartAt string `json:"start_at,omitempty"`

	// types
	Types []string `json:"types"`

	// tz
	Tz string `json:"tz,omitempty"`

	// unavailable
	Unavailable bool `json:"unavailable,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduler create schedule block body
func (o *SchedulerCreateScheduleBlockBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerCreateScheduleBlockBody) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(o.Schedule) { // not required
		return nil
	}

	if o.Schedule != nil {
		if err := o.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule")
			}
			return err
		}
	}

	return nil
}

func (o *SchedulerCreateScheduleBlockBody) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(o.Services) { // not required
		return nil
	}

	for i := 0; i < len(o.Services); i++ {
		if swag.IsZero(o.Services[i]) { // not required
			continue
		}

		if o.Services[i] != nil {
			if err := o.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scheduler create schedule block body based on the context it is used
func (o *SchedulerCreateScheduleBlockBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerCreateScheduleBlockBody) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if o.Schedule != nil {
		if err := o.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule")
			}
			return err
		}
	}

	return nil
}

func (o *SchedulerCreateScheduleBlockBody) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Services); i++ {

		if o.Services[i] != nil {
			if err := o.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchedulerCreateScheduleBlockBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchedulerCreateScheduleBlockBody) UnmarshalBinary(b []byte) error {
	var res SchedulerCreateScheduleBlockBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SchedulerCreateScheduleBlockParamsBodySchedule scheduler create schedule block params body schedule
swagger:model SchedulerCreateScheduleBlockParamsBodySchedule
*/
type SchedulerCreateScheduleBlockParamsBodySchedule struct {

	// active
	Active bool `json:"active,omitempty"`

	// active from
	ActiveFrom string `json:"active_from,omitempty"`

	// active to
	ActiveTo string `json:"active_to,omitempty"`

	// blocks
	Blocks []*models.SchedulingScheduleBlock `json:"blocks"`

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
	Provider *models.SchedulingProvider `json:"provider,omitempty"`

	// timezone
	Timezone *models.SchedulingTimezone `json:"timezone,omitempty"`

	// tz
	Tz string `json:"tz,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this scheduler create schedule block params body schedule
func (o *SchedulerCreateScheduleBlockParamsBodySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) validateBlocks(formats strfmt.Registry) error {
	if swag.IsZero(o.Blocks) { // not required
		return nil
	}

	for i := 0; i < len(o.Blocks); i++ {
		if swag.IsZero(o.Blocks[i]) { // not required
			continue
		}

		if o.Blocks[i] != nil {
			if err := o.Blocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "schedule" + "." + "blocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "schedule" + "." + "blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(o.Provider) { // not required
		return nil
	}

	if o.Provider != nil {
		if err := o.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule" + "." + "provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule" + "." + "provider")
			}
			return err
		}
	}

	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(o.Timezone) { // not required
		return nil
	}

	if o.Timezone != nil {
		if err := o.Timezone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule" + "." + "timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule" + "." + "timezone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scheduler create schedule block params body schedule based on the context it is used
func (o *SchedulerCreateScheduleBlockParamsBodySchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBlocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) contextValidateBlocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Blocks); i++ {

		if o.Blocks[i] != nil {
			if err := o.Blocks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "schedule" + "." + "blocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "schedule" + "." + "blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if o.Provider != nil {
		if err := o.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule" + "." + "provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule" + "." + "provider")
			}
			return err
		}
	}

	return nil
}

func (o *SchedulerCreateScheduleBlockParamsBodySchedule) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if o.Timezone != nil {
		if err := o.Timezone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "schedule" + "." + "timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "schedule" + "." + "timezone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchedulerCreateScheduleBlockParamsBodySchedule) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchedulerCreateScheduleBlockParamsBodySchedule) UnmarshalBinary(b []byte) error {
	var res SchedulerCreateScheduleBlockParamsBodySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
