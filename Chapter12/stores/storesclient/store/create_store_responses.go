// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/stores/storesclient/models"
)

// CreateStoreReader is a Reader for the CreateStore structure.
type CreateStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateStoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStoreOK creates a CreateStoreOK with default headers values
func NewCreateStoreOK() *CreateStoreOK {
	return &CreateStoreOK{}
}

/* CreateStoreOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateStoreOK struct {
	Payload *models.StorespbCreateStoreResponse
}

// IsSuccess returns true when this create store o k response has a 2xx status code
func (o *CreateStoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create store o k response has a 3xx status code
func (o *CreateStoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create store o k response has a 4xx status code
func (o *CreateStoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create store o k response has a 5xx status code
func (o *CreateStoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create store o k response a status code equal to that given
func (o *CreateStoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateStoreOK) Error() string {
	return fmt.Sprintf("[POST /api/stores][%d] createStoreOK  %+v", 200, o.Payload)
}

func (o *CreateStoreOK) String() string {
	return fmt.Sprintf("[POST /api/stores][%d] createStoreOK  %+v", 200, o.Payload)
}

func (o *CreateStoreOK) GetPayload() *models.StorespbCreateStoreResponse {
	return o.Payload
}

func (o *CreateStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorespbCreateStoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStoreDefault creates a CreateStoreDefault with default headers values
func NewCreateStoreDefault(code int) *CreateStoreDefault {
	return &CreateStoreDefault{
		_statusCode: code,
	}
}

/* CreateStoreDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateStoreDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the create store default response
func (o *CreateStoreDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create store default response has a 2xx status code
func (o *CreateStoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create store default response has a 3xx status code
func (o *CreateStoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create store default response has a 4xx status code
func (o *CreateStoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create store default response has a 5xx status code
func (o *CreateStoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create store default response a status code equal to that given
func (o *CreateStoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateStoreDefault) Error() string {
	return fmt.Sprintf("[POST /api/stores][%d] createStore default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStoreDefault) String() string {
	return fmt.Sprintf("[POST /api/stores][%d] createStore default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStoreDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateStoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
