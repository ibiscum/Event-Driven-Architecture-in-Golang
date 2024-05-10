package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(orderHandlers, domain.ShoppingListCompletedEvent)
}
