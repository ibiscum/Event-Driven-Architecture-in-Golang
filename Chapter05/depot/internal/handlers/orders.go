package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers)
}
