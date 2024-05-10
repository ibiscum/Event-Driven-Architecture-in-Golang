package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/baskets/internal/domain"
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

func (h OrderHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	checkedOut := event.(*domain.BasketCheckedOut)
	_, err := h.orders.Save(ctx, checkedOut.Basket)
	return err
}
