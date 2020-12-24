// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/go-strava-api-v3/models"
)

// GetZonesByActivityIDReader is a Reader for the GetZonesByActivityID structure.
type GetZonesByActivityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesByActivityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZonesByActivityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetZonesByActivityIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetZonesByActivityIDOK creates a GetZonesByActivityIDOK with default headers values
func NewGetZonesByActivityIDOK() *GetZonesByActivityIDOK {
	return &GetZonesByActivityIDOK{}
}

/*GetZonesByActivityIDOK handles this case with default header values.

Activity Zones.
*/
type GetZonesByActivityIDOK struct {
	Payload []*models.ActivityZone
}

func (o *GetZonesByActivityIDOK) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityIdOK  %+v", 200, o.Payload)
}

func (o *GetZonesByActivityIDOK) GetPayload() []*models.ActivityZone {
	return o.Payload
}

func (o *GetZonesByActivityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesByActivityIDDefault creates a GetZonesByActivityIDDefault with default headers values
func NewGetZonesByActivityIDDefault(code int) *GetZonesByActivityIDDefault {
	return &GetZonesByActivityIDDefault{
		_statusCode: code,
	}
}

/*GetZonesByActivityIDDefault handles this case with default header values.

Unexpected error.
*/
type GetZonesByActivityIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get zones by activity Id default response
func (o *GetZonesByActivityIDDefault) Code() int {
	return o._statusCode
}

func (o *GetZonesByActivityIDDefault) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityId default  %+v", o._statusCode, o.Payload)
}

func (o *GetZonesByActivityIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetZonesByActivityIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
