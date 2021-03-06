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

	"github.com/shillingio/scheduling.go/models"
)

// NewSchedulerSearchAvailabilityParams creates a new SchedulerSearchAvailabilityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerSearchAvailabilityParams() *SchedulerSearchAvailabilityParams {
	return &SchedulerSearchAvailabilityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerSearchAvailabilityParamsWithTimeout creates a new SchedulerSearchAvailabilityParams object
// with the ability to set a timeout on a request.
func NewSchedulerSearchAvailabilityParamsWithTimeout(timeout time.Duration) *SchedulerSearchAvailabilityParams {
	return &SchedulerSearchAvailabilityParams{
		timeout: timeout,
	}
}

// NewSchedulerSearchAvailabilityParamsWithContext creates a new SchedulerSearchAvailabilityParams object
// with the ability to set a context for a request.
func NewSchedulerSearchAvailabilityParamsWithContext(ctx context.Context) *SchedulerSearchAvailabilityParams {
	return &SchedulerSearchAvailabilityParams{
		Context: ctx,
	}
}

// NewSchedulerSearchAvailabilityParamsWithHTTPClient creates a new SchedulerSearchAvailabilityParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerSearchAvailabilityParamsWithHTTPClient(client *http.Client) *SchedulerSearchAvailabilityParams {
	return &SchedulerSearchAvailabilityParams{
		HTTPClient: client,
	}
}

/* SchedulerSearchAvailabilityParams contains all the parameters to send to the API endpoint
   for the scheduler search availability operation.

   Typically these are written to a http.Request.
*/
type SchedulerSearchAvailabilityParams struct {

	// Availability.
	Availability *models.SchedulerAvailabilitySearchRequest

	// ID.
	ID *string

	// OrganizationID.
	OrganizationID *string

	/* ProviderDuration.

	   The duration of the appointment in minutes

	   Format: int32
	*/
	ProviderDuration *int32

	/* ProviderExclude.

	     The ID of a specific providers to exclude those providers from search
	results
	*/
	ProviderExclude []string

	/* ProviderInPerson.

	   Return only in-person providers (default: false)
	*/
	ProviderInPerson *bool

	/* ProviderInclude.

	   The IDs of specific providers in which to search for availability
	*/
	ProviderInclude []string

	/* ProviderLimit.

	   Limit the number of results to return. (default: 10)

	   Format: int32
	*/
	ProviderLimit *int32

	/* ProviderLocationLatitude.

	   Specific latitude point to use

	   Format: float
	*/
	ProviderLocationLatitude *float32

	/* ProviderLocationLongitude.

	   Specific longitude point to use

	   Format: float
	*/
	ProviderLocationLongitude *float32

	/* ProviderLocationQuery.

	   Address query (ex: 123 Main St., San Diego, CA, 92101, San Diego, CA, etc.)
	*/
	ProviderLocationQuery *string

	/* ProviderLocationZip.

	   Zip code
	*/
	ProviderLocationZip *string

	/* ProviderOffset.

	     Offset the results count if pagination is necessary based on the limit and
	total fields in the response

	     Format: int32
	*/
	ProviderOffset *int32

	/* ProviderRangeFrom.

	   The start of the requested date/time range in RFC3339 format
	*/
	ProviderRangeFrom *string

	/* ProviderRangeTo.

	   The end of the requested date/time range in RFC3339 format
	*/
	ProviderRangeTo *string

	/* ProviderServiceID.

	   The service UUID requested by the patient
	*/
	ProviderServiceID *string

	/* ProviderState.

	   State the provider is licensed in for the requested service
	*/
	ProviderState *string

	/* ProviderZip.

	   ZIP code the patient resides in
	*/
	ProviderZip *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler search availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerSearchAvailabilityParams) WithDefaults() *SchedulerSearchAvailabilityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler search availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerSearchAvailabilityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithTimeout(timeout time.Duration) *SchedulerSearchAvailabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithContext(ctx context.Context) *SchedulerSearchAvailabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithHTTPClient(client *http.Client) *SchedulerSearchAvailabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailability adds the availability to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithAvailability(availability *models.SchedulerAvailabilitySearchRequest) *SchedulerSearchAvailabilityParams {
	o.SetAvailability(availability)
	return o
}

// SetAvailability adds the availability to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetAvailability(availability *models.SchedulerAvailabilitySearchRequest) {
	o.Availability = availability
}

