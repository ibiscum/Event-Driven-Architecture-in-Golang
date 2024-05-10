package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/baskets/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/baskets/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
