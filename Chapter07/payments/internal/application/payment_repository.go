package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/payments/internal/models"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *models.Payment) error
	Find(ctx context.Context, paymentID string) (*models.Payment, error)
}