// WithID adds the id to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithID(id *string) *SchedulerSearchAvailabilityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetID(id *string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithOrganizationID(organizationID *string) *SchedulerSearchAvailabilityParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithProviderDuration adds the providerDuration to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderDuration(providerDuration *int32) *SchedulerSearchAvailabilityParams {
	o.SetProviderDuration(providerDuration)
	return o
}

// SetProviderDuration adds the providerDuration to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderDuration(providerDuration *int32) {
	o.ProviderDuration = providerDuration
}

// WithProviderExclude adds the providerExclude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderExclude(providerExclude []string) *SchedulerSearchAvailabilityParams {
	o.SetProviderExclude(providerExclude)
	return o
}

// SetProviderExclude adds the providerExclude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderExclude(providerExclude []string) {
	o.ProviderExclude = providerExclude
}

// WithProviderInPerson adds the providerInPerson to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderInPerson(providerInPerson *bool) *SchedulerSearchAvailabilityParams {
	o.SetProviderInPerson(providerInPerson)
	return o
}

// SetProviderInPerson adds the providerInPerson to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderInPerson(providerInPerson *bool) {
	o.ProviderInPerson = providerInPerson
}

// WithProviderInclude adds the providerInclude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderInclude(providerInclude []string) *SchedulerSearchAvailabilityParams {
	o.SetProviderInclude(providerInclude)
	return o
}

// SetProviderInclude adds the providerInclude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderInclude(providerInclude []string) {
	o.ProviderInclude = providerInclude
}

// WithProviderLimit adds the providerLimit to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderLimit(providerLimit *int32) *SchedulerSearchAvailabilityParams {
	o.SetProviderLimit(providerLimit)
	return o
}

// SetProviderLimit adds the providerLimit to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderLimit(providerLimit *int32) {
	o.ProviderLimit = providerLimit
}

// WithProviderLocationLatitude adds the providerLocationLatitude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderLocationLatitude(providerLocationLatitude *float32) *SchedulerSearchAvailabilityParams {
	o.SetProviderLocationLatitude(providerLocationLatitude)
	return o
}

// SetProviderLocationLatitude adds the providerLocationLatitude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderLocationLatitude(providerLocationLatitude *float32) {
	o.ProviderLocationLatitude = providerLocationLatitude
}

// WithProviderLocationLongitude adds the providerLocationLongitude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderLocationLongitude(providerLocationLongitude *float32) *SchedulerSearchAvailabilityParams {
	o.SetProviderLocationLongitude(providerLocationLongitude)
	return o
}

// SetProviderLocationLongitude adds the providerLocationLongitude to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderLocationLongitude(providerLocationLongitude *float32) {
	o.ProviderLocationLongitude = providerLocationLongitude
}

// WithProviderLocationQuery adds the providerLocationQuery to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderLocationQuery(providerLocationQuery *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderLocationQuery(providerLocationQuery)
	return o
}

// SetProviderLocationQuery adds the providerLocationQuery to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderLocationQuery(providerLocationQuery *string) {
	o.ProviderLocationQuery = providerLocationQuery
}

// WithProviderLocationZip adds the providerLocationZip to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderLocationZip(providerLocationZip *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderLocationZip(providerLocationZip)
	return o
}

// SetProviderLocationZip adds the providerLocationZip to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderLocationZip(providerLocationZip *string) {
	o.ProviderLocationZip = providerLocationZip
}

// WithProviderOffset adds the providerOffset to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderOffset(providerOffset *int32) *SchedulerSearchAvailabilityParams {
	o.SetProviderOffset(providerOffset)
	return o
}

// SetProviderOffset adds the providerOffset to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderOffset(providerOffset *int32) {
	o.ProviderOffset = providerOffset
}

// WithProviderRangeFrom adds the providerRangeFrom to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderRangeFrom(providerRangeFrom *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderRangeFrom(providerRangeFrom)
	return o
}

// SetProviderRangeFrom adds the providerRangeFrom to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderRangeFrom(providerRangeFrom *string) {
	o.ProviderRangeFrom = providerRangeFrom
}

// WithProviderRangeTo adds the providerRangeTo to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderRangeTo(providerRangeTo *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderRangeTo(providerRangeTo)
	return o
}

// SetProviderRangeTo adds the providerRangeTo to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderRangeTo(providerRangeTo *string) {
	o.ProviderRangeTo = providerRangeTo
}

