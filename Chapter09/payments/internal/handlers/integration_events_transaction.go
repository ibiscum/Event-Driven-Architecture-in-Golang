package handlers

import (
	"context"
	"database/sql"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/di"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/ordering/orderingpb"
)

func RegisterIntegrationEventHandlersTx(container di.Container) error {
	evtMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
		ctx = container.Scoped(ctx)
		defer func(tx *sql.Tx) {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}(di.Get(ctx, "tx").(*sql.Tx))

		evtHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewEventMessageHandler(
				di.Get(ctx, "registry").(registry.Registry),
				di.Get(ctx, "integrationEventHandlers").(ddd.EventHandler[ddd.Event]),
			),
			di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return evtHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	return subscriber.Subscribe(orderingpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderingpb.OrderReadiedEvent,
	}, am.GroupName("payment-orders"))
}
