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

// NewSchedulerSearchProvidersParams creates a new SchedulerSearchProvidersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulerSearchProvidersParams() *SchedulerSearchProvidersParams {
	return &SchedulerSearchProvidersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulerSearchProvidersParamsWithTimeout creates a new SchedulerSearchProvidersParams object
// with the ability to set a timeout on a request.
func NewSchedulerSearchProvidersParamsWithTimeout(timeout time.Duration) *SchedulerSearchProvidersParams {
	return &SchedulerSearchProvidersParams{
		timeout: timeout,
	}
}

// NewSchedulerSearchProvidersParamsWithContext creates a new SchedulerSearchProvidersParams object
// with the ability to set a context for a request.
func NewSchedulerSearchProvidersParamsWithContext(ctx context.Context) *SchedulerSearchProvidersParams {
	return &SchedulerSearchProvidersParams{
		Context: ctx,
	}
}

// NewSchedulerSearchProvidersParamsWithHTTPClient creates a new SchedulerSearchProvidersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulerSearchProvidersParamsWithHTTPClient(client *http.Client) *SchedulerSearchProvidersParams {
	return &SchedulerSearchProvidersParams{
		HTTPClient: client,
	}
}

/* SchedulerSearchProvidersParams contains all the parameters to send to the API endpoint
   for the scheduler search providers operation.

   Typically these are written to a http.Request.
*/
type SchedulerSearchProvidersParams struct {

	/* AvailabilityDuration.

	   The duration of the appointment in minutes.

	   Format: int32
	*/
	AvailabilityDuration *int32

	/* AvailabilityExclude.

	     The ID of a specific providers to exclude those providers from search
	results.
	*/
	AvailabilityExclude []string

	// AvailabilityExcludeByRef.
	AvailabilityExcludeByRef *bool

	/* AvailabilityInPerson.

	   Return only in-person providers (default: false).
	*/
	AvailabilityInPerson *bool

	/* AvailabilityInclude.

	     The IDs of specific providers to retrieve appointments for just those
	providers.
	*/
	AvailabilityInclude []string

	// AvailabilityIncludeByRef.
	AvailabilityIncludeByRef *bool

	// AvailabilityLanguage.
	AvailabilityLanguage *string

	/* AvailabilityLimit.

	   Limit the number of results to return. (default: 10).

	   Format: int32
	*/
	AvailabilityLimit *int32

	/* AvailabilityLocationLatitude.

	   Specific latitude point to use.

	   Format: float
	*/
	AvailabilityLocationLatitude *float32

	/* AvailabilityLocationLongitude.

	   Specific longitude point to use.

	   Format: float
	*/
	AvailabilityLocationLongitude *float32

	/* AvailabilityLocationQuery.

	   Address query (ex: 123 Main St., San Diego, CA, 92101, San Diego, CA, etc.).
	*/
	AvailabilityLocationQuery *string

	/* AvailabilityLocationZip.

	   Zip code.
	*/
	AvailabilityLocationZip *string

	/* AvailabilityOffset.

	     Offset the results count if pagination is necessary based on the limit and
	total fields in the response.

	     Format: int32
	*/
	AvailabilityOffset *int32

	/* AvailabilityRangeFrom.

	   The start of the requested date/time range in RFC3339 format.
	*/
	AvailabilityRangeFrom *string

	/* AvailabilityRangeTo.

	   The end of the requested date/time range in RFC3339 format.
	*/
	AvailabilityRangeTo *string

	/* AvailabilityServiceID.

	   The service UUID requested by the patient.
	*/
	AvailabilityServiceID *string

	/* AvailabilityState.

	   State the provider is licensed in for the requested service.
	*/
	AvailabilityState *string

	// AvailabilityTimezoneCode.
	//
	// Format: int32
	AvailabilityTimezoneCode *int32

	// AvailabilityTimezoneName.
	//
	// Format: int32
	AvailabilityTimezoneName *int32

	// AvailabilityTimezoneOffset.
	//
	// Format: int32
	AvailabilityTimezoneOffset *int32

	/* AvailabilityZip.

	   ZIP code the patient resides in.
	*/
	AvailabilityZip *string

	// Body.
	Body *models.SchedulerProviderSearchRequest

	// OrganizationID.
	OrganizationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduler search providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerSearchProvidersParams) WithDefaults() *SchedulerSearchProvidersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduler search providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulerSearchProvidersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithTimeout(timeout time.Duration) *SchedulerSearchProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithContext(ctx context.Context) *SchedulerSearchProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithHTTPClient(client *http.Client) *SchedulerSearchProvidersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityDuration adds the availabilityDuration to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityDuration(availabilityDuration *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityDuration(availabilityDuration)
	return o
}

