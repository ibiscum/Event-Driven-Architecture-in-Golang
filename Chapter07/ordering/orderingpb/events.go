package orderingpb

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/registry/serdes"
)

const (
	OrderAggregateChannel = "mallbots.ordering.events.Order"

	OrderCreatedEvent   = "ordersapi.OrderCreated"
	OrderReadiedEvent   = "ordersapi.OrderReadied"
	OrderCanceledEvent  = "ordersapi.OrderCanceled"
	OrderCompletedEvent = "ordersapi.OrderCompleted"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Order events
	if err := serde.Register(&OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderCompleted{}); err != nil {
		return err
	}

	return nil
}

func (*OrderCreated) Key() string   { return OrderCreatedEvent }
func (*OrderReadied) Key() string   { return OrderReadiedEvent }
func (*OrderCanceled) Key() string  { return OrderCanceledEvent }
func (*OrderCompleted) Key() string { return OrderCompletedEvent }
