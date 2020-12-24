// Code generated by go-swagger; DO NOT EDIT.

package athletes

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

// NewGetLoggedInAthleteZonesParams creates a new GetLoggedInAthleteZonesParams object
// with the default values initialized.
func NewGetLoggedInAthleteZonesParams() *GetLoggedInAthleteZonesParams {

	return &GetLoggedInAthleteZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedInAthleteZonesParamsWithTimeout creates a new GetLoggedInAthleteZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoggedInAthleteZonesParamsWithTimeout(timeout time.Duration) *GetLoggedInAthleteZonesParams {

	return &GetLoggedInAthleteZonesParams{

		timeout: timeout,
	}
}

// NewGetLoggedInAthleteZonesParamsWithContext creates a new GetLoggedInAthleteZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoggedInAthleteZonesParamsWithContext(ctx context.Context) *GetLoggedInAthleteZonesParams {

	return &GetLoggedInAthleteZonesParams{

		Context: ctx,
	}
}

// NewGetLoggedInAthleteZonesParamsWithHTTPClient creates a new GetLoggedInAthleteZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoggedInAthleteZonesParamsWithHTTPClient(client *http.Client) *GetLoggedInAthleteZonesParams {

	return &GetLoggedInAthleteZonesParams{
		HTTPClient: client,
	}
}

/*GetLoggedInAthleteZonesParams contains all the parameters to send to the API endpoint
for the get logged in athlete zones operation typically these are written to a http.Request
*/
type GetLoggedInAthleteZonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) WithTimeout(timeout time.Duration) *GetLoggedInAthleteZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) WithContext(ctx context.Context) *GetLoggedInAthleteZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) WithHTTPClient(client *http.Client) *GetLoggedInAthleteZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logged in athlete zones params
func (o *GetLoggedInAthleteZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInAthleteZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
