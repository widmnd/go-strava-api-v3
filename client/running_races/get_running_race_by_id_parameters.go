// Code generated by go-swagger; DO NOT EDIT.

package running_races

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

// NewGetRunningRaceByIDParams creates a new GetRunningRaceByIDParams object
// with the default values initialized.
func NewGetRunningRaceByIDParams() *GetRunningRaceByIDParams {
	var ()
	return &GetRunningRaceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunningRaceByIDParamsWithTimeout creates a new GetRunningRaceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunningRaceByIDParamsWithTimeout(timeout time.Duration) *GetRunningRaceByIDParams {
	var ()
	return &GetRunningRaceByIDParams{

		timeout: timeout,
	}
}

// NewGetRunningRaceByIDParamsWithContext creates a new GetRunningRaceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunningRaceByIDParamsWithContext(ctx context.Context) *GetRunningRaceByIDParams {
	var ()
	return &GetRunningRaceByIDParams{

		Context: ctx,
	}
}

// NewGetRunningRaceByIDParamsWithHTTPClient creates a new GetRunningRaceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunningRaceByIDParamsWithHTTPClient(client *http.Client) *GetRunningRaceByIDParams {
	var ()
	return &GetRunningRaceByIDParams{
		HTTPClient: client,
	}
}

/*GetRunningRaceByIDParams contains all the parameters to send to the API endpoint
for the get running race by Id operation typically these are written to a http.Request
*/
type GetRunningRaceByIDParams struct {

	/*ID
	  The identifier of the running race.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get running race by Id params
func (o *GetRunningRaceByIDParams) WithTimeout(timeout time.Duration) *GetRunningRaceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get running race by Id params
func (o *GetRunningRaceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get running race by Id params
func (o *GetRunningRaceByIDParams) WithContext(ctx context.Context) *GetRunningRaceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get running race by Id params
func (o *GetRunningRaceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get running race by Id params
func (o *GetRunningRaceByIDParams) WithHTTPClient(client *http.Client) *GetRunningRaceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get running race by Id params
func (o *GetRunningRaceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get running race by Id params
func (o *GetRunningRaceByIDParams) WithID(id int64) *GetRunningRaceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get running race by Id params
func (o *GetRunningRaceByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunningRaceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}