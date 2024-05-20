package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/ordering/internal/domain"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orders    domain.OrderRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewReadyOrderHandler(orders domain.OrderRepository, publisher ddd.EventPublisher[ddd.Event]) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders:    orders,
		publisher: publisher,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Ready()
	if err != nil {
		return nil
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
