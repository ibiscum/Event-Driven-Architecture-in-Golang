package notifications

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/amotel"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/amprom"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/jetstream"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/postgresotel"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/system"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/tm"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/constants"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/notifications/internal/handlers"
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
	inboxStore := pg.NewInboxStore(constants.InboxTableName, svc.DB())
	messageSubscriber := am.NewMessageSubscriber(
		jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger()),
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(constants.ServiceName),
	)
	customers := postgres.NewCustomerCacheRepository(
		constants.CustomersCacheTableName,
		postgresotel.Trace(svc.DB()),
		grpc.NewCustomerRepository(svc.Config().Rpc.Service(constants.CustomersServiceName)),
	)

	// setup application
	app := application.New(customers)
	integrationEventHandlers := handlers.NewIntegrationEventHandlers(
		reg, app, customers,
		tm.InboxHandler(inboxStore),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, svc.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}
