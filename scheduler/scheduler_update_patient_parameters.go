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
)

// NewSchedulerUpdatePatientParams creates a new SchedulerUpdatePatientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerUpdatePatientParams() *SchedulerUpdatePatientParams {
	return &SchedulerUpdatePatientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerUpdatePatientParamsWithTimeout creates a new SchedulerUpdatePatientParams object
// with the ability to set a timeout on a request.
func NewSchedulerUpdatePatientParamsWithTimeout(timeout time.Duration) *SchedulerUpdatePatientParams {
	return &SchedulerUpdatePatientParams{
		timeout: timeout,
	}
}

// NewSchedulerUpdatePatientParamsWithContext creates a new SchedulerUpdatePatientParams object
// with the ability to set a context for a request.
func NewSchedulerUpdatePatientParamsWithContext(ctx context.Context) *SchedulerUpdatePatientParams {
	return &SchedulerUpdatePatientParams{
		Context: ctx,
	}
}

// NewSchedulerUpdatePatientParamsWithHTTPClient creates a new SchedulerUpdatePatientParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerUpdatePatientParamsWithHTTPClient(client *http.Client) *SchedulerUpdatePatientParams {
	return &SchedulerUpdatePatientParams{
		HTTPClient: client,
	}
}

/* SchedulerUpdatePatientParams contains all the parameters to send to the API endpoint
   for the scheduler update patient operation.

   Typically these are written to a http.Request.
*/
type SchedulerUpdatePatientParams struct {

	// Body.
	Body SchedulerUpdatePatientBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler update patient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerUpdatePatientParams) WithDefaults() *SchedulerUpdatePatientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler update patient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerUpdatePatientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) WithTimeout(timeout time.Duration) *SchedulerUpdatePatientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) WithContext(ctx context.Context) *SchedulerUpdatePatientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) WithHTTPClient(client *http.Client) *SchedulerUpdatePatientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) WithBody(body SchedulerUpdatePatientBody) *SchedulerUpdatePatientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) SetBody(body SchedulerUpdatePatientBody) {
	o.Body = body
}

// WithID adds the id to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) WithID(id string) *SchedulerUpdatePatientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the scheduler update patient params
func (o *SchedulerUpdatePatientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerUpdatePatientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
