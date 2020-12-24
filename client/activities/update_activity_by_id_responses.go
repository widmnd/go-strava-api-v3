// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/strava-api-v3/models"
)

// UpdateActivityByIDReader is a Reader for the UpdateActivityByID structure.
type UpdateActivityByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateActivityByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateActivityByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateActivityByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateActivityByIDOK creates a UpdateActivityByIDOK with default headers values
func NewUpdateActivityByIDOK() *UpdateActivityByIDOK {
	return &UpdateActivityByIDOK{}
}

/*UpdateActivityByIDOK handles this case with default header values.

The activity's detailed representation.
*/
type UpdateActivityByIDOK struct {
	Payload *models.DetailedActivity
}

func (o *UpdateActivityByIDOK) Error() string {
	return fmt.Sprintf("[PUT /activities/{id}][%d] updateActivityByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateActivityByIDOK) GetPayload() *models.DetailedActivity {
	return o.Payload
}

func (o *UpdateActivityByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActivityByIDDefault creates a UpdateActivityByIDDefault with default headers values
func NewUpdateActivityByIDDefault(code int) *UpdateActivityByIDDefault {
	return &UpdateActivityByIDDefault{
		_statusCode: code,
	}
}

/*UpdateActivityByIDDefault handles this case with default header values.

Unexpected error.
*/
type UpdateActivityByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the update activity by Id default response
func (o *UpdateActivityByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateActivityByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /activities/{id}][%d] updateActivityById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateActivityByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *UpdateActivityByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
