package ordering

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, shopping, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())
	// setup application handlers
	notificationHandlers := logging.LogDomainEventHandlerAccess(
		application.NewNotificationHandlers(notifications),
		mono.Logger(),
	)
	invoiceHandlers := logging.LogDomainEventHandlerAccess(
		application.NewInvoiceHandlers(invoices),
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)
	handlers.RegisterInvoiceHandlers(invoiceHandlers, domainDispatcher)

	return nil
}
