package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
)

type DomainEventHandlers interface {
	OnShoppingListCreated(ctx context.Context, event ddd.Event) error
	OnShoppingListCanceled(ctx context.Context, event ddd.Event) error
	OnShoppingListAssigned(ctx context.Context, event ddd.Event) error
	OnShoppingListCompleted(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnShoppingListCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnShoppingListCanceled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnShoppingListAssigned(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	return nil
}
