// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/widmnd/strava-api-v3/models"
)

// GetSegmentStreamsReader is a Reader for the GetSegmentStreams structure.
type GetSegmentStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSegmentStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSegmentStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSegmentStreamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSegmentStreamsOK creates a GetSegmentStreamsOK with default headers values
func NewGetSegmentStreamsOK() *GetSegmentStreamsOK {
	return &GetSegmentStreamsOK{}
}

/*GetSegmentStreamsOK handles this case with default header values.

The set of requested streams.
*/
type GetSegmentStreamsOK struct {
	Payload *models.StreamSet
}

func (o *GetSegmentStreamsOK) Error() string {
	return fmt.Sprintf("[GET /segments/{id}/streams][%d] getSegmentStreamsOK  %+v", 200, o.Payload)
}

func (o *GetSegmentStreamsOK) GetPayload() *models.StreamSet {
	return o.Payload
}

func (o *GetSegmentStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StreamSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSegmentStreamsDefault creates a GetSegmentStreamsDefault with default headers values
func NewGetSegmentStreamsDefault(code int) *GetSegmentStreamsDefault {
	return &GetSegmentStreamsDefault{
		_statusCode: code,
	}
}

/*GetSegmentStreamsDefault handles this case with default header values.

Unexpected error.
*/
type GetSegmentStreamsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get segment streams default response
func (o *GetSegmentStreamsDefault) Code() int {
	return o._statusCode
}

func (o *GetSegmentStreamsDefault) Error() string {
	return fmt.Sprintf("[GET /segments/{id}/streams][%d] getSegmentStreams default  %+v", o._statusCode, o.Payload)
}

func (o *GetSegmentStreamsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetSegmentStreamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
