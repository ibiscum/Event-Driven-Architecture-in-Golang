package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
