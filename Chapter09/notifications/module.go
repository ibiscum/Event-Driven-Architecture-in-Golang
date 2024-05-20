package notifications

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/notifications/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/notifications/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/notifications/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/notifications/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/notifications/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/ordering/orderingpb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := postgres.NewCustomerCacheRepository("notifications.customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(app, customers),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}
