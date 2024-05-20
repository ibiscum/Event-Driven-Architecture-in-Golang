package payments

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/paymentspb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = paymentspb.Registrations(reg); err != nil {
		return err
	}
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	invoices := postgres.NewInvoiceRepository("payments.invoices", mono.DB())
	payments := postgres.NewPaymentRepository("payments.payments", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(invoices, payments, domainDispatcher),
		mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationHandlers(app),
		"IntegrationEvents", mono.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err = grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

	return
}
