package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/baskets/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/internal/ddd"
)

type OrderHandlers[T ddd.AggregateEvent] struct {
	orders domain.OrderRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*OrderHandlers[ddd.AggregateEvent])(nil)

func NewOrderHandlers(orders domain.OrderRepository) OrderHandlers[ddd.AggregateEvent] {
	return OrderHandlers[ddd.AggregateEvent]{
		orders: orders,
	}
}

func (h OrderHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	}
	return nil
}

func (h OrderHandlers[T]) onBasketCheckedOut(ctx context.Context, event ddd.AggregateEvent) error {
	checkedOut := event.Payload().(*domain.BasketCheckedOut)
	_, err := h.orders.Save(ctx, checkedOut.PaymentID, checkedOut.CustomerID, checkedOut.Items)
	return err
}
