package queries

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall domain.MallRepository
}

func NewGetStoresHandler(mall domain.MallRepository) GetStoresHandler {
	return GetStoresHandler{mall: mall}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.MallStore, error) {
	return h.mall.All(ctx)
}
