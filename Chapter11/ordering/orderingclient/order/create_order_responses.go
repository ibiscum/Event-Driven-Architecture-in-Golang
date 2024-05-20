// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/ordering/orderingclient/models"
)

// CreateOrderReader is a Reader for the CreateOrder structure.
type CreateOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrderOK creates a CreateOrderOK with default headers values
func NewCreateOrderOK() *CreateOrderOK {
	return &CreateOrderOK{}
}

/* CreateOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOrderOK struct {
	Payload *models.OrderingpbCreateOrderResponse
}

func (o *CreateOrderOK) Error() string {
	return fmt.Sprintf("[POST /api/ordering][%d] createOrderOK  %+v", 200, o.Payload)
}
func (o *CreateOrderOK) GetPayload() *models.OrderingpbCreateOrderResponse {
	return o.Payload
}

func (o *CreateOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderingpbCreateOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrderDefault creates a CreateOrderDefault with default headers values
func NewCreateOrderDefault(code int) *CreateOrderDefault {
	return &CreateOrderDefault{
		_statusCode: code,
	}
}

/* CreateOrderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateOrderDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the create order default response
func (o *CreateOrderDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrderDefault) Error() string {
	return fmt.Sprintf("[POST /api/ordering][%d] createOrder default  %+v", o._statusCode, o.Payload)
}
func (o *CreateOrderDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
