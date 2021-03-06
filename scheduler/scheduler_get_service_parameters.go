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

// NewSchedulerGetServiceParams creates a new SchedulerGetServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerGetServiceParams() *SchedulerGetServiceParams {
	return &SchedulerGetServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerGetServiceParamsWithTimeout creates a new SchedulerGetServiceParams object
// with the ability to set a timeout on a request.
func NewSchedulerGetServiceParamsWithTimeout(timeout time.Duration) *SchedulerGetServiceParams {
	return &SchedulerGetServiceParams{
		timeout: timeout,
	}
}

// NewSchedulerGetServiceParamsWithContext creates a new SchedulerGetServiceParams object
// with the ability to set a context for a request.
func NewSchedulerGetServiceParamsWithContext(ctx context.Context) *SchedulerGetServiceParams {
	return &SchedulerGetServiceParams{
		Context: ctx,
	}
}

// NewSchedulerGetServiceParamsWithHTTPClient creates a new SchedulerGetServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerGetServiceParamsWithHTTPClient(client *http.Client) *SchedulerGetServiceParams {
	return &SchedulerGetServiceParams{
		HTTPClient: client,
	}
}

/* SchedulerGetServiceParams contains all the parameters to send to the API endpoint
   for the scheduler get service operation.

   Typically these are written to a http.Request.
*/
type SchedulerGetServiceParams struct {

	// Active.
	Active *bool

	// ConfigDurations.
	ConfigDurations []int32

	// ConfigFixedDuration.
	//
	// Format: int32
	ConfigFixedDuration *int32

	// CreatedAt.
	CreatedAt *string

	// CreatedBy.
	CreatedBy *string

	// Description.
	Description *string

	// ID.
	ID string

	// Name.
	Name *string

	// OrgID.
	OrgID *string

	// RefID.
	RefID *string

	// UpdatedAt.
	UpdatedAt *string

	// UpdatedBy.
	UpdatedBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler get service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetServiceParams) WithDefaults() *SchedulerGetServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler get service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerGetServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler get service params
