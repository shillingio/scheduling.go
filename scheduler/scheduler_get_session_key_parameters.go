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
	"github.com/go-openapi/swag"
)

// NewSchedulerGetSessionKeyParams creates a new SchedulerGetSessionKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerGetSessionKeyParams() *SchedulerGetSessionKeyParams {
	return &SchedulerGetSessionKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerGetSessionKeyParamsWithTimeout creates a new SchedulerGetSessionKeyParams object
// with the ability to set a timeout on a request.
func NewSchedulerGetSessionKeyParamsWithTimeout(timeout time.Duration) *SchedulerGetSessionKeyParams {
	return &SchedulerGetSessionKeyParams{
		timeout: timeout,
	}
}

// NewSchedulerGetSessionKeyParamsWithContext creates a new SchedulerGetSessionKeyParams object
// with the ability to set a context for a request.
func NewSchedulerGetSessionKeyParamsWithContext(ctx context.Context) *SchedulerGetSessionKeyParams {
	return &SchedulerGetSessionKeyParams{
		Context: ctx,
	}
}

// NewSchedulerGetSessionKeyParamsWithHTTPClient creates a new SchedulerGetSessionKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerGetSessionKeyParamsWithHTTPClient(client *http.Client) *SchedulerGetSessionKeyParams {
	return &SchedulerGetSessionKeyParams{
		HTTPClient: client,
	}
}

/* SchedulerGetSessionKeyParams contains all the parameters to send to the API endpoint
   for the scheduler get session key operation.

   Typically these are written to a http.Request.
*/
type SchedulerGetSessionKeyParams struct {

	// AppointmentID.
	AppointmentID *string

	// Key.
	Key string

	// PatientActive.
	PatientActive *bool

	// PatientCreatedAt.
	PatientCreatedAt *string

	// PatientCreatedBy.
	PatientCreatedBy *string

	// PatientGivenName.
	PatientGivenName *string

	// PatientID.
	PatientID *string

	// PatientOrgID.
	PatientOrgID *string

	// PatientPrefix.
	PatientPrefix *string

	// PatientRefID.
	PatientRefID *string

	// PatientSuffix.
	PatientSuffix *string

	// PatientSurname.
	PatientSurname *string

	// PatientUpdatedAt.
	PatientUpdatedAt *string

	// PatientUpdatedBy.
	PatientUpdatedBy *string

	// PrivateKey.
	PrivateKey *string

	// ProviderID.
	ProviderID *string

	// RangeFrom.
	RangeFrom *string

	// RangeTo.
	RangeTo *string

	// ServiceID.
	ServiceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler get session key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetSessionKeyParams) WithDefaults() *SchedulerGetSessionKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler get session key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetSessionKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithTimeout(timeout time.Duration) *SchedulerGetSessionKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithContext(ctx context.Context) *SchedulerGetSessionKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithHTTPClient(client *http.Client) *SchedulerGetSessionKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithAppointmentID(appointmentID *string) *SchedulerGetSessionKeyParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetAppointmentID(appointmentID *string) {
	o.AppointmentID = appointmentID
}

// WithKey adds the key to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithKey(key string) *SchedulerGetSessionKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetKey(key string) {
	o.Key = key
}

// WithPatientActive adds the patientActive to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientActive(patientActive *bool) *SchedulerGetSessionKeyParams {
	o.SetPatientActive(patientActive)
	return o
}

// SetPatientActive adds the patientActive to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientActive(patientActive *bool) {
	o.PatientActive = patientActive
}

// WithPatientCreatedAt adds the patientCreatedAt to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientCreatedAt(patientCreatedAt *string) *SchedulerGetSessionKeyParams {
	o.SetPatientCreatedAt(patientCreatedAt)
	return o
}

// SetPatientCreatedAt adds the patientCreatedAt to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientCreatedAt(patientCreatedAt *string) {
	o.PatientCreatedAt = patientCreatedAt
}

