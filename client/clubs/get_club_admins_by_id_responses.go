// Code generated by go-swagger; DO NOT EDIT.

package clubs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/go-strava-api-v3/models"
)

// GetClubAdminsByIDReader is a Reader for the GetClubAdminsByID structure.
type GetClubAdminsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClubAdminsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClubAdminsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClubAdminsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClubAdminsByIDOK creates a GetClubAdminsByIDOK with default headers values
func NewGetClubAdminsByIDOK() *GetClubAdminsByIDOK {
	return &GetClubAdminsByIDOK{}
}

/*GetClubAdminsByIDOK handles this case with default header values.

A list of summary athlete representations.
*/
type GetClubAdminsByIDOK struct {
	Payload []*models.SummaryAthlete
}

func (o *GetClubAdminsByIDOK) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/admins][%d] getClubAdminsByIdOK  %+v", 200, o.Payload)
}

func (o *GetClubAdminsByIDOK) GetPayload() []*models.SummaryAthlete {
	return o.Payload
}

func (o *GetClubAdminsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClubAdminsByIDDefault creates a GetClubAdminsByIDDefault with default headers values
func NewGetClubAdminsByIDDefault(code int) *GetClubAdminsByIDDefault {
	return &GetClubAdminsByIDDefault{
		_statusCode: code,
	}
}

/*GetClubAdminsByIDDefault handles this case with default header values.

Unexpected error.
*/
type GetClubAdminsByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get club admins by Id default response
func (o *GetClubAdminsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClubAdminsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/admins][%d] getClubAdminsById default  %+v", o._statusCode, o.Payload)
}

func (o *GetClubAdminsByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetClubAdminsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
