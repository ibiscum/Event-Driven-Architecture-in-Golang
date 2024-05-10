package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.OrderCreatedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderCanceledEvent, notificationHandlers)
}
