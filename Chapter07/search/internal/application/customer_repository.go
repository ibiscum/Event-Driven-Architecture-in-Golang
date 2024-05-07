package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/search/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
