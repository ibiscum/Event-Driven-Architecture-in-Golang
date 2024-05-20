package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/stores/internal/domain"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	stores    domain.StoreRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewEnableParticipationHandler(stores domain.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) EnableParticipationHandler {
	return EnableParticipationHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.EnableParticipation()
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
