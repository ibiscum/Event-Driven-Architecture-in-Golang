package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/depot/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/storespb"
)

type ProductRepository struct {
	client storespb.StoresServiceClient
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(conn *grpc.ClientConn) ProductRepository {
	return ProductRepository{client: storespb.NewStoresServiceClient(conn)}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	resp, err := r.client.GetProduct(ctx, &storespb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(resp.Product), nil
}

func (r ProductRepository) productToDomain(product *storespb.Product) *domain.Product {
	return &domain.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}
