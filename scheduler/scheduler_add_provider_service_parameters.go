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

// NewSchedulerAddProviderServiceParams creates a new SchedulerAddProviderServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerAddProviderServiceParams() *SchedulerAddProviderServiceParams {
	return &SchedulerAddProviderServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerAddProviderServiceParamsWithTimeout creates a new SchedulerAddProviderServiceParams object
// with the ability to set a timeout on a request.
func NewSchedulerAddProviderServiceParamsWithTimeout(timeout time.Duration) *SchedulerAddProviderServiceParams {
	return &SchedulerAddProviderServiceParams{
		timeout: timeout,
	}
}

// NewSchedulerAddProviderServiceParamsWithContext creates a new SchedulerAddProviderServiceParams object
// with the ability to set a context for a request.
func NewSchedulerAddProviderServiceParamsWithContext(ctx context.Context) *SchedulerAddProviderServiceParams {
	return &SchedulerAddProviderServiceParams{
		Context: ctx,
	}
}

// NewSchedulerAddProviderServiceParamsWithHTTPClient creates a new SchedulerAddProviderServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerAddProviderServiceParamsWithHTTPClient(client *http.Client) *SchedulerAddProviderServiceParams {
	return &SchedulerAddProviderServiceParams{
		HTTPClient: client,
	}
}

/* SchedulerAddProviderServiceParams contains all the parameters to send to the API endpoint
   for the scheduler add provider service operation.

   Typically these are written to a http.Request.
*/
type SchedulerAddProviderServiceParams struct {

	// Body.
	Body *models.SchedulingProviderService

	// ProviderID.
	ProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler add provider service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerAddProviderServiceParams) WithDefaults() *SchedulerAddProviderServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler add provider service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerAddProviderServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) WithTimeout(timeout time.Duration) *SchedulerAddProviderServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) WithContext(ctx context.Context) *SchedulerAddProviderServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) WithHTTPClient(client *http.Client) *SchedulerAddProviderServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) WithBody(body *models.SchedulingProviderService) *SchedulerAddProviderServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) SetBody(body *models.SchedulingProviderService) {
	o.Body = body
}

// WithProviderID adds the providerID to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) WithProviderID(providerID string) *SchedulerAddProviderServiceParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the scheduler add provider service params
func (o *SchedulerAddProviderServiceParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerAddProviderServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param provider.id
	if err := r.SetPathParam("provider.id", o.ProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
