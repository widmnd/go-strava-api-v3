// Code generated by go-swagger; DO NOT EDIT.

package segments

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

// NewExploreSegmentsParams creates a new ExploreSegmentsParams object
// with the default values initialized.
func NewExploreSegmentsParams() *ExploreSegmentsParams {
	var ()
	return &ExploreSegmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExploreSegmentsParamsWithTimeout creates a new ExploreSegmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExploreSegmentsParamsWithTimeout(timeout time.Duration) *ExploreSegmentsParams {
	var ()
	return &ExploreSegmentsParams{

		timeout: timeout,
	}
}

// NewExploreSegmentsParamsWithContext creates a new ExploreSegmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewExploreSegmentsParamsWithContext(ctx context.Context) *ExploreSegmentsParams {
	var ()
	return &ExploreSegmentsParams{

		Context: ctx,
	}
}

// NewExploreSegmentsParamsWithHTTPClient creates a new ExploreSegmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExploreSegmentsParamsWithHTTPClient(client *http.Client) *ExploreSegmentsParams {
	var ()
	return &ExploreSegmentsParams{
		HTTPClient: client,
	}
}

/*ExploreSegmentsParams contains all the parameters to send to the API endpoint
for the explore segments operation typically these are written to a http.Request
*/
type ExploreSegmentsParams struct {

	/*ActivityType
	  Desired activity type.

	*/
	ActivityType *string
	/*Bounds
	  The latitude and longitude for two points describing a rectangular boundary for the search: [southwest corner latitutde, southwest corner longitude, northeast corner latitude, northeast corner longitude]

	*/
	Bounds []float32
	/*MaxCat
	  The maximum climbing category.

	*/
	MaxCat *int64
	/*MinCat
	  The minimum climbing category.

	*/
	MinCat *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the explore segments params
func (o *ExploreSegmentsParams) WithTimeout(timeout time.Duration) *ExploreSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the explore segments params
func (o *ExploreSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the explore segments params
func (o *ExploreSegmentsParams) WithContext(ctx context.Context) *ExploreSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the explore segments params
func (o *ExploreSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the explore segments params
func (o *ExploreSegmentsParams) WithHTTPClient(client *http.Client) *ExploreSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the explore segments params
func (o *ExploreSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityType adds the activityType to the explore segments params
func (o *ExploreSegmentsParams) WithActivityType(activityType *string) *ExploreSegmentsParams {
	o.SetActivityType(activityType)
	return o
}

// SetActivityType adds the activityType to the explore segments params
func (o *ExploreSegmentsParams) SetActivityType(activityType *string) {
	o.ActivityType = activityType
}

// WithBounds adds the bounds to the explore segments params
func (o *ExploreSegmentsParams) WithBounds(bounds []float32) *ExploreSegmentsParams {
	o.SetBounds(bounds)
	return o
}

// SetBounds adds the bounds to the explore segments params
func (o *ExploreSegmentsParams) SetBounds(bounds []float32) {
	o.Bounds = bounds
}

// WithMaxCat adds the maxCat to the explore segments params
func (o *ExploreSegmentsParams) WithMaxCat(maxCat *int64) *ExploreSegmentsParams {
	o.SetMaxCat(maxCat)
	return o
}

// SetMaxCat adds the maxCat to the explore segments params
func (o *ExploreSegmentsParams) SetMaxCat(maxCat *int64) {
	o.MaxCat = maxCat
}

// WithMinCat adds the minCat to the explore segments params
func (o *ExploreSegmentsParams) WithMinCat(minCat *int64) *ExploreSegmentsParams {
	o.SetMinCat(minCat)
	return o
}

// SetMinCat adds the minCat to the explore segments params
func (o *ExploreSegmentsParams) SetMinCat(minCat *int64) {
	o.MinCat = minCat
}

// WriteToRequest writes these params to a swagger request
func (o *ExploreSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActivityType != nil {

		// query param activity_type
		var qrActivityType string
		if o.ActivityType != nil {
			qrActivityType = *o.ActivityType
		}
		qActivityType := qrActivityType
		if qActivityType != "" {
			if err := r.SetQueryParam("activity_type", qActivityType); err != nil {
				return err
			}
		}

	}

	var valuesBounds []string
	for _, v := range o.Bounds {
		valuesBounds = append(valuesBounds, swag.FormatFloat32(v))
	}

	joinedBounds := swag.JoinByFormat(valuesBounds, "csv")
	// query array param bounds
	if err := r.SetQueryParam("bounds", joinedBounds...); err != nil {
		return err
	}

	if o.MaxCat != nil {

		// query param max_cat
		var qrMaxCat int64
		if o.MaxCat != nil {
			qrMaxCat = *o.MaxCat
		}
		qMaxCat := swag.FormatInt64(qrMaxCat)
		if qMaxCat != "" {
			if err := r.SetQueryParam("max_cat", qMaxCat); err != nil {
				return err
			}
		}

	}

	if o.MinCat != nil {

		// query param min_cat
		var qrMinCat int64
		if o.MinCat != nil {
			qrMinCat = *o.MinCat
		}
		qMinCat := swag.FormatInt64(qrMinCat)
		if qMinCat != "" {
			if err := r.SetQueryParam("min_cat", qMinCat); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
