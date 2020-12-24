// Code generated by go-swagger; DO NOT EDIT.

package segment_efforts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/go-strava-api-v3/models"
)

// GetEffortsBySegmentIDReader is a Reader for the GetEffortsBySegmentID structure.
type GetEffortsBySegmentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEffortsBySegmentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEffortsBySegmentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEffortsBySegmentIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEffortsBySegmentIDOK creates a GetEffortsBySegmentIDOK with default headers values
func NewGetEffortsBySegmentIDOK() *GetEffortsBySegmentIDOK {
	return &GetEffortsBySegmentIDOK{}
}

/*GetEffortsBySegmentIDOK handles this case with default header values.

List of segment efforts.
*/
type GetEffortsBySegmentIDOK struct {
	Payload []*models.DetailedSegmentEffort
}

func (o *GetEffortsBySegmentIDOK) Error() string {
	return fmt.Sprintf("[GET /segment_efforts][%d] getEffortsBySegmentIdOK  %+v", 200, o.Payload)
}

func (o *GetEffortsBySegmentIDOK) GetPayload() []*models.DetailedSegmentEffort {
	return o.Payload
}

func (o *GetEffortsBySegmentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEffortsBySegmentIDDefault creates a GetEffortsBySegmentIDDefault with default headers values
func NewGetEffortsBySegmentIDDefault(code int) *GetEffortsBySegmentIDDefault {
	return &GetEffortsBySegmentIDDefault{
		_statusCode: code,
	}
}

/*GetEffortsBySegmentIDDefault handles this case with default header values.

Unexpected error.
*/
type GetEffortsBySegmentIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get efforts by segment Id default response
func (o *GetEffortsBySegmentIDDefault) Code() int {
	return o._statusCode
}

func (o *GetEffortsBySegmentIDDefault) Error() string {
	return fmt.Sprintf("[GET /segment_efforts][%d] getEffortsBySegmentId default  %+v", o._statusCode, o.Payload)
}

func (o *GetEffortsBySegmentIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetEffortsBySegmentIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
