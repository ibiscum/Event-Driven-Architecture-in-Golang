package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/search/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}

type CustomerCacheRepository interface {
	Add(ctx context.Context, customerID, name string) error
	CustomerRepository
}
