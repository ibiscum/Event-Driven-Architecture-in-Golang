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

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/stores/storesclient/models"
)

// NewIncreaseProductPriceParams creates a new IncreaseProductPriceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIncreaseProductPriceParams() *IncreaseProductPriceParams {
	return &IncreaseProductPriceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIncreaseProductPriceParamsWithTimeout creates a new IncreaseProductPriceParams object
// with the ability to set a timeout on a request.
func NewIncreaseProductPriceParamsWithTimeout(timeout time.Duration) *IncreaseProductPriceParams {
	return &IncreaseProductPriceParams{
		timeout: timeout,
	}
}

// NewIncreaseProductPriceParamsWithContext creates a new IncreaseProductPriceParams object
// with the ability to set a context for a request.
func NewIncreaseProductPriceParamsWithContext(ctx context.Context) *IncreaseProductPriceParams {
	return &IncreaseProductPriceParams{
		Context: ctx,
	}
}

// NewIncreaseProductPriceParamsWithHTTPClient creates a new IncreaseProductPriceParams object
// with the ability to set a custom HTTPClient for a request.
func NewIncreaseProductPriceParamsWithHTTPClient(client *http.Client) *IncreaseProductPriceParams {
	return &IncreaseProductPriceParams{
		HTTPClient: client,
	}
}

/* IncreaseProductPriceParams contains all the parameters to send to the API endpoint
   for the increase product price operation.

   Typically these are written to a http.Request.
*/
type IncreaseProductPriceParams struct {

	// Body.
	Body *models.IncreaseProductPriceParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the increase product price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IncreaseProductPriceParams) WithDefaults() *IncreaseProductPriceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the increase product price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IncreaseProductPriceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the increase product price params
func (o *IncreaseProductPriceParams) WithTimeout(timeout time.Duration) *IncreaseProductPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the increase product price params
func (o *IncreaseProductPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the increase product price params
func (o *IncreaseProductPriceParams) WithContext(ctx context.Context) *IncreaseProductPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the increase product price params
func (o *IncreaseProductPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the increase product price params
func (o *IncreaseProductPriceParams) WithHTTPClient(client *http.Client) *IncreaseProductPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the increase product price params
func (o *IncreaseProductPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the increase product price params
func (o *IncreaseProductPriceParams) WithBody(body *models.IncreaseProductPriceParamsBody) *IncreaseProductPriceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the increase product price params
func (o *IncreaseProductPriceParams) SetBody(body *models.IncreaseProductPriceParamsBody) {
	o.Body = body
}

// WithID adds the id to the increase product price params
func (o *IncreaseProductPriceParams) WithID(id string) *IncreaseProductPriceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the increase product price params
func (o *IncreaseProductPriceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IncreaseProductPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
