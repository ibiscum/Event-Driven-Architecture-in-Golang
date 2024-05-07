// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/basketsclient/models"
)

// StartBasketReader is a Reader for the StartBasket structure.
type StartBasketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartBasketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartBasketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartBasketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartBasketOK creates a StartBasketOK with default headers values
func NewStartBasketOK() *StartBasketOK {
	return &StartBasketOK{}
}

/* StartBasketOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartBasketOK struct {
	Payload *models.BasketspbStartBasketResponse
}

// IsSuccess returns true when this start basket o k response has a 2xx status code
func (o *StartBasketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start basket o k response has a 3xx status code
func (o *StartBasketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start basket o k response has a 4xx status code
func (o *StartBasketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start basket o k response has a 5xx status code
func (o *StartBasketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start basket o k response a status code equal to that given
func (o *StartBasketOK) IsCode(code int) bool {
	return code == 200
}

func (o *StartBasketOK) Error() string {
	return fmt.Sprintf("[POST /api/baskets][%d] startBasketOK  %+v", 200, o.Payload)
}

func (o *StartBasketOK) String() string {
	return fmt.Sprintf("[POST /api/baskets][%d] startBasketOK  %+v", 200, o.Payload)
}

func (o *StartBasketOK) GetPayload() *models.BasketspbStartBasketResponse {
	return o.Payload
}

func (o *StartBasketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasketspbStartBasketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBasketDefault creates a StartBasketDefault with default headers values
func NewStartBasketDefault(code int) *StartBasketDefault {
	return &StartBasketDefault{
		_statusCode: code,
	}
}

/* StartBasketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartBasketDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the start basket default response
func (o *StartBasketDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this start basket default response has a 2xx status code
func (o *StartBasketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start basket default response has a 3xx status code
func (o *StartBasketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start basket default response has a 4xx status code
func (o *StartBasketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start basket default response has a 5xx status code
func (o *StartBasketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start basket default response a status code equal to that given
func (o *StartBasketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StartBasketDefault) Error() string {
	return fmt.Sprintf("[POST /api/baskets][%d] startBasket default  %+v", o._statusCode, o.Payload)
}

func (o *StartBasketDefault) String() string {
	return fmt.Sprintf("[POST /api/baskets][%d] startBasket default  %+v", o._statusCode, o.Payload)
}

func (o *StartBasketDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *StartBasketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