func (o *SchedulerGetServiceParams) WithTimeout(timeout time.Duration) *SchedulerGetServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler get service params
func (o *SchedulerGetServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler get service params
func (o *SchedulerGetServiceParams) WithContext(ctx context.Context) *SchedulerGetServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler get service params
func (o *SchedulerGetServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler get service params
func (o *SchedulerGetServiceParams) WithHTTPClient(client *http.Client) *SchedulerGetServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler get service params
func (o *SchedulerGetServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActive adds the active to the scheduler get service params
func (o *SchedulerGetServiceParams) WithActive(active *bool) *SchedulerGetServiceParams {
	o.SetActive(active)
	return o
}

// SetActive adds the active to the scheduler get service params
func (o *SchedulerGetServiceParams) SetActive(active *bool) {
	o.Active = active
}

// WithConfigDurations adds the configDurations to the scheduler get service params
func (o *SchedulerGetServiceParams) WithConfigDurations(configDurations []int32) *SchedulerGetServiceParams {
	o.SetConfigDurations(configDurations)
	return o
}

// SetConfigDurations adds the configDurations to the scheduler get service params
func (o *SchedulerGetServiceParams) SetConfigDurations(configDurations []int32) {
	o.ConfigDurations = configDurations
}

// WithConfigFixedDuration adds the configFixedDuration to the scheduler get service params
func (o *SchedulerGetServiceParams) WithConfigFixedDuration(configFixedDuration *int32) *SchedulerGetServiceParams {
	o.SetConfigFixedDuration(configFixedDuration)
	return o
}

// SetConfigFixedDuration adds the configFixedDuration to the scheduler get service params
func (o *SchedulerGetServiceParams) SetConfigFixedDuration(configFixedDuration *int32) {
	o.ConfigFixedDuration = configFixedDuration
}

// WithCreatedAt adds the createdAt to the scheduler get service params
func (o *SchedulerGetServiceParams) WithCreatedAt(createdAt *string) *SchedulerGetServiceParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the scheduler get service params
func (o *SchedulerGetServiceParams) SetCreatedAt(createdAt *string) {
	o.CreatedAt = createdAt
}

// WithCreatedBy adds the createdBy to the scheduler get service params
func (o *SchedulerGetServiceParams) WithCreatedBy(createdBy *string) *SchedulerGetServiceParams {
	o.SetCreatedBy(createdBy)
	return o
}

// SetCreatedBy adds the createdBy to the scheduler get service params
func (o *SchedulerGetServiceParams) SetCreatedBy(createdBy *string) {
	o.CreatedBy = createdBy
}

// WithDescription adds the description to the scheduler get service params
func (o *SchedulerGetServiceParams) WithDescription(description *string) *SchedulerGetServiceParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the scheduler get service params
func (o *SchedulerGetServiceParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the scheduler get service params
func (o *SchedulerGetServiceParams) WithID(id string) *SchedulerGetServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the scheduler get service params
func (o *SchedulerGetServiceParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the scheduler get service params
func (o *SchedulerGetServiceParams) WithName(name *string) *SchedulerGetServiceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the scheduler get service params
func (o *SchedulerGetServiceParams) SetName(name *string) {
	o.Name = name
}

// WithOrgID adds the orgID to the scheduler get service params
func (o *SchedulerGetServiceParams) WithOrgID(orgID *string) *SchedulerGetServiceParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the scheduler get service params
func (o *SchedulerGetServiceParams) SetOrgID(orgID *string) {
	o.OrgID = orgID
}

// WithRefID adds the refID to the scheduler get service params
func (o *SchedulerGetServiceParams) WithRefID(refID *string) *SchedulerGetServiceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the scheduler get service params
func (o *SchedulerGetServiceParams) SetRefID(refID *string) {
	o.RefID = refID
}

// WithUpdatedAt adds the updatedAt to the scheduler get service params
func (o *SchedulerGetServiceParams) WithUpdatedAt(updatedAt *string) *SchedulerGetServiceParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the scheduler get service params
func (o *SchedulerGetServiceParams) SetUpdatedAt(updatedAt *string) {
	o.UpdatedAt = updatedAt
}

// WithUpdatedBy adds the updatedBy to the scheduler get service params
func (o *SchedulerGetServiceParams) WithUpdatedBy(updatedBy *string) *SchedulerGetServiceParams {
	o.SetUpdatedBy(updatedBy)
	return o
}

// SetUpdatedBy adds the updatedBy to the scheduler get service params
func (o *SchedulerGetServiceParams) SetUpdatedBy(updatedBy *string) {
	o.UpdatedBy = updatedBy
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerGetServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Active != nil {

		// query param active
		var qrActive bool

		if o.Active != nil {
			qrActive = *o.Active
		}
		qActive := swag.FormatBool(qrActive)
		if qActive != "" {

			if err := r.SetQueryParam("active", qActive); err != nil {
				return err
			}
		}
	}

	if o.ConfigDurations != nil {

		// binding items for config.durations
		joinedConfigDurations := o.bindParamConfigDurations(reg)

		// query array param config.durations
		if err := r.SetQueryParam("config.durations", joinedConfigDurations...); err != nil {
			return err
		}
	}

	if o.ConfigFixedDuration != nil {

		// query param config.fixed_duration
		var qrConfigFixedDuration int32

		if o.ConfigFixedDuration != nil {
			qrConfigFixedDuration = *o.ConfigFixedDuration
		}
		qConfigFixedDuration := swag.FormatInt32(qrConfigFixedDuration)
		if qConfigFixedDuration != "" {

			if err := r.SetQueryParam("config.fixed_duration", qConfigFixedDuration); err != nil {
				return err
			}
		}
	}

	if o.CreatedAt != nil {

		// query param created_at
		var qrCreatedAt string

		if o.CreatedAt != nil {
			qrCreatedAt = *o.CreatedAt
		}
		qCreatedAt := qrCreatedAt
		if qCreatedAt != "" {

			if err := r.SetQueryParam("created_at", qCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.CreatedBy != nil {

		// query param created_by
		var qrCreatedBy string

		if o.CreatedBy != nil {
			qrCreatedBy = *o.CreatedBy
		}
		qCreatedBy := qrCreatedBy
		if qCreatedBy != "" {

			if err := r.SetQueryParam("created_by", qCreatedBy); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrgID != nil {

		// query param org_id
		var qrOrgID string

		if o.OrgID != nil {
			qrOrgID = *o.OrgID
		}
		qOrgID := qrOrgID
		if qOrgID != "" {

			if err := r.SetQueryParam("org_id", qOrgID); err != nil {
				return err
			}
		}
	}

	if o.RefID != nil {

		// query param ref_id
		var qrRefID string

		if o.RefID != nil {
			qrRefID = *o.RefID
		}
		qRefID := qrRefID
		if qRefID != "" {

			if err := r.SetQueryParam("ref_id", qRefID); err != nil {
				return err
			}
		}
	}

	if o.UpdatedAt != nil {

		// query param updated_at
		var qrUpdatedAt string

		if o.UpdatedAt != nil {
			qrUpdatedAt = *o.UpdatedAt
		}
		qUpdatedAt := qrUpdatedAt
		if qUpdatedAt != "" {

			if err := r.SetQueryParam("updated_at", qUpdatedAt); err != nil {
				return err
			}
		}
	}

	if o.UpdatedBy != nil {

		// query param updated_by
		var qrUpdatedBy string

		if o.UpdatedBy != nil {
			qrUpdatedBy = *o.UpdatedBy
		}
		qUpdatedBy := qrUpdatedBy
		if qUpdatedBy != "" {

			if err := r.SetQueryParam("updated_by", qUpdatedBy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSchedulerGetService binds the parameter config.durations
func (o *SchedulerGetServiceParams) bindParamConfigDurations(formats strfmt.Registry) []string {
	configDurationsIR := o.ConfigDurations

	var configDurationsIC []string
	for _, configDurationsIIR := range configDurationsIR { // explode []int32

		configDurationsIIV := swag.FormatInt32(configDurationsIIR) // int32 as string
		configDurationsIC = append(configDurationsIC, configDurationsIIV)
	}

	// items.CollectionFormat: "multi"
	configDurationsIS := swag.JoinByFormat(configDurationsIC, "multi")

	return configDurationsIS
}
