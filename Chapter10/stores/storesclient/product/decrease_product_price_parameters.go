// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/storesclient/models"
)

// NewDecreaseProductPriceParams creates a new DecreaseProductPriceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDecreaseProductPriceParams() *DecreaseProductPriceParams {
	return &DecreaseProductPriceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDecreaseProductPriceParamsWithTimeout creates a new DecreaseProductPriceParams object
// with the ability to set a timeout on a request.
func NewDecreaseProductPriceParamsWithTimeout(timeout time.Duration) *DecreaseProductPriceParams {
	return &DecreaseProductPriceParams{
		timeout: timeout,
	}
}

// NewDecreaseProductPriceParamsWithContext creates a new DecreaseProductPriceParams object
// with the ability to set a context for a request.
func NewDecreaseProductPriceParamsWithContext(ctx context.Context) *DecreaseProductPriceParams {
	return &DecreaseProductPriceParams{
		Context: ctx,
	}
}

// NewDecreaseProductPriceParamsWithHTTPClient creates a new DecreaseProductPriceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDecreaseProductPriceParamsWithHTTPClient(client *http.Client) *DecreaseProductPriceParams {
	return &DecreaseProductPriceParams{
		HTTPClient: client,
	}
}

/* DecreaseProductPriceParams contains all the parameters to send to the API endpoint
   for the decrease product price operation.

   Typically these are written to a http.Request.
*/
type DecreaseProductPriceParams struct {

	// Body.
	Body *models.DecreaseProductPriceParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the decrease product price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecreaseProductPriceParams) WithDefaults() *DecreaseProductPriceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the decrease product price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecreaseProductPriceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the decrease product price params
func (o *DecreaseProductPriceParams) WithTimeout(timeout time.Duration) *DecreaseProductPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decrease product price params
func (o *DecreaseProductPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decrease product price params
func (o *DecreaseProductPriceParams) WithContext(ctx context.Context) *DecreaseProductPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decrease product price params
func (o *DecreaseProductPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decrease product price params
func (o *DecreaseProductPriceParams) WithHTTPClient(client *http.Client) *DecreaseProductPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decrease product price params
func (o *DecreaseProductPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the decrease product price params
func (o *DecreaseProductPriceParams) WithBody(body *models.DecreaseProductPriceParamsBody) *DecreaseProductPriceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the decrease product price params
func (o *DecreaseProductPriceParams) SetBody(body *models.DecreaseProductPriceParamsBody) {
	o.Body = body
}

// WithID adds the id to the decrease product price params
func (o *DecreaseProductPriceParams) WithID(id string) *DecreaseProductPriceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the decrease product price params
func (o *DecreaseProductPriceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DecreaseProductPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
