package payments

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/payments/paymentsclient"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/payments/paymentsclient/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/payments/paymentsclient/payment"
)

type Client interface {
	AuthorizePayment(ctx context.Context, customerID string, amount float64) (string, error)
}

type client struct {
	c *paymentsclient.Payments
}

func NewClient(transport runtime.ClientTransport) Client {
	return &client{
		c: paymentsclient.New(transport, strfmt.Default),
	}
}

func (c *client) AuthorizePayment(ctx context.Context, customerID string, amount float64) (string, error) {
	resp, err := c.c.Payment.AuthorizePayment(&payment.AuthorizePaymentParams{
		Body: &models.PaymentspbAuthorizePaymentRequest{
			Amount:     amount,
			CustomerID: customerID,
		},
		Context: ctx,
	})
	if err != nil {
		return "", err
	}
	return resp.GetPayload().ID, nil
}
