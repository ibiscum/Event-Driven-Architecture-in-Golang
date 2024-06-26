package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/stores/storespb"
)

func RegisterStoreHandlers(storeHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreParticipatingToggledEvent,
		storespb.StoreRebrandedEvent,
	})
}
