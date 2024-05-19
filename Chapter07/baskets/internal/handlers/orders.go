package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/baskets/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(orderHandlers, domain.BasketCheckedOutEvent)
}
