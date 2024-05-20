package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/payments/internal/models"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.Event], domainSubscriber ddd.EventSubscriber[ddd.Event]) {
	domainSubscriber.Subscribe(eventHandlers,
		models.InvoicePaidEvent,
	)
}
