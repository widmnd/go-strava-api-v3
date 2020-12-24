// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/strava-api-v3/models"
)

// GetUploadByIDReader is a Reader for the GetUploadByID structure.
type GetUploadByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUploadByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUploadByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUploadByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUploadByIDOK creates a GetUploadByIDOK with default headers values
func NewGetUploadByIDOK() *GetUploadByIDOK {
	return &GetUploadByIDOK{}
}

/*GetUploadByIDOK handles this case with default header values.

Representation of the upload.
*/
type GetUploadByIDOK struct {
	Payload *models.Upload
}

func (o *GetUploadByIDOK) Error() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadByIdOK  %+v", 200, o.Payload)
}

func (o *GetUploadByIDOK) GetPayload() *models.Upload {
	return o.Payload
}

func (o *GetUploadByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Upload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadByIDDefault creates a GetUploadByIDDefault with default headers values
func NewGetUploadByIDDefault(code int) *GetUploadByIDDefault {
	return &GetUploadByIDDefault{
		_statusCode: code,
	}
}

/*GetUploadByIDDefault handles this case with default header values.

Unexpected error.
*/
type GetUploadByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get upload by Id default response
func (o *GetUploadByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetUploadByIDDefault) Error() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadById default  %+v", o._statusCode, o.Payload)
}

func (o *GetUploadByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetUploadByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
