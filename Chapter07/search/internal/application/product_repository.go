package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/search/internal/models"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*models.Product, error)
}
