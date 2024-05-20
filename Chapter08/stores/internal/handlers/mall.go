package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(mallHandlers,
		domain.StoreCreatedEvent,
		domain.StoreParticipationEnabledEvent,
		domain.StoreParticipationDisabledEvent,
		domain.StoreRebrandedEvent,
	)
}