// SetAvailabilityDuration adds the availabilityDuration to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityDuration(availabilityDuration *int32) {
	o.AvailabilityDuration = availabilityDuration
}

// WithAvailabilityExclude adds the availabilityExclude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityExclude(availabilityExclude []string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityExclude(availabilityExclude)
	return o
}

// SetAvailabilityExclude adds the availabilityExclude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityExclude(availabilityExclude []string) {
	o.AvailabilityExclude = availabilityExclude
}

// WithAvailabilityExcludeByRef adds the availabilityExcludeByRef to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityExcludeByRef(availabilityExcludeByRef *bool) *SchedulerSearchProvidersParams {
	o.SetAvailabilityExcludeByRef(availabilityExcludeByRef)
	return o
}

// SetAvailabilityExcludeByRef adds the availabilityExcludeByRef to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityExcludeByRef(availabilityExcludeByRef *bool) {
	o.AvailabilityExcludeByRef = availabilityExcludeByRef
}

// WithAvailabilityInPerson adds the availabilityInPerson to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityInPerson(availabilityInPerson *bool) *SchedulerSearchProvidersParams {
	o.SetAvailabilityInPerson(availabilityInPerson)
	return o
}

// SetAvailabilityInPerson adds the availabilityInPerson to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityInPerson(availabilityInPerson *bool) {
	o.AvailabilityInPerson = availabilityInPerson
}

// WithAvailabilityInclude adds the availabilityInclude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityInclude(availabilityInclude []string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityInclude(availabilityInclude)
	return o
}

// SetAvailabilityInclude adds the availabilityInclude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityInclude(availabilityInclude []string) {
	o.AvailabilityInclude = availabilityInclude
}

// WithAvailabilityIncludeByRef adds the availabilityIncludeByRef to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityIncludeByRef(availabilityIncludeByRef *bool) *SchedulerSearchProvidersParams {
	o.SetAvailabilityIncludeByRef(availabilityIncludeByRef)
	return o
}

// SetAvailabilityIncludeByRef adds the availabilityIncludeByRef to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityIncludeByRef(availabilityIncludeByRef *bool) {
	o.AvailabilityIncludeByRef = availabilityIncludeByRef
}

// WithAvailabilityLanguage adds the availabilityLanguage to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLanguage(availabilityLanguage *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLanguage(availabilityLanguage)
	return o
}

// SetAvailabilityLanguage adds the availabilityLanguage to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLanguage(availabilityLanguage *string) {
	o.AvailabilityLanguage = availabilityLanguage
}

// WithAvailabilityLimit adds the availabilityLimit to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLimit(availabilityLimit *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLimit(availabilityLimit)
	return o
}

// SetAvailabilityLimit adds the availabilityLimit to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLimit(availabilityLimit *int32) {
	o.AvailabilityLimit = availabilityLimit
}

// WithAvailabilityLocationLatitude adds the availabilityLocationLatitude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLocationLatitude(availabilityLocationLatitude *float32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLocationLatitude(availabilityLocationLatitude)
	return o
}

// SetAvailabilityLocationLatitude adds the availabilityLocationLatitude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLocationLatitude(availabilityLocationLatitude *float32) {
	o.AvailabilityLocationLatitude = availabilityLocationLatitude
}

