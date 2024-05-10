package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(notificationHandlers,
		domain.OrderCreatedEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
	)
}
