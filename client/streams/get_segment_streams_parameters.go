// Code generated by go-swagger; DO NOT EDIT.

package streams

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

// NewGetSegmentStreamsParams creates a new GetSegmentStreamsParams object
// with the default values initialized.
func NewGetSegmentStreamsParams() *GetSegmentStreamsParams {
	var (
		keyByTypeDefault = bool(true)
	)
	return &GetSegmentStreamsParams{
		KeyByType: keyByTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentStreamsParamsWithTimeout creates a new GetSegmentStreamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSegmentStreamsParamsWithTimeout(timeout time.Duration) *GetSegmentStreamsParams {
	var (
		keyByTypeDefault = bool(true)
	)
	return &GetSegmentStreamsParams{
		KeyByType: keyByTypeDefault,

		timeout: timeout,
	}
}

// NewGetSegmentStreamsParamsWithContext creates a new GetSegmentStreamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSegmentStreamsParamsWithContext(ctx context.Context) *GetSegmentStreamsParams {
	var (
		keyByTypeDefault = bool(true)
	)
	return &GetSegmentStreamsParams{
		KeyByType: keyByTypeDefault,

		Context: ctx,
	}
}

// NewGetSegmentStreamsParamsWithHTTPClient creates a new GetSegmentStreamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSegmentStreamsParamsWithHTTPClient(client *http.Client) *GetSegmentStreamsParams {
	var (
		keyByTypeDefault = bool(true)
	)
	return &GetSegmentStreamsParams{
		KeyByType:  keyByTypeDefault,
		HTTPClient: client,
	}
}

/*GetSegmentStreamsParams contains all the parameters to send to the API endpoint
for the get segment streams operation typically these are written to a http.Request
*/
type GetSegmentStreamsParams struct {

	/*ID
	  The identifier of the segment.

	*/
	ID int64
	/*KeyByType
	  Must be true.

	*/
	KeyByType bool
	/*Keys
	  The types of streams to return.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get segment streams params
func (o *GetSegmentStreamsParams) WithTimeout(timeout time.Duration) *GetSegmentStreamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment streams params
func (o *GetSegmentStreamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment streams params
func (o *GetSegmentStreamsParams) WithContext(ctx context.Context) *GetSegmentStreamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment streams params
func (o *GetSegmentStreamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment streams params
func (o *GetSegmentStreamsParams) WithHTTPClient(client *http.Client) *GetSegmentStreamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment streams params
func (o *GetSegmentStreamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get segment streams params
func (o *GetSegmentStreamsParams) WithID(id int64) *GetSegmentStreamsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get segment streams params
func (o *GetSegmentStreamsParams) SetID(id int64) {
	o.ID = id
}

// WithKeyByType adds the keyByType to the get segment streams params
func (o *GetSegmentStreamsParams) WithKeyByType(keyByType bool) *GetSegmentStreamsParams {
	o.SetKeyByType(keyByType)
	return o
}

// SetKeyByType adds the keyByType to the get segment streams params
func (o *GetSegmentStreamsParams) SetKeyByType(keyByType bool) {
	o.KeyByType = keyByType
}

// WithKeys adds the keys to the get segment streams params
func (o *GetSegmentStreamsParams) WithKeys(keys []string) *GetSegmentStreamsParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the get segment streams params
func (o *GetSegmentStreamsParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentStreamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param key_by_type
	qrKeyByType := o.KeyByType
	qKeyByType := swag.FormatBool(qrKeyByType)
	if qKeyByType != "" {
		if err := r.SetQueryParam("key_by_type", qKeyByType); err != nil {
			return err
		}
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "csv")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
