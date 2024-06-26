package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/payments/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/payments/paymentspb"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(app application.App) ddd.CommandHandler[ddd.Command] {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.RawMessageSubscriber, handlers am.RawMessageHandler) error {
	_, err := subscriber.Subscribe(paymentspb.CommandChannel, handlers, am.MessageFilter{
		paymentspb.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case paymentspb.ConfirmPaymentCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doConfirmPayment(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*paymentspb.ConfirmPayment)

	return nil, h.app.ConfirmPayment(ctx, application.ConfirmPayment{ID: payload.GetId()})
}
