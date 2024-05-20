// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/stores/storesclient/models"
)

// DecreaseProductPriceReader is a Reader for the DecreaseProductPrice structure.
type DecreaseProductPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecreaseProductPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecreaseProductPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDecreaseProductPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDecreaseProductPriceOK creates a DecreaseProductPriceOK with default headers values
func NewDecreaseProductPriceOK() *DecreaseProductPriceOK {
	return &DecreaseProductPriceOK{}
}

/* DecreaseProductPriceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DecreaseProductPriceOK struct {
	Payload models.StorespbDecreaseProductPriceResponse
}

func (o *DecreaseProductPriceOK) Error() string {
	return fmt.Sprintf("[PUT /api/stores/products/{id}/decreasePrice][%d] decreaseProductPriceOK  %+v", 200, o.Payload)
}
func (o *DecreaseProductPriceOK) GetPayload() models.StorespbDecreaseProductPriceResponse {
	return o.Payload
}

func (o *DecreaseProductPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseProductPriceDefault creates a DecreaseProductPriceDefault with default headers values
func NewDecreaseProductPriceDefault(code int) *DecreaseProductPriceDefault {
	return &DecreaseProductPriceDefault{
		_statusCode: code,
	}
}

/* DecreaseProductPriceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DecreaseProductPriceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the decrease product price default response
func (o *DecreaseProductPriceDefault) Code() int {
	return o._statusCode
}

func (o *DecreaseProductPriceDefault) Error() string {
	return fmt.Sprintf("[PUT /api/stores/products/{id}/decreasePrice][%d] decreaseProductPrice default  %+v", o._statusCode, o.Payload)
}
func (o *DecreaseProductPriceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DecreaseProductPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
