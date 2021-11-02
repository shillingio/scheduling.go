// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/shillingio/scheduling.go/models"
)

// NewSchedulerCreateAppointmentParams creates a new SchedulerCreateAppointmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerCreateAppointmentParams() *SchedulerCreateAppointmentParams {
	return &SchedulerCreateAppointmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerCreateAppointmentParamsWithTimeout creates a new SchedulerCreateAppointmentParams object
// with the ability to set a timeout on a request.
func NewSchedulerCreateAppointmentParamsWithTimeout(timeout time.Duration) *SchedulerCreateAppointmentParams {
	return &SchedulerCreateAppointmentParams{
		timeout: timeout,
	}
}

// NewSchedulerCreateAppointmentParamsWithContext creates a new SchedulerCreateAppointmentParams object
// with the ability to set a context for a request.
func NewSchedulerCreateAppointmentParamsWithContext(ctx context.Context) *SchedulerCreateAppointmentParams {
	return &SchedulerCreateAppointmentParams{
		Context: ctx,
	}
}

// NewSchedulerCreateAppointmentParamsWithHTTPClient creates a new SchedulerCreateAppointmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerCreateAppointmentParamsWithHTTPClient(client *http.Client) *SchedulerCreateAppointmentParams {
	return &SchedulerCreateAppointmentParams{
		HTTPClient: client,
	}
}

/* SchedulerCreateAppointmentParams contains all the parameters to send to the API endpoint
   for the scheduler create appointment operation.

   Typically these are written to a http.Request.
*/
type SchedulerCreateAppointmentParams struct {

	// Body.
	Body *models.SchedulingAppointment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler create appointment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerCreateAppointmentParams) WithDefaults() *SchedulerCreateAppointmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler create appointment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerCreateAppointmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) WithTimeout(timeout time.Duration) *SchedulerCreateAppointmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) WithContext(ctx context.Context) *SchedulerCreateAppointmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) WithHTTPClient(client *http.Client) *SchedulerCreateAppointmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) WithBody(body *models.SchedulingAppointment) *SchedulerCreateAppointmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scheduler create appointment params
func (o *SchedulerCreateAppointmentParams) SetBody(body *models.SchedulingAppointment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerCreateAppointmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
