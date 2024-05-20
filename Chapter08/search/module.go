package search

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/search/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/storespb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := postgres.NewCustomerCacheRepository("search.customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))
	stores := postgres.NewStoreCacheRepository("search.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
	products := postgres.NewProductCacheRepository("search.products_cache", mono.DB(), grpc.NewProductRepository(conn))
	orders := postgres.NewOrderRepository("search.orders", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(orders),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(orders, customers, stores, products),
		"IntegrationEvents", mono.Logger(),
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

	return nil
}