// WithAvailabilityLocationLongitude adds the availabilityLocationLongitude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLocationLongitude(availabilityLocationLongitude *float32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLocationLongitude(availabilityLocationLongitude)
	return o
}

// SetAvailabilityLocationLongitude adds the availabilityLocationLongitude to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLocationLongitude(availabilityLocationLongitude *float32) {
	o.AvailabilityLocationLongitude = availabilityLocationLongitude
}

// WithAvailabilityLocationQuery adds the availabilityLocationQuery to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLocationQuery(availabilityLocationQuery *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLocationQuery(availabilityLocationQuery)
	return o
}

// SetAvailabilityLocationQuery adds the availabilityLocationQuery to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLocationQuery(availabilityLocationQuery *string) {
	o.AvailabilityLocationQuery = availabilityLocationQuery
}

// WithAvailabilityLocationZip adds the availabilityLocationZip to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityLocationZip(availabilityLocationZip *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityLocationZip(availabilityLocationZip)
	return o
}

// SetAvailabilityLocationZip adds the availabilityLocationZip to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityLocationZip(availabilityLocationZip *string) {
	o.AvailabilityLocationZip = availabilityLocationZip
}

// WithAvailabilityOffset adds the availabilityOffset to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityOffset(availabilityOffset *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityOffset(availabilityOffset)
	return o
}

// SetAvailabilityOffset adds the availabilityOffset to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityOffset(availabilityOffset *int32) {
	o.AvailabilityOffset = availabilityOffset
}

// WithAvailabilityRangeFrom adds the availabilityRangeFrom to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityRangeFrom(availabilityRangeFrom *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityRangeFrom(availabilityRangeFrom)
	return o
}

// SetAvailabilityRangeFrom adds the availabilityRangeFrom to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityRangeFrom(availabilityRangeFrom *string) {
	o.AvailabilityRangeFrom = availabilityRangeFrom
}

// WithAvailabilityRangeTo adds the availabilityRangeTo to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityRangeTo(availabilityRangeTo *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityRangeTo(availabilityRangeTo)
	return o
}

// SetAvailabilityRangeTo adds the availabilityRangeTo to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityRangeTo(availabilityRangeTo *string) {
	o.AvailabilityRangeTo = availabilityRangeTo
}

// WithAvailabilityServiceID adds the availabilityServiceID to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityServiceID(availabilityServiceID *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityServiceID(availabilityServiceID)
	return o
}

// SetAvailabilityServiceID adds the availabilityServiceId to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityServiceID(availabilityServiceID *string) {
	o.AvailabilityServiceID = availabilityServiceID
}

// WithAvailabilityState adds the availabilityState to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityState(availabilityState *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityState(availabilityState)
	return o
}

// SetAvailabilityState adds the availabilityState to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityState(availabilityState *string) {
	o.AvailabilityState = availabilityState
}

// WithAvailabilityTimezoneCode adds the availabilityTimezoneCode to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityTimezoneCode(availabilityTimezoneCode *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityTimezoneCode(availabilityTimezoneCode)
	return o
}

// SetAvailabilityTimezoneCode adds the availabilityTimezoneCode to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityTimezoneCode(availabilityTimezoneCode *int32) {
	o.AvailabilityTimezoneCode = availabilityTimezoneCode
}

// WithAvailabilityTimezoneName adds the availabilityTimezoneName to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityTimezoneName(availabilityTimezoneName *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityTimezoneName(availabilityTimezoneName)
	return o
}

// SetAvailabilityTimezoneName adds the availabilityTimezoneName to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityTimezoneName(availabilityTimezoneName *int32) {
	o.AvailabilityTimezoneName = availabilityTimezoneName
}

// WithAvailabilityTimezoneOffset adds the availabilityTimezoneOffset to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityTimezoneOffset(availabilityTimezoneOffset *int32) *SchedulerSearchProvidersParams {
	o.SetAvailabilityTimezoneOffset(availabilityTimezoneOffset)
	return o
}

