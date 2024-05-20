package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscriber, orchestrator sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return orchestrator.HandleReply(ctx, replyMsg)
	})
	_, err := subscriber.Subscribe(orchestrator.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
	return err
}
