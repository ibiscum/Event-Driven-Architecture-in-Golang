package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
)

func RegisterCustomerHandlers(customerHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return customerHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(customerspb.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerspb.CustomerRegisteredEvent,
	}, am.GroupName("search-customers"))
}
