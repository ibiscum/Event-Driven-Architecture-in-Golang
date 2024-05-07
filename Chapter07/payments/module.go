package payments

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/payments/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	invoices := postgres.NewInvoiceRepository("payments.invoices", mono.DB())
	payments := postgres.NewPaymentRepository("payments.payments", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(invoices, payments, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewOrderHandlers(app),
		"Order", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterOrderHandlers(orderHandlers, eventStream); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}
