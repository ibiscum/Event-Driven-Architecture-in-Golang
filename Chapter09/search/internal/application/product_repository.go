package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/search/internal/models"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*models.Product, error)
}

type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string) error
	Rebrand(ctx context.Context, productID, name string) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}
