package postgres

import (
	"context"
	"fmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/models"
)

type PaymentRepository struct {
	tableName string
	db        postgres.DB
}

var _ application.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(tableName string, db postgres.DB) PaymentRepository {
	return PaymentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r PaymentRepository) Save(ctx context.Context, payment *models.Payment) error {
	const query = "INSERT INTO %s (id, customer_id, amount) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, r.table(query), payment.ID, payment.CustomerID, payment.Amount)

	return err
}

func (r PaymentRepository) Find(ctx context.Context, paymentID string) (*models.Payment, error) {
	const query = "SELECT customer_id, amount FROM %s WHERE id = $1 LIMIT 1"

	payment := &models.Payment{
		ID: paymentID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), paymentID).Scan(&payment.CustomerID, &payment.Amount)

	return payment, err
}

func (r PaymentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
