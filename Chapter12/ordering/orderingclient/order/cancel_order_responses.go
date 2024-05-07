// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/orderingclient/models"
)

// CancelOrderReader is a Reader for the CancelOrder structure.
type CancelOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelOrderOK creates a CancelOrderOK with default headers values
func NewCancelOrderOK() *CancelOrderOK {
	return &CancelOrderOK{}
}

/* CancelOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type CancelOrderOK struct {
	Payload models.OrderingpbCancelOrderResponse
}

// IsSuccess returns true when this cancel order o k response has a 2xx status code
func (o *CancelOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel order o k response has a 3xx status code
func (o *CancelOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel order o k response has a 4xx status code
func (o *CancelOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel order o k response has a 5xx status code
func (o *CancelOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel order o k response a status code equal to that given
func (o *CancelOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *CancelOrderOK) Error() string {
	return fmt.Sprintf("[DELETE /api/ordering/{id}][%d] cancelOrderOK  %+v", 200, o.Payload)
}

func (o *CancelOrderOK) String() string {
	return fmt.Sprintf("[DELETE /api/ordering/{id}][%d] cancelOrderOK  %+v", 200, o.Payload)
}

func (o *CancelOrderOK) GetPayload() models.OrderingpbCancelOrderResponse {
	return o.Payload
}

func (o *CancelOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelOrderDefault creates a CancelOrderDefault with default headers values
func NewCancelOrderDefault(code int) *CancelOrderDefault {
	return &CancelOrderDefault{
		_statusCode: code,
	}
}

/* CancelOrderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CancelOrderDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the cancel order default response
func (o *CancelOrderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cancel order default response has a 2xx status code
func (o *CancelOrderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel order default response has a 3xx status code
func (o *CancelOrderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel order default response has a 4xx status code
func (o *CancelOrderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel order default response has a 5xx status code
func (o *CancelOrderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel order default response a status code equal to that given
func (o *CancelOrderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CancelOrderDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/ordering/{id}][%d] cancelOrder default  %+v", o._statusCode, o.Payload)
}

func (o *CancelOrderDefault) String() string {
	return fmt.Sprintf("[DELETE /api/ordering/{id}][%d] cancelOrder default  %+v", o._statusCode, o.Payload)
}

func (o *CancelOrderDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CancelOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
