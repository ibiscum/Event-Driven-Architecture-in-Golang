package application

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/payments/internal/models"
)

type InvoiceRepository interface {
	Find(ctx context.Context, invoiceID string) (*models.Invoice, error)
	Save(ctx context.Context, invoice *models.Invoice) error
	Update(ctx context.Context, invoice *models.Invoice) error
}
