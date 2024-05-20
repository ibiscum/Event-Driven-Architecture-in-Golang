package customers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers, domainDispatcher),
		mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainEventHandlers, domainDispatcher)
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

	return nil
}
