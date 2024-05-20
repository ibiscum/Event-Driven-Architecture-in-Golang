package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/Chapter10/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/Chapter10/internal/ddd"
)

type CompleteShoppingList struct {
	ID string
}

type CompleteShoppingListHandler struct {
	shoppingLists   domain.ShoppingListRepository
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewCompleteShoppingListHandler(shoppingLists domain.ShoppingListRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent],
) CompleteShoppingListHandler {
	return CompleteShoppingListHandler{
		shoppingLists:   shoppingLists,
		domainPublisher: domainPublisher,
	}
}

func (h CompleteShoppingListHandler) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = list.Complete(); err != nil {
		return err
	}

	if err = h.shoppingLists.Update(ctx, list); err != nil {
		return nil
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
