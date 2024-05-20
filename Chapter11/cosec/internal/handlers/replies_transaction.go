package handlers

import (
	"context"
	"database/sql"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/cosec/internal"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/di"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/sec"
)

func RegisterReplyHandlersTx(container di.Container) error {
	replyMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
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

		replyHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewReplyMessageHandler(
				di.Get(ctx, "registry").(registry.Registry),
				di.Get(ctx, "orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
			),
			di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return replyHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	_, err := subscriber.Subscribe(internal.CreateOrderReplyChannel, replyMsgHandler, am.GroupName("cosec-replies"))
	return err
}
