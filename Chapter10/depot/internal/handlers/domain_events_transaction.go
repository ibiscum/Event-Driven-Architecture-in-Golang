package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/di"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		domainHandlers := di.Get(ctx, "domainEventHandlers").(ddd.EventHandler[ddd.AggregateEvent])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent])

	RegisterDomainEventHandlers(subscriber, handlers)
}
