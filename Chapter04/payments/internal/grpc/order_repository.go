package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/payments/internal/application"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Complete(ctx context.Context, invoiceID, orderID string) error {
	_, err := r.client.CompleteOrder(ctx, &orderingpb.CompleteOrderRequest{
		Id:        orderID,
		InvoiceId: invoiceID,
	})
	return err
}
