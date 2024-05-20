package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/di"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.Event](func(ctx context.Context, event ddd.Event) error {
		domainHandlers := di.Get(ctx, "domainEventHandlers").(ddd.EventHandler[ddd.Event])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.Event])

	RegisterDomainEventHandlers(subscriber, handlers)
}
