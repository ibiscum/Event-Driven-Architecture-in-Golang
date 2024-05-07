// Code generated by go-swagger; DO NOT EDIT.

package participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/stores/storesclient/models"
)

// DisableParticipationReader is a Reader for the DisableParticipation structure.
type DisableParticipationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableParticipationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableParticipationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableParticipationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableParticipationOK creates a DisableParticipationOK with default headers values
func NewDisableParticipationOK() *DisableParticipationOK {
	return &DisableParticipationOK{}
}

/* DisableParticipationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DisableParticipationOK struct {
	Payload models.StorespbDisableParticipationResponse
}

func (o *DisableParticipationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/stores/{id}/participating][%d] disableParticipationOK  %+v", 200, o.Payload)
}
func (o *DisableParticipationOK) GetPayload() models.StorespbDisableParticipationResponse {
	return o.Payload
}

func (o *DisableParticipationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableParticipationDefault creates a DisableParticipationDefault with default headers values
func NewDisableParticipationDefault(code int) *DisableParticipationDefault {
	return &DisableParticipationDefault{
		_statusCode: code,
	}
}

/* DisableParticipationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DisableParticipationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the disable participation default response
func (o *DisableParticipationDefault) Code() int {
	return o._statusCode
}

func (o *DisableParticipationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/stores/{id}/participating][%d] disableParticipation default  %+v", o._statusCode, o.Payload)
}
func (o *DisableParticipationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DisableParticipationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
