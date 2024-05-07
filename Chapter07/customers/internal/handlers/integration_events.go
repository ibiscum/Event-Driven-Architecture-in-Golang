package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/customers/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.CustomerRegisteredEvent,
		domain.CustomerSmsChangedEvent,
		domain.CustomerEnabledEvent,
		domain.CustomerDisabledEvent,
	)
}
