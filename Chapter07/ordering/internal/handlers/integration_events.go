package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/ordering/internal/domain"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.OrderCreatedEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
		domain.OrderCompletedEvent,
	)
}
