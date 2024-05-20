// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/Chapter10/customers/customersclient/models"
)

// GetCustomerReader is a Reader for the GetCustomer structure.
type GetCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomerOK creates a GetCustomerOK with default headers values
func NewGetCustomerOK() *GetCustomerOK {
	return &GetCustomerOK{}
}

/* GetCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetCustomerOK struct {
	Payload *models.CustomerspbGetCustomerResponse
}

func (o *GetCustomerOK) Error() string {
	return fmt.Sprintf("[GET /api/customers/{id}][%d] getCustomerOK  %+v", 200, o.Payload)
}
func (o *GetCustomerOK) GetPayload() *models.CustomerspbGetCustomerResponse {
	return o.Payload
}

func (o *GetCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerspbGetCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerDefault creates a GetCustomerDefault with default headers values
func NewGetCustomerDefault(code int) *GetCustomerDefault {
	return &GetCustomerDefault{
		_statusCode: code,
	}
}

/* GetCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetCustomerDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the get customer default response
func (o *GetCustomerDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomerDefault) Error() string {
	return fmt.Sprintf("[GET /api/customers/{id}][%d] getCustomer default  %+v", o._statusCode, o.Payload)
}
func (o *GetCustomerDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
