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

// NewStarSegmentParams creates a new StarSegmentParams object
// with the default values initialized.
func NewStarSegmentParams() *StarSegmentParams {
	var (
		starredDefault = bool(false)
	)
	return &StarSegmentParams{
		Starred: starredDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStarSegmentParamsWithTimeout creates a new StarSegmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStarSegmentParamsWithTimeout(timeout time.Duration) *StarSegmentParams {
	var (
		starredDefault = bool(false)
	)
	return &StarSegmentParams{
		Starred: starredDefault,

		timeout: timeout,
	}
}

// NewStarSegmentParamsWithContext creates a new StarSegmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewStarSegmentParamsWithContext(ctx context.Context) *StarSegmentParams {
	var (
		starredDefault = bool(false)
	)
	return &StarSegmentParams{
		Starred: starredDefault,

		Context: ctx,
	}
}

// NewStarSegmentParamsWithHTTPClient creates a new StarSegmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStarSegmentParamsWithHTTPClient(client *http.Client) *StarSegmentParams {
	var (
		starredDefault = bool(false)
	)
	return &StarSegmentParams{
		Starred:    starredDefault,
		HTTPClient: client,
	}
}

/*StarSegmentParams contains all the parameters to send to the API endpoint
for the star segment operation typically these are written to a http.Request
*/
type StarSegmentParams struct {

	/*ID
	  The identifier of the segment to star.

	*/
	ID int64
	/*Starred
	  If true, star the segment; if false, unstar the segment.

	*/
	Starred bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the star segment params
func (o *StarSegmentParams) WithTimeout(timeout time.Duration) *StarSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the star segment params
func (o *StarSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the star segment params
func (o *StarSegmentParams) WithContext(ctx context.Context) *StarSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the star segment params
func (o *StarSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the star segment params
func (o *StarSegmentParams) WithHTTPClient(client *http.Client) *StarSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the star segment params
func (o *StarSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the star segment params
func (o *StarSegmentParams) WithID(id int64) *StarSegmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the star segment params
func (o *StarSegmentParams) SetID(id int64) {
	o.ID = id
}

// WithStarred adds the starred to the star segment params
func (o *StarSegmentParams) WithStarred(starred bool) *StarSegmentParams {
	o.SetStarred(starred)
	return o
}

// SetStarred adds the starred to the star segment params
func (o *StarSegmentParams) SetStarred(starred bool) {
	o.Starred = starred
}

// WriteToRequest writes these params to a swagger request
func (o *StarSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// form param starred
	frStarred := o.Starred
	fStarred := swag.FormatBool(frStarred)
	if fStarred != "" {
		if err := r.SetFormParam("starred", fStarred); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}