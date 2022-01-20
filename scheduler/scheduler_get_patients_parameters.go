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

// NewSchedulerGetPatientsParams creates a new SchedulerGetPatientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerGetPatientsParams() *SchedulerGetPatientsParams {
	return &SchedulerGetPatientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerGetPatientsParamsWithTimeout creates a new SchedulerGetPatientsParams object
// with the ability to set a timeout on a request.
func NewSchedulerGetPatientsParamsWithTimeout(timeout time.Duration) *SchedulerGetPatientsParams {
	return &SchedulerGetPatientsParams{
		timeout: timeout,
	}
}

// NewSchedulerGetPatientsParamsWithContext creates a new SchedulerGetPatientsParams object
// with the ability to set a context for a request.
func NewSchedulerGetPatientsParamsWithContext(ctx context.Context) *SchedulerGetPatientsParams {
	return &SchedulerGetPatientsParams{
		Context: ctx,
	}
}

// NewSchedulerGetPatientsParamsWithHTTPClient creates a new SchedulerGetPatientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerGetPatientsParamsWithHTTPClient(client *http.Client) *SchedulerGetPatientsParams {
	return &SchedulerGetPatientsParams{
		HTTPClient: client,
	}
}

/* SchedulerGetPatientsParams contains all the parameters to send to the API endpoint
   for the scheduler get patients operation.

   Typically these are written to a http.Request.
*/
type SchedulerGetPatientsParams struct {

	// ID.
	ID *string

	// Ids.
	Ids []int32

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// OrganizationID.
	OrganizationID *string

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

	// ProviderID.
	ProviderID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler get patients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetPatientsParams) WithDefaults() *SchedulerGetPatientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler get patients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetPatientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithTimeout(timeout time.Duration) *SchedulerGetPatientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithContext(ctx context.Context) *SchedulerGetPatientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithHTTPClient(client *http.Client) *SchedulerGetPatientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithID(id *string) *SchedulerGetPatientsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetID(id *string) {
	o.ID = id
}

// WithIds adds the ids to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithIds(ids []int32) *SchedulerGetPatientsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetIds(ids []int32) {
	o.Ids = ids
}

// WithLimit adds the limit to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithLimit(limit *int32) *SchedulerGetPatientsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithOffset(offset *int32) *SchedulerGetPatientsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithOrganizationID(organizationID *string) *SchedulerGetPatientsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithPatientActive adds the patientActive to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientActive(patientActive *bool) *SchedulerGetPatientsParams {
	o.SetPatientActive(patientActive)
	return o
}

// SetPatientActive adds the patientActive to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientActive(patientActive *bool) {
	o.PatientActive = patientActive
}

// WithPatientCreatedAt adds the patientCreatedAt to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientCreatedAt(patientCreatedAt *string) *SchedulerGetPatientsParams {
	o.SetPatientCreatedAt(patientCreatedAt)
	return o
}

// SetPatientCreatedAt adds the patientCreatedAt to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientCreatedAt(patientCreatedAt *string) {
	o.PatientCreatedAt = patientCreatedAt
}

// WithPatientCreatedBy adds the patientCreatedBy to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientCreatedBy(patientCreatedBy *string) *SchedulerGetPatientsParams {
	o.SetPatientCreatedBy(patientCreatedBy)
	return o
}

// SetPatientCreatedBy adds the patientCreatedBy to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientCreatedBy(patientCreatedBy *string) {
	o.PatientCreatedBy = patientCreatedBy
}

// WithPatientGivenName adds the patientGivenName to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientGivenName(patientGivenName *string) *SchedulerGetPatientsParams {
	o.SetPatientGivenName(patientGivenName)
	return o
}

// SetPatientGivenName adds the patientGivenName to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientGivenName(patientGivenName *string) {
	o.PatientGivenName = patientGivenName
}

// WithPatientID adds the patientID to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientID(patientID *string) *SchedulerGetPatientsParams {
	o.SetPatientID(patientID)
	return o
}

// SetPatientID adds the patientId to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientID(patientID *string) {
	o.PatientID = patientID
}

// WithPatientOrgID adds the patientOrgID to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientOrgID(patientOrgID *string) *SchedulerGetPatientsParams {
	o.SetPatientOrgID(patientOrgID)
	return o
}

// SetPatientOrgID adds the patientOrgId to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientOrgID(patientOrgID *string) {
	o.PatientOrgID = patientOrgID
}

// WithPatientPrefix adds the patientPrefix to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientPrefix(patientPrefix *string) *SchedulerGetPatientsParams {
	o.SetPatientPrefix(patientPrefix)
	return o
}

// SetPatientPrefix adds the patientPrefix to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientPrefix(patientPrefix *string) {
	o.PatientPrefix = patientPrefix
}

// WithPatientRefID adds the patientRefID to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientRefID(patientRefID *string) *SchedulerGetPatientsParams {
	o.SetPatientRefID(patientRefID)
	return o
}

// SetPatientRefID adds the patientRefId to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientRefID(patientRefID *string) {
	o.PatientRefID = patientRefID
}

// WithPatientSuffix adds the patientSuffix to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientSuffix(patientSuffix *string) *SchedulerGetPatientsParams {
	o.SetPatientSuffix(patientSuffix)
	return o
}

// SetPatientSuffix adds the patientSuffix to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientSuffix(patientSuffix *string) {
	o.PatientSuffix = patientSuffix
}

// WithPatientSurname adds the patientSurname to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientSurname(patientSurname *string) *SchedulerGetPatientsParams {
	o.SetPatientSurname(patientSurname)
	return o
}

// SetPatientSurname adds the patientSurname to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientSurname(patientSurname *string) {
	o.PatientSurname = patientSurname
}

// WithPatientUpdatedAt adds the patientUpdatedAt to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientUpdatedAt(patientUpdatedAt *string) *SchedulerGetPatientsParams {
	o.SetPatientUpdatedAt(patientUpdatedAt)
	return o
}

// SetPatientUpdatedAt adds the patientUpdatedAt to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientUpdatedAt(patientUpdatedAt *string) {
	o.PatientUpdatedAt = patientUpdatedAt
}

// WithPatientUpdatedBy adds the patientUpdatedBy to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithPatientUpdatedBy(patientUpdatedBy *string) *SchedulerGetPatientsParams {
	o.SetPatientUpdatedBy(patientUpdatedBy)
	return o
}

// SetPatientUpdatedBy adds the patientUpdatedBy to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetPatientUpdatedBy(patientUpdatedBy *string) {
	o.PatientUpdatedBy = patientUpdatedBy
}

// WithProviderID adds the providerID to the scheduler get patients params
func (o *SchedulerGetPatientsParams) WithProviderID(providerID *string) *SchedulerGetPatientsParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the scheduler get patients params
func (o *SchedulerGetPatientsParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerGetPatientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrganizationID != nil {

		// query param organization_id
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organization_id", qOrganizationID); err != nil {
				return err
			}
		}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSchedulerGetPatients binds the parameter ids
func (o *SchedulerGetPatientsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []int32

		idsIIV := swag.FormatInt32(idsIIR) // int32 as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
