package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
)

type OrderHandlers struct {
	orders domain.OrderRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*OrderHandlers)(nil)

func NewOrderHandlers(orders domain.OrderRepository) OrderHandlers {
	return OrderHandlers{
		orders: orders,
	}
}

func (h OrderHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	completed := event.(*domain.ShoppingListCompleted)
	return h.orders.Ready(ctx, completed.ShoppingList.OrderID)
}
