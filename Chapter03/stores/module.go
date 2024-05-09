package stores

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	stores := postgres.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := postgres.NewProductRepository("stores.products", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
