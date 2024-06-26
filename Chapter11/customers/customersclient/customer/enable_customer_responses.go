// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/customers/customersclient/models"
)

// EnableCustomerReader is a Reader for the EnableCustomer structure.
type EnableCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableCustomerOK creates a EnableCustomerOK with default headers values
func NewEnableCustomerOK() *EnableCustomerOK {
	return &EnableCustomerOK{}
}

/* EnableCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnableCustomerOK struct {
	Payload models.CustomerspbEnableCustomerResponse
}

func (o *EnableCustomerOK) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomerOK  %+v", 200, o.Payload)
}
func (o *EnableCustomerOK) GetPayload() models.CustomerspbEnableCustomerResponse {
	return o.Payload
}

func (o *EnableCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableCustomerDefault creates a EnableCustomerDefault with default headers values
func NewEnableCustomerDefault(code int) *EnableCustomerDefault {
	return &EnableCustomerDefault{
		_statusCode: code,
	}
}

/* EnableCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnableCustomerDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the enable customer default response
func (o *EnableCustomerDefault) Code() int {
	return o._statusCode
}

func (o *EnableCustomerDefault) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomer default  %+v", o._statusCode, o.Payload)
}
func (o *EnableCustomerDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *EnableCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
