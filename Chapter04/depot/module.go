package depot

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/depot/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	shoppingLists := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(application.New(shoppingLists, stores, products, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orders),
		mono.Logger(),
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

	return nil
}
