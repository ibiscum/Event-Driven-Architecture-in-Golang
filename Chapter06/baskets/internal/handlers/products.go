package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/stores/storespb"
)

func RegisterProductHandlers(productHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return productHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductPriceIncreasedEvent,
		storespb.ProductPriceDecreasedEvent,
		storespb.ProductRemovedEvent,
	})
}
