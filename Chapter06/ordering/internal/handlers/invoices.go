package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(invoiceHandlers, domain.OrderReadiedEvent)
}
