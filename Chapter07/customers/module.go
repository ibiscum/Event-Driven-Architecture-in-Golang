package customers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/customers/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/Chapter07/internal/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers, domainDispatcher),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}
