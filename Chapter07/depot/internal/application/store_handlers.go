package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/stores/storespb"
)

type StoreHandlers[T ddd.Event] struct {
	cache domain.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*StoreHandlers[ddd.Event])(nil)

func NewStoreHandlers(cache domain.StoreCacheRepository) StoreHandlers[ddd.Event] {
	return StoreHandlers[ddd.Event]{
		cache: cache,
	}
}

func (h StoreHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}

	return nil
}

func (h StoreHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreCreated)
	return h.cache.Add(ctx, payload.GetId(), payload.GetName(), payload.GetLocation())
}

func (h StoreHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreRebranded)
	return h.cache.Rename(ctx, payload.GetId(), payload.GetName())
}
