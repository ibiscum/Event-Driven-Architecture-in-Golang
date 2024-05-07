// Code generated by go-swagger; DO NOT EDIT.

package shopping_list

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

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/depot/depotclient/models"
)

// NewCreateShoppingListParams creates a new CreateShoppingListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateShoppingListParams() *CreateShoppingListParams {
	return &CreateShoppingListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateShoppingListParamsWithTimeout creates a new CreateShoppingListParams object
// with the ability to set a timeout on a request.
func NewCreateShoppingListParamsWithTimeout(timeout time.Duration) *CreateShoppingListParams {
	return &CreateShoppingListParams{
		timeout: timeout,
	}
}

// NewCreateShoppingListParamsWithContext creates a new CreateShoppingListParams object
// with the ability to set a context for a request.
func NewCreateShoppingListParamsWithContext(ctx context.Context) *CreateShoppingListParams {
	return &CreateShoppingListParams{
		Context: ctx,
	}
}

// NewCreateShoppingListParamsWithHTTPClient creates a new CreateShoppingListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateShoppingListParamsWithHTTPClient(client *http.Client) *CreateShoppingListParams {
	return &CreateShoppingListParams{
		HTTPClient: client,
	}
}

/* CreateShoppingListParams contains all the parameters to send to the API endpoint
   for the create shopping list operation.

   Typically these are written to a http.Request.
*/
type CreateShoppingListParams struct {

	// Body.
	Body *models.DepotpbCreateShoppingListRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create shopping list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShoppingListParams) WithDefaults() *CreateShoppingListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create shopping list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShoppingListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create shopping list params
func (o *CreateShoppingListParams) WithTimeout(timeout time.Duration) *CreateShoppingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create shopping list params
func (o *CreateShoppingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create shopping list params
func (o *CreateShoppingListParams) WithContext(ctx context.Context) *CreateShoppingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create shopping list params
func (o *CreateShoppingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create shopping list params
func (o *CreateShoppingListParams) WithHTTPClient(client *http.Client) *CreateShoppingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create shopping list params
func (o *CreateShoppingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create shopping list params
func (o *CreateShoppingListParams) WithBody(body *models.DepotpbCreateShoppingListRequest) *CreateShoppingListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create shopping list params
func (o *CreateShoppingListParams) SetBody(body *models.DepotpbCreateShoppingListRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateShoppingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