// WithProviderServiceID adds the providerServiceID to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderServiceID(providerServiceID *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderServiceID(providerServiceID)
	return o
}

// SetProviderServiceID adds the providerServiceId to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderServiceID(providerServiceID *string) {
	o.ProviderServiceID = providerServiceID
}

// WithProviderState adds the providerState to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderState(providerState *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderState(providerState)
	return o
}

// SetProviderState adds the providerState to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderState(providerState *string) {
	o.ProviderState = providerState
}

// WithProviderZip adds the providerZip to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) WithProviderZip(providerZip *string) *SchedulerSearchAvailabilityParams {
	o.SetProviderZip(providerZip)
	return o
}

// SetProviderZip adds the providerZip to the scheduler search availability params
func (o *SchedulerSearchAvailabilityParams) SetProviderZip(providerZip *string) {
	o.ProviderZip = providerZip
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerSearchAvailabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Availability != nil {
		if err := r.SetBodyParam(o.Availability); err != nil {
			return err
		}
	}

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

	if o.ProviderDuration != nil {

		// query param provider.duration
		var qrProviderDuration int32

		if o.ProviderDuration != nil {
			qrProviderDuration = *o.ProviderDuration
		}
		qProviderDuration := swag.FormatInt32(qrProviderDuration)
		if qProviderDuration != "" {

			if err := r.SetQueryParam("provider.duration", qProviderDuration); err != nil {
				return err
			}
		}
	}

	if o.ProviderExclude != nil {

		// binding items for provider.exclude
		joinedProviderExclude := o.bindParamProviderExclude(reg)

		// query array param provider.exclude
		if err := r.SetQueryParam("provider.exclude", joinedProviderExclude...); err != nil {
			return err
		}
	}

	if o.ProviderInPerson != nil {

		// query param provider.in_person
		var qrProviderInPerson bool

		if o.ProviderInPerson != nil {
			qrProviderInPerson = *o.ProviderInPerson
		}
		qProviderInPerson := swag.FormatBool(qrProviderInPerson)
		if qProviderInPerson != "" {

			if err := r.SetQueryParam("provider.in_person", qProviderInPerson); err != nil {
				return err
			}
		}
	}

	if o.ProviderInclude != nil {

		// binding items for provider.include
		joinedProviderInclude := o.bindParamProviderInclude(reg)

		// query array param provider.include
		if err := r.SetQueryParam("provider.include", joinedProviderInclude...); err != nil {
			return err
		}
	}

	if o.ProviderLimit != nil {

		// query param provider.limit
		var qrProviderLimit int32

		if o.ProviderLimit != nil {
			qrProviderLimit = *o.ProviderLimit
		}
		qProviderLimit := swag.FormatInt32(qrProviderLimit)
		if qProviderLimit != "" {

			if err := r.SetQueryParam("provider.limit", qProviderLimit); err != nil {
				return err
			}
		}
	}

	if o.ProviderLocationLatitude != nil {

		// query param provider.location.latitude
		var qrProviderLocationLatitude float32

		if o.ProviderLocationLatitude != nil {
			qrProviderLocationLatitude = *o.ProviderLocationLatitude
		}
		qProviderLocationLatitude := swag.FormatFloat32(qrProviderLocationLatitude)
		if qProviderLocationLatitude != "" {

			if err := r.SetQueryParam("provider.location.latitude", qProviderLocationLatitude); err != nil {
				return err
			}
		}
	}

	if o.ProviderLocationLongitude != nil {

		// query param provider.location.longitude
		var qrProviderLocationLongitude float32

		if o.ProviderLocationLongitude != nil {
			qrProviderLocationLongitude = *o.ProviderLocationLongitude
		}
		qProviderLocationLongitude := swag.FormatFloat32(qrProviderLocationLongitude)
		if qProviderLocationLongitude != "" {

			if err := r.SetQueryParam("provider.location.longitude", qProviderLocationLongitude); err != nil {
				return err
			}
		}
	}

	if o.ProviderLocationQuery != nil {

		// query param provider.location.query
		var qrProviderLocationQuery string

		if o.ProviderLocationQuery != nil {
			qrProviderLocationQuery = *o.ProviderLocationQuery
		}
		qProviderLocationQuery := qrProviderLocationQuery
		if qProviderLocationQuery != "" {

			if err := r.SetQueryParam("provider.location.query", qProviderLocationQuery); err != nil {
				return err
			}
		}
	}

	if o.ProviderLocationZip != nil {

		// query param provider.location.zip
		var qrProviderLocationZip string

		if o.ProviderLocationZip != nil {
			qrProviderLocationZip = *o.ProviderLocationZip
		}
		qProviderLocationZip := qrProviderLocationZip
		if qProviderLocationZip != "" {

			if err := r.SetQueryParam("provider.location.zip", qProviderLocationZip); err != nil {
				return err
			}
		}
	}

	if o.ProviderOffset != nil {

		// query param provider.offset
		var qrProviderOffset int32

		if o.ProviderOffset != nil {
			qrProviderOffset = *o.ProviderOffset
		}
		qProviderOffset := swag.FormatInt32(qrProviderOffset)
		if qProviderOffset != "" {

			if err := r.SetQueryParam("provider.offset", qProviderOffset); err != nil {
				return err
			}
		}
	}

	if o.ProviderRangeFrom != nil {

		// query param provider.range_from
		var qrProviderRangeFrom string

		if o.ProviderRangeFrom != nil {
			qrProviderRangeFrom = *o.ProviderRangeFrom
		}
		qProviderRangeFrom := qrProviderRangeFrom
		if qProviderRangeFrom != "" {

			if err := r.SetQueryParam("provider.range_from", qProviderRangeFrom); err != nil {
				return err
			}
		}
	}

	if o.ProviderRangeTo != nil {

		// query param provider.range_to
		var qrProviderRangeTo string

		if o.ProviderRangeTo != nil {
			qrProviderRangeTo = *o.ProviderRangeTo
		}
		qProviderRangeTo := qrProviderRangeTo
		if qProviderRangeTo != "" {

			if err := r.SetQueryParam("provider.range_to", qProviderRangeTo); err != nil {
				return err
			}
		}
	}

	if o.ProviderServiceID != nil {

		// query param provider.service_id
		var qrProviderServiceID string

		if o.ProviderServiceID != nil {
			qrProviderServiceID = *o.ProviderServiceID
		}
		qProviderServiceID := qrProviderServiceID
		if qProviderServiceID != "" {

			if err := r.SetQueryParam("provider.service_id", qProviderServiceID); err != nil {
				return err
			}
		}
	}

	if o.ProviderState != nil {

		// query param provider.state
		var qrProviderState string

		if o.ProviderState != nil {
			qrProviderState = *o.ProviderState
		}
		qProviderState := qrProviderState
		if qProviderState != "" {

			if err := r.SetQueryParam("provider.state", qProviderState); err != nil {
				return err
			}
		}
	}

	if o.ProviderZip != nil {

		// query param provider.zip
		var qrProviderZip string

		if o.ProviderZip != nil {
			qrProviderZip = *o.ProviderZip
		}
		qProviderZip := qrProviderZip
		if qProviderZip != "" {

			if err := r.SetQueryParam("provider.zip", qProviderZip); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSchedulerSearchAvailability binds the parameter provider.exclude
func (o *SchedulerSearchAvailabilityParams) bindParamProviderExclude(formats strfmt.Registry) []string {
	providerExcludeIR := o.ProviderExclude

	var providerExcludeIC []string
	for _, providerExcludeIIR := range providerExcludeIR { // explode []string

		providerExcludeIIV := providerExcludeIIR // string as string
		providerExcludeIC = append(providerExcludeIC, providerExcludeIIV)
	}

	// items.CollectionFormat: "multi"
	providerExcludeIS := swag.JoinByFormat(providerExcludeIC, "multi")

	return providerExcludeIS
}

// bindParamSchedulerSearchAvailability binds the parameter provider.include
func (o *SchedulerSearchAvailabilityParams) bindParamProviderInclude(formats strfmt.Registry) []string {
	providerIncludeIR := o.ProviderInclude

	var providerIncludeIC []string
	for _, providerIncludeIIR := range providerIncludeIR { // explode []string

		providerIncludeIIV := providerIncludeIIR // string as string
		providerIncludeIC = append(providerIncludeIC, providerIncludeIIV)
	}

	// items.CollectionFormat: "multi"
	providerIncludeIS := swag.JoinByFormat(providerIncludeIC, "multi")

	return providerIncludeIS
}
