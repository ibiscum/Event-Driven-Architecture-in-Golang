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

// NewAddProductParams creates a new AddProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddProductParams() *AddProductParams {
	return &AddProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddProductParamsWithTimeout creates a new AddProductParams object
// with the ability to set a timeout on a request.
func NewAddProductParamsWithTimeout(timeout time.Duration) *AddProductParams {
	return &AddProductParams{
		timeout: timeout,
	}
}

// NewAddProductParamsWithContext creates a new AddProductParams object
// with the ability to set a context for a request.
func NewAddProductParamsWithContext(ctx context.Context) *AddProductParams {
	return &AddProductParams{
		Context: ctx,
	}
}

// NewAddProductParamsWithHTTPClient creates a new AddProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddProductParamsWithHTTPClient(client *http.Client) *AddProductParams {
	return &AddProductParams{
		HTTPClient: client,
	}
}

/* AddProductParams contains all the parameters to send to the API endpoint
   for the add product operation.

   Typically these are written to a http.Request.
*/
type AddProductParams struct {

	// Body.
	Body *models.AddProductParamsBody

	// StoreID.
	StoreID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProductParams) WithDefaults() *AddProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add product params
func (o *AddProductParams) WithTimeout(timeout time.Duration) *AddProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add product params
func (o *AddProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add product params
func (o *AddProductParams) WithContext(ctx context.Context) *AddProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add product params
func (o *AddProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add product params
func (o *AddProductParams) WithHTTPClient(client *http.Client) *AddProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add product params
func (o *AddProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add product params
func (o *AddProductParams) WithBody(body *models.AddProductParamsBody) *AddProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add product params
func (o *AddProductParams) SetBody(body *models.AddProductParamsBody) {
	o.Body = body
}

// WithStoreID adds the storeID to the add product params
func (o *AddProductParams) WithStoreID(storeID string) *AddProductParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the add product params
func (o *AddProductParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *AddProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param storeId
	if err := r.SetPathParam("storeId", o.StoreID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
