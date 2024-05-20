package customers

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers/customersclient"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers/customersclient/customer"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers/customersclient/models"
)

type Client interface {
	RegisterCustomer(ctx context.Context, name, smsNumber string) (string, error)
}

type client struct {
	c *customersclient.Customers
}

func NewClient(transport runtime.ClientTransport) Client {
	return &client{
		c: customersclient.New(transport, strfmt.Default),
	}
}

func (c *client) RegisterCustomer(ctx context.Context, name, smsNumber string) (string, error) {
	resp, err := c.c.Customer.RegisterCustomer(&customer.RegisterCustomerParams{
		Body: &models.CustomerspbRegisterCustomerRequest{
			Name:      name,
			SmsNumber: smsNumber,
		},
		Context: ctx,
	})
	if err != nil {
		return "", err
	}

	return resp.GetPayload().ID, nil
}
