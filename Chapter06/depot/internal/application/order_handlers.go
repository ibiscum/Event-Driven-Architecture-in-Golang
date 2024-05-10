package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
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
	case domain.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}
	return nil
}

func (h OrderHandlers[T]) onShoppingListCompleted(ctx context.Context, event ddd.AggregateEvent) error {
	completed := event.Payload().(*domain.ShoppingListCompleted)
	return h.orders.Ready(ctx, completed.ShoppingList.OrderID)
}
