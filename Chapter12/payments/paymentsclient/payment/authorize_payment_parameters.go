// Code generated by go-swagger; DO NOT EDIT.

package payment

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

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/payments/paymentsclient/models"
)

// NewAuthorizePaymentParams creates a new AuthorizePaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthorizePaymentParams() *AuthorizePaymentParams {
	return &AuthorizePaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizePaymentParamsWithTimeout creates a new AuthorizePaymentParams object
// with the ability to set a timeout on a request.
func NewAuthorizePaymentParamsWithTimeout(timeout time.Duration) *AuthorizePaymentParams {
	return &AuthorizePaymentParams{
		timeout: timeout,
	}
}

// NewAuthorizePaymentParamsWithContext creates a new AuthorizePaymentParams object
// with the ability to set a context for a request.
func NewAuthorizePaymentParamsWithContext(ctx context.Context) *AuthorizePaymentParams {
	return &AuthorizePaymentParams{
		Context: ctx,
	}
}

// NewAuthorizePaymentParamsWithHTTPClient creates a new AuthorizePaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthorizePaymentParamsWithHTTPClient(client *http.Client) *AuthorizePaymentParams {
	return &AuthorizePaymentParams{
		HTTPClient: client,
	}
}

/* AuthorizePaymentParams contains all the parameters to send to the API endpoint
   for the authorize payment operation.

   Typically these are written to a http.Request.
*/
type AuthorizePaymentParams struct {

	// Body.
	Body *models.PaymentspbAuthorizePaymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authorize payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizePaymentParams) WithDefaults() *AuthorizePaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authorize payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizePaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authorize payment params
func (o *AuthorizePaymentParams) WithTimeout(timeout time.Duration) *AuthorizePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorize payment params
func (o *AuthorizePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorize payment params
func (o *AuthorizePaymentParams) WithContext(ctx context.Context) *AuthorizePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorize payment params
func (o *AuthorizePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorize payment params
func (o *AuthorizePaymentParams) WithHTTPClient(client *http.Client) *AuthorizePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorize payment params
func (o *AuthorizePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the authorize payment params
func (o *AuthorizePaymentParams) WithBody(body *models.PaymentspbAuthorizePaymentRequest) *AuthorizePaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authorize payment params
func (o *AuthorizePaymentParams) SetBody(body *models.PaymentspbAuthorizePaymentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
