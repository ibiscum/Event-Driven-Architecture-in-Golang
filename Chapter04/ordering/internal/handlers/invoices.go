package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
