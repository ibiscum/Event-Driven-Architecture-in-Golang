package depot

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/depot/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter07/stores/storespb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	shoppingLists := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := postgres.NewStoreCacheRepository("depot.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
	products := postgres.NewProductCacheRepository("depot.products_cache", mono.DB(), grpc.NewProductRepository(conn))
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(shoppingLists, stores, products, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewOrderHandlers(orders),
		"Order", mono.Logger(),
	)
	storeHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewStoreHandlers(stores),
		"Store", mono.Logger(),
	)
	productHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewProductHandlers(products),
		"Product", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)
	if err = handlers.RegisterStoreHandlers(storeHandlers, eventStream); err != nil {
		return err
	}
	if err = handlers.RegisterProductHandlers(productHandlers, eventStream); err != nil {
		return err
	}

	return nil
}
