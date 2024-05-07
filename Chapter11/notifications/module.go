package notifications

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/system"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/orderingpb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger()))
	conn, err := grpc.Dial(ctx, svc.Config().Rpc.Service("CUSTOMERS"))
	if err != nil {
		return err
	}
	customers := postgres.NewCustomerCacheRepository("notifications.customers_cache", svc.DB(), grpc.NewCustomerRepository(conn))

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers),
		svc.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(app, customers),
		"IntegrationEvents", svc.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, svc.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}