// WithPatientCreatedBy adds the patientCreatedBy to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientCreatedBy(patientCreatedBy *string) *SchedulerGetSessionKeyParams {
	o.SetPatientCreatedBy(patientCreatedBy)
	return o
}

// SetPatientCreatedBy adds the patientCreatedBy to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientCreatedBy(patientCreatedBy *string) {
	o.PatientCreatedBy = patientCreatedBy
}

// WithPatientGivenName adds the patientGivenName to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientGivenName(patientGivenName *string) *SchedulerGetSessionKeyParams {
	o.SetPatientGivenName(patientGivenName)
	return o
}

// SetPatientGivenName adds the patientGivenName to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientGivenName(patientGivenName *string) {
	o.PatientGivenName = patientGivenName
}

// WithPatientID adds the patientID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientID(patientID *string) *SchedulerGetSessionKeyParams {
	o.SetPatientID(patientID)
	return o
}

// SetPatientID adds the patientId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientID(patientID *string) {
	o.PatientID = patientID
}

// WithPatientOrgID adds the patientOrgID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientOrgID(patientOrgID *string) *SchedulerGetSessionKeyParams {
	o.SetPatientOrgID(patientOrgID)
	return o
}

// SetPatientOrgID adds the patientOrgId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientOrgID(patientOrgID *string) {
	o.PatientOrgID = patientOrgID
}

// WithPatientPrefix adds the patientPrefix to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientPrefix(patientPrefix *string) *SchedulerGetSessionKeyParams {
	o.SetPatientPrefix(patientPrefix)
	return o
}

// SetPatientPrefix adds the patientPrefix to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientPrefix(patientPrefix *string) {
	o.PatientPrefix = patientPrefix
}

// WithPatientRefID adds the patientRefID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientRefID(patientRefID *string) *SchedulerGetSessionKeyParams {
	o.SetPatientRefID(patientRefID)
	return o
}

// SetPatientRefID adds the patientRefId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientRefID(patientRefID *string) {
	o.PatientRefID = patientRefID
}

// WithPatientSuffix adds the patientSuffix to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientSuffix(patientSuffix *string) *SchedulerGetSessionKeyParams {
	o.SetPatientSuffix(patientSuffix)
	return o
}

// SetPatientSuffix adds the patientSuffix to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientSuffix(patientSuffix *string) {
	o.PatientSuffix = patientSuffix
}

// WithPatientSurname adds the patientSurname to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientSurname(patientSurname *string) *SchedulerGetSessionKeyParams {
	o.SetPatientSurname(patientSurname)
	return o
}

// SetPatientSurname adds the patientSurname to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientSurname(patientSurname *string) {
	o.PatientSurname = patientSurname
}

// WithPatientUpdatedAt adds the patientUpdatedAt to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientUpdatedAt(patientUpdatedAt *string) *SchedulerGetSessionKeyParams {
	o.SetPatientUpdatedAt(patientUpdatedAt)
	return o
}

// SetPatientUpdatedAt adds the patientUpdatedAt to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientUpdatedAt(patientUpdatedAt *string) {
	o.PatientUpdatedAt = patientUpdatedAt
}

// WithPatientUpdatedBy adds the patientUpdatedBy to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPatientUpdatedBy(patientUpdatedBy *string) *SchedulerGetSessionKeyParams {
	o.SetPatientUpdatedBy(patientUpdatedBy)
	return o
}

// SetPatientUpdatedBy adds the patientUpdatedBy to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPatientUpdatedBy(patientUpdatedBy *string) {
	o.PatientUpdatedBy = patientUpdatedBy
}

// WithPrivateKey adds the privateKey to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithPrivateKey(privateKey *string) *SchedulerGetSessionKeyParams {
	o.SetPrivateKey(privateKey)
	return o
}

// SetPrivateKey adds the privateKey to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetPrivateKey(privateKey *string) {
	o.PrivateKey = privateKey
}

