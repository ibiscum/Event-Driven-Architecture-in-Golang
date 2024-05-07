package baskets

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	baskets := postgres.NewBasketRepository("baskets.baskets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, orders, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orders),
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return
}
