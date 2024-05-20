package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
)

func RegisterCustomerHandlers(customerHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return customerHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(customerspb.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerspb.CustomerRegisteredEvent,
		customerspb.CustomerSmsChangedEvent,
	}, am.GroupName("notification-customers"))
}
