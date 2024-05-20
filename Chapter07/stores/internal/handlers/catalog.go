package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(catalogHandlers,
		domain.ProductAddedEvent,
		domain.ProductRebrandedEvent,
		domain.ProductPriceIncreasedEvent,
		domain.ProductPriceDecreasedEvent,
		domain.ProductRemovedEvent,
	)
}