// WithProviderID adds the providerID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithProviderID(providerID *string) *SchedulerGetSessionKeyParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithRangeFrom adds the rangeFrom to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithRangeFrom(rangeFrom *string) *SchedulerGetSessionKeyParams {
	o.SetRangeFrom(rangeFrom)
	return o
}

// SetRangeFrom adds the rangeFrom to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetRangeFrom(rangeFrom *string) {
	o.RangeFrom = rangeFrom
}

// WithRangeTo adds the rangeTo to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithRangeTo(rangeTo *string) *SchedulerGetSessionKeyParams {
	o.SetRangeTo(rangeTo)
	return o
}

// SetRangeTo adds the rangeTo to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetRangeTo(rangeTo *string) {
	o.RangeTo = rangeTo
}

// WithServiceID adds the serviceID to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) WithServiceID(serviceID *string) *SchedulerGetSessionKeyParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the scheduler get session key params
func (o *SchedulerGetSessionKeyParams) SetServiceID(serviceID *string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerGetSessionKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppointmentID != nil {

		// query param appointment_id
		var qrAppointmentID string

		if o.AppointmentID != nil {
			qrAppointmentID = *o.AppointmentID
		}
		qAppointmentID := qrAppointmentID
		if qAppointmentID != "" {

			if err := r.SetQueryParam("appointment_id", qAppointmentID); err != nil {
				return err
			}
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if o.PatientActive != nil {

		// query param patient.active
		var qrPatientActive bool

		if o.PatientActive != nil {
			qrPatientActive = *o.PatientActive
		}
		qPatientActive := swag.FormatBool(qrPatientActive)
		if qPatientActive != "" {

			if err := r.SetQueryParam("patient.active", qPatientActive); err != nil {
				return err
			}
		}
	}

	if o.PatientCreatedAt != nil {

		// query param patient.created_at
		var qrPatientCreatedAt string

		if o.PatientCreatedAt != nil {
			qrPatientCreatedAt = *o.PatientCreatedAt
		}
		qPatientCreatedAt := qrPatientCreatedAt
		if qPatientCreatedAt != "" {

			if err := r.SetQueryParam("patient.created_at", qPatientCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.PatientCreatedBy != nil {

		// query param patient.created_by
		var qrPatientCreatedBy string

		if o.PatientCreatedBy != nil {
			qrPatientCreatedBy = *o.PatientCreatedBy
		}
		qPatientCreatedBy := qrPatientCreatedBy
		if qPatientCreatedBy != "" {

			if err := r.SetQueryParam("patient.created_by", qPatientCreatedBy); err != nil {
				return err
			}
		}
	}

	if o.PatientGivenName != nil {

		// query param patient.given_name
		var qrPatientGivenName string

		if o.PatientGivenName != nil {
			qrPatientGivenName = *o.PatientGivenName
		}
		qPatientGivenName := qrPatientGivenName
		if qPatientGivenName != "" {

			if err := r.SetQueryParam("patient.given_name", qPatientGivenName); err != nil {
				return err
			}
		}
	}

	if o.PatientID != nil {

		// query param patient.id
		var qrPatientID string

		if o.PatientID != nil {
			qrPatientID = *o.PatientID
		}
		qPatientID := qrPatientID
		if qPatientID != "" {

			if err := r.SetQueryParam("patient.id", qPatientID); err != nil {
				return err
			}
		}
	}

	if o.PatientOrgID != nil {

		// query param patient.org_id
		var qrPatientOrgID string

		if o.PatientOrgID != nil {
			qrPatientOrgID = *o.PatientOrgID
		}
		qPatientOrgID := qrPatientOrgID
		if qPatientOrgID != "" {

			if err := r.SetQueryParam("patient.org_id", qPatientOrgID); err != nil {
				return err
			}
		}
	}

	if o.PatientPrefix != nil {

		// query param patient.prefix
		var qrPatientPrefix string

		if o.PatientPrefix != nil {
			qrPatientPrefix = *o.PatientPrefix
		}
		qPatientPrefix := qrPatientPrefix
		if qPatientPrefix != "" {

			if err := r.SetQueryParam("patient.prefix", qPatientPrefix); err != nil {
				return err
			}
		}
	}

	if o.PatientRefID != nil {

		// query param patient.ref_id
		var qrPatientRefID string

		if o.PatientRefID != nil {
			qrPatientRefID = *o.PatientRefID
		}
		qPatientRefID := qrPatientRefID
		if qPatientRefID != "" {

			if err := r.SetQueryParam("patient.ref_id", qPatientRefID); err != nil {
				return err
			}
		}
	}

	if o.PatientSuffix != nil {

		// query param patient.suffix
		var qrPatientSuffix string

		if o.PatientSuffix != nil {
			qrPatientSuffix = *o.PatientSuffix
		}
		qPatientSuffix := qrPatientSuffix
		if qPatientSuffix != "" {

			if err := r.SetQueryParam("patient.suffix", qPatientSuffix); err != nil {
				return err
			}
		}
	}

	if o.PatientSurname != nil {

		// query param patient.surname
		var qrPatientSurname string

		if o.PatientSurname != nil {
			qrPatientSurname = *o.PatientSurname
		}
		qPatientSurname := qrPatientSurname
		if qPatientSurname != "" {

			if err := r.SetQueryParam("patient.surname", qPatientSurname); err != nil {
				return err
			}
		}
	}

	if o.PatientUpdatedAt != nil {

		// query param patient.updated_at
		var qrPatientUpdatedAt string

		if o.PatientUpdatedAt != nil {
			qrPatientUpdatedAt = *o.PatientUpdatedAt
		}
		qPatientUpdatedAt := qrPatientUpdatedAt
		if qPatientUpdatedAt != "" {

			if err := r.SetQueryParam("patient.updated_at", qPatientUpdatedAt); err != nil {
				return err
			}
		}
	}

	if o.PatientUpdatedBy != nil {

		// query param patient.updated_by
		var qrPatientUpdatedBy string

		if o.PatientUpdatedBy != nil {
			qrPatientUpdatedBy = *o.PatientUpdatedBy
		}
		qPatientUpdatedBy := qrPatientUpdatedBy
		if qPatientUpdatedBy != "" {

			if err := r.SetQueryParam("patient.updated_by", qPatientUpdatedBy); err != nil {
				return err
			}
		}
	}

	if o.PrivateKey != nil {

		// query param private_key
		var qrPrivateKey string

		if o.PrivateKey != nil {
			qrPrivateKey = *o.PrivateKey
		}
		qPrivateKey := qrPrivateKey
		if qPrivateKey != "" {

			if err := r.SetQueryParam("private_key", qPrivateKey); err != nil {
				return err
			}
		}
	}

	if o.ProviderID != nil {

		// query param provider_id
		var qrProviderID string

		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {

			if err := r.SetQueryParam("provider_id", qProviderID); err != nil {
				return err
			}
		}
	}

	if o.RangeFrom != nil {

		// query param range_from
		var qrRangeFrom string

		if o.RangeFrom != nil {
			qrRangeFrom = *o.RangeFrom
		}
		qRangeFrom := qrRangeFrom
		if qRangeFrom != "" {

			if err := r.SetQueryParam("range_from", qRangeFrom); err != nil {
				return err
			}
		}
	}

	if o.RangeTo != nil {

		// query param range_to
		var qrRangeTo string

		if o.RangeTo != nil {
			qrRangeTo = *o.RangeTo
		}
		qRangeTo := qrRangeTo
		if qRangeTo != "" {

			if err := r.SetQueryParam("range_to", qRangeTo); err != nil {
				return err
			}
		}
	}

	if o.ServiceID != nil {

		// query param service_id
		var qrServiceID string

		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := qrServiceID
		if qServiceID != "" {

			if err := r.SetQueryParam("service_id", qServiceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}