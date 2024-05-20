package notifications

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/amotel"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/amprom"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/jetstream"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/postgresotel"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/system"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/tm"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications/internal/constants"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/ordering/orderingpb"
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
