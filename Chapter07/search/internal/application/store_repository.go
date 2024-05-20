package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/search/internal/models"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*models.Store, error)
}
