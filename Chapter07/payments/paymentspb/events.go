package paymentspb

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/registry/serdes"
)

const (
	InvoiceAggregateChannel = "mallbots.payments.events.Invoice"

	InvoicePaidEvent = "paymentsapi.InvoicePaid"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Invoice events
	if err := serde.Register(&InvoicePaid{}); err != nil {
		return err
	}

	return nil
}

func (*InvoicePaid) Key() string { return InvoicePaidEvent }
