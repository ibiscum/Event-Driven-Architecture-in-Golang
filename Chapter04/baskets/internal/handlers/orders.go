package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
