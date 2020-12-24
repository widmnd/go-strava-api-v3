// Code generated by go-swagger; DO NOT EDIT.

package clubs

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

// NewGetClubByIDParams creates a new GetClubByIDParams object
// with the default values initialized.
func NewGetClubByIDParams() *GetClubByIDParams {
	var ()
	return &GetClubByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClubByIDParamsWithTimeout creates a new GetClubByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClubByIDParamsWithTimeout(timeout time.Duration) *GetClubByIDParams {
	var ()
	return &GetClubByIDParams{

		timeout: timeout,
	}
}

// NewGetClubByIDParamsWithContext creates a new GetClubByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClubByIDParamsWithContext(ctx context.Context) *GetClubByIDParams {
	var ()
	return &GetClubByIDParams{

		Context: ctx,
	}
}

// NewGetClubByIDParamsWithHTTPClient creates a new GetClubByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClubByIDParamsWithHTTPClient(client *http.Client) *GetClubByIDParams {
	var ()
	return &GetClubByIDParams{
		HTTPClient: client,
	}
}

/*GetClubByIDParams contains all the parameters to send to the API endpoint
for the get club by Id operation typically these are written to a http.Request
*/
type GetClubByIDParams struct {

	/*ID
	  The identifier of the club.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get club by Id params
func (o *GetClubByIDParams) WithTimeout(timeout time.Duration) *GetClubByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get club by Id params
func (o *GetClubByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get club by Id params
func (o *GetClubByIDParams) WithContext(ctx context.Context) *GetClubByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get club by Id params
func (o *GetClubByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get club by Id params
func (o *GetClubByIDParams) WithHTTPClient(client *http.Client) *GetClubByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get club by Id params
func (o *GetClubByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get club by Id params
func (o *GetClubByIDParams) WithID(id int64) *GetClubByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get club by Id params
func (o *GetClubByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClubByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
