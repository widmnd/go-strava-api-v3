// Code generated by go-swagger; DO NOT EDIT.

package athletes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new athletes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for athletes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetLoggedInAthlete(params *GetLoggedInAthleteParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoggedInAthleteOK, error)

	GetLoggedInAthleteZones(params *GetLoggedInAthleteZonesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoggedInAthleteZonesOK, error)

	GetStats(params *GetStatsParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatsOK, error)

	UpdateLoggedInAthlete(params *UpdateLoggedInAthleteParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLoggedInAthleteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLoggedInAthlete gets authenticated athlete

  Returns the currently authenticated athlete. Tokens with profile:read_all scope will receive a detailed athlete representation; all others will receive a summary representation.
*/
func (a *Client) GetLoggedInAthlete(params *GetLoggedInAthleteParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoggedInAthleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoggedInAthleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoggedInAthlete",
		Method:             "GET",
		PathPattern:        "/athlete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoggedInAthleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLoggedInAthleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLoggedInAthleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLoggedInAthleteZones gets zones

  Returns the the authenticated athlete's heart rate and power zones. Requires profile:read_all.
*/
func (a *Client) GetLoggedInAthleteZones(params *GetLoggedInAthleteZonesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoggedInAthleteZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoggedInAthleteZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoggedInAthleteZones",
		Method:             "GET",
		PathPattern:        "/athlete/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoggedInAthleteZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLoggedInAthleteZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLoggedInAthleteZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetStats gets athlete stats

  Returns the activity stats of an athlete. Only includes data from activities set to Everyone visibilty.
*/
func (a *Client) GetStats(params *GetStatsParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStats",
		Method:             "GET",
		PathPattern:        "/athletes/{id}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateLoggedInAthlete updates athlete

  Update the currently authenticated athlete. Requires profile:write scope.
*/
func (a *Client) UpdateLoggedInAthlete(params *UpdateLoggedInAthleteParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLoggedInAthleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLoggedInAthleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLoggedInAthlete",
		Method:             "PUT",
		PathPattern:        "/athlete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLoggedInAthleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateLoggedInAthleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateLoggedInAthleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}