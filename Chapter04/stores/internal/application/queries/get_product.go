package queries

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/stores/internal/domain"
)

type GetProduct struct {
	ID string
}

type GetProductHandler struct {
	products domain.ProductRepository
}

func NewGetProductHandler(products domain.ProductRepository) GetProductHandler {
	return GetProductHandler{products: products}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.Product, error) {
	return h.products.Find(ctx, query.ID)
}
