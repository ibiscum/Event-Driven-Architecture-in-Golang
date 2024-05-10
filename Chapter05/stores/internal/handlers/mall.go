package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.StoreCreatedEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreParticipationEnabledEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreParticipationDisabledEvent, mallHandlers)
	domainSubscriber.Subscribe(domain.StoreRebrandedEvent, mallHandlers)
}