// SetAvailabilityTimezoneOffset adds the availabilityTimezoneOffset to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityTimezoneOffset(availabilityTimezoneOffset *int32) {
	o.AvailabilityTimezoneOffset = availabilityTimezoneOffset
}

// WithAvailabilityZip adds the availabilityZip to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithAvailabilityZip(availabilityZip *string) *SchedulerSearchProvidersParams {
	o.SetAvailabilityZip(availabilityZip)
	return o
}

// SetAvailabilityZip adds the availabilityZip to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetAvailabilityZip(availabilityZip *string) {
	o.AvailabilityZip = availabilityZip
}

// WithBody adds the body to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithBody(body *models.SchedulerProviderSearchRequest) *SchedulerSearchProvidersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetBody(body *models.SchedulerProviderSearchRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) WithOrganizationID(organizationID *string) *SchedulerSearchProvidersParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the scheduler search providers params
func (o *SchedulerSearchProvidersParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulerSearchProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityDuration != nil {

		// query param availability.duration
		var qrAvailabilityDuration int32

		if o.AvailabilityDuration != nil {
			qrAvailabilityDuration = *o.AvailabilityDuration
		}
		qAvailabilityDuration := swag.FormatInt32(qrAvailabilityDuration)
		if qAvailabilityDuration != "" {

			if err := r.SetQueryParam("availability.duration", qAvailabilityDuration); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityExclude != nil {

		// binding items for availability.exclude
		joinedAvailabilityExclude := o.bindParamAvailabilityExclude(reg)

		// query array param availability.exclude
		if err := r.SetQueryParam("availability.exclude", joinedAvailabilityExclude...); err != nil {
			return err
		}
	}

	if o.AvailabilityExcludeByRef != nil {

		// query param availability.exclude_by_ref
		var qrAvailabilityExcludeByRef bool

		if o.AvailabilityExcludeByRef != nil {
			qrAvailabilityExcludeByRef = *o.AvailabilityExcludeByRef
		}
		qAvailabilityExcludeByRef := swag.FormatBool(qrAvailabilityExcludeByRef)
		if qAvailabilityExcludeByRef != "" {

			if err := r.SetQueryParam("availability.exclude_by_ref", qAvailabilityExcludeByRef); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityInPerson != nil {

		// query param availability.in_person
		var qrAvailabilityInPerson bool

		if o.AvailabilityInPerson != nil {
			qrAvailabilityInPerson = *o.AvailabilityInPerson
		}
		qAvailabilityInPerson := swag.FormatBool(qrAvailabilityInPerson)
		if qAvailabilityInPerson != "" {

			if err := r.SetQueryParam("availability.in_person", qAvailabilityInPerson); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityInclude != nil {

		// binding items for availability.include
		joinedAvailabilityInclude := o.bindParamAvailabilityInclude(reg)

		// query array param availability.include
		if err := r.SetQueryParam("availability.include", joinedAvailabilityInclude...); err != nil {
			return err
		}
	}

	if o.AvailabilityIncludeByRef != nil {

		// query param availability.include_by_ref
		var qrAvailabilityIncludeByRef bool

		if o.AvailabilityIncludeByRef != nil {
			qrAvailabilityIncludeByRef = *o.AvailabilityIncludeByRef
		}
		qAvailabilityIncludeByRef := swag.FormatBool(qrAvailabilityIncludeByRef)
		if qAvailabilityIncludeByRef != "" {

			if err := r.SetQueryParam("availability.include_by_ref", qAvailabilityIncludeByRef); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLanguage != nil {

		// query param availability.language
		var qrAvailabilityLanguage string

		if o.AvailabilityLanguage != nil {
			qrAvailabilityLanguage = *o.AvailabilityLanguage
		}
		qAvailabilityLanguage := qrAvailabilityLanguage
		if qAvailabilityLanguage != "" {

			if err := r.SetQueryParam("availability.language", qAvailabilityLanguage); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLimit != nil {

		// query param availability.limit
		var qrAvailabilityLimit int32

		if o.AvailabilityLimit != nil {
			qrAvailabilityLimit = *o.AvailabilityLimit
		}
		qAvailabilityLimit := swag.FormatInt32(qrAvailabilityLimit)
		if qAvailabilityLimit != "" {

			if err := r.SetQueryParam("availability.limit", qAvailabilityLimit); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLocationLatitude != nil {

		// query param availability.location.latitude
		var qrAvailabilityLocationLatitude float32

		if o.AvailabilityLocationLatitude != nil {
			qrAvailabilityLocationLatitude = *o.AvailabilityLocationLatitude
		}
		qAvailabilityLocationLatitude := swag.FormatFloat32(qrAvailabilityLocationLatitude)
		if qAvailabilityLocationLatitude != "" {

			if err := r.SetQueryParam("availability.location.latitude", qAvailabilityLocationLatitude); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLocationLongitude != nil {

		// query param availability.location.longitude
		var qrAvailabilityLocationLongitude float32

		if o.AvailabilityLocationLongitude != nil {
			qrAvailabilityLocationLongitude = *o.AvailabilityLocationLongitude
		}
		qAvailabilityLocationLongitude := swag.FormatFloat32(qrAvailabilityLocationLongitude)
		if qAvailabilityLocationLongitude != "" {

			if err := r.SetQueryParam("availability.location.longitude", qAvailabilityLocationLongitude); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLocationQuery != nil {

		// query param availability.location.query
		var qrAvailabilityLocationQuery string

		if o.AvailabilityLocationQuery != nil {
			qrAvailabilityLocationQuery = *o.AvailabilityLocationQuery
		}
		qAvailabilityLocationQuery := qrAvailabilityLocationQuery
		if qAvailabilityLocationQuery != "" {

			if err := r.SetQueryParam("availability.location.query", qAvailabilityLocationQuery); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityLocationZip != nil {

		// query param availability.location.zip
		var qrAvailabilityLocationZip string

		if o.AvailabilityLocationZip != nil {
			qrAvailabilityLocationZip = *o.AvailabilityLocationZip
		}
		qAvailabilityLocationZip := qrAvailabilityLocationZip
		if qAvailabilityLocationZip != "" {

			if err := r.SetQueryParam("availability.location.zip", qAvailabilityLocationZip); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityOffset != nil {

		// query param availability.offset
		var qrAvailabilityOffset int32

		if o.AvailabilityOffset != nil {
			qrAvailabilityOffset = *o.AvailabilityOffset
		}
		qAvailabilityOffset := swag.FormatInt32(qrAvailabilityOffset)
		if qAvailabilityOffset != "" {

			if err := r.SetQueryParam("availability.offset", qAvailabilityOffset); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityRangeFrom != nil {

		// query param availability.range_from
		var qrAvailabilityRangeFrom string

		if o.AvailabilityRangeFrom != nil {
			qrAvailabilityRangeFrom = *o.AvailabilityRangeFrom
		}
		qAvailabilityRangeFrom := qrAvailabilityRangeFrom
		if qAvailabilityRangeFrom != "" {

			if err := r.SetQueryParam("availability.range_from", qAvailabilityRangeFrom); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityRangeTo != nil {

		// query param availability.range_to
		var qrAvailabilityRangeTo string

		if o.AvailabilityRangeTo != nil {
			qrAvailabilityRangeTo = *o.AvailabilityRangeTo
		}
		qAvailabilityRangeTo := qrAvailabilityRangeTo
		if qAvailabilityRangeTo != "" {

			if err := r.SetQueryParam("availability.range_to", qAvailabilityRangeTo); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityServiceID != nil {

		// query param availability.service_id
		var qrAvailabilityServiceID string

		if o.AvailabilityServiceID != nil {
			qrAvailabilityServiceID = *o.AvailabilityServiceID
		}
		qAvailabilityServiceID := qrAvailabilityServiceID
		if qAvailabilityServiceID != "" {

			if err := r.SetQueryParam("availability.service_id", qAvailabilityServiceID); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityState != nil {

		// query param availability.state
		var qrAvailabilityState string

		if o.AvailabilityState != nil {
			qrAvailabilityState = *o.AvailabilityState
		}
		qAvailabilityState := qrAvailabilityState
		if qAvailabilityState != "" {

			if err := r.SetQueryParam("availability.state", qAvailabilityState); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityTimezoneCode != nil {

		// query param availability.timezone_code
		var qrAvailabilityTimezoneCode int32

		if o.AvailabilityTimezoneCode != nil {
			qrAvailabilityTimezoneCode = *o.AvailabilityTimezoneCode
		}
		qAvailabilityTimezoneCode := swag.FormatInt32(qrAvailabilityTimezoneCode)
		if qAvailabilityTimezoneCode != "" {

			if err := r.SetQueryParam("availability.timezone_code", qAvailabilityTimezoneCode); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityTimezoneName != nil {

		// query param availability.timezone_name
		var qrAvailabilityTimezoneName int32

		if o.AvailabilityTimezoneName != nil {
			qrAvailabilityTimezoneName = *o.AvailabilityTimezoneName
		}
		qAvailabilityTimezoneName := swag.FormatInt32(qrAvailabilityTimezoneName)
		if qAvailabilityTimezoneName != "" {

			if err := r.SetQueryParam("availability.timezone_name", qAvailabilityTimezoneName); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityTimezoneOffset != nil {

		// query param availability.timezone_offset
		var qrAvailabilityTimezoneOffset int32

		if o.AvailabilityTimezoneOffset != nil {
			qrAvailabilityTimezoneOffset = *o.AvailabilityTimezoneOffset
		}
		qAvailabilityTimezoneOffset := swag.FormatInt32(qrAvailabilityTimezoneOffset)
		if qAvailabilityTimezoneOffset != "" {

			if err := r.SetQueryParam("availability.timezone_offset", qAvailabilityTimezoneOffset); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityZip != nil {

		// query param availability.zip
		var qrAvailabilityZip string

		if o.AvailabilityZip != nil {
			qrAvailabilityZip = *o.AvailabilityZip
		}
		qAvailabilityZip := qrAvailabilityZip
		if qAvailabilityZip != "" {

			if err := r.SetQueryParam("availability.zip", qAvailabilityZip); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSchedulerSearchProviders binds the parameter availability.exclude
func (o *SchedulerSearchProvidersParams) bindParamAvailabilityExclude(formats strfmt.Registry) []string {
	availabilityExcludeIR := o.AvailabilityExclude

	var availabilityExcludeIC []string
	for _, availabilityExcludeIIR := range availabilityExcludeIR { // explode []string

		availabilityExcludeIIV := availabilityExcludeIIR // string as string
		availabilityExcludeIC = append(availabilityExcludeIC, availabilityExcludeIIV)
	}

	// items.CollectionFormat: "multi"
	availabilityExcludeIS := swag.JoinByFormat(availabilityExcludeIC, "multi")

	return availabilityExcludeIS
}

// bindParamSchedulerSearchProviders binds the parameter availability.include
func (o *SchedulerSearchProvidersParams) bindParamAvailabilityInclude(formats strfmt.Registry) []string {
	availabilityIncludeIR := o.AvailabilityInclude

	var availabilityIncludeIC []string
	for _, availabilityIncludeIIR := range availabilityIncludeIR { // explode []string

		availabilityIncludeIIV := availabilityIncludeIIR // string as string
		availabilityIncludeIC = append(availabilityIncludeIC, availabilityIncludeIIV)
	}

	// items.CollectionFormat: "multi"
	availabilityIncludeIS := swag.JoinByFormat(availabilityIncludeIC, "multi")

	return availabilityIncludeIS
}