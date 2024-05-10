package queries

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter05/stores/internal/domain"
)

type GetParticipatingStores struct{}

type GetParticipatingStoresHandler struct {
	mall domain.MallRepository
}

func NewGetParticipatingStoresHandler(mall domain.MallRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{mall: mall}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.MallStore, error) {
	return h.mall.AllParticipating(ctx)
}
