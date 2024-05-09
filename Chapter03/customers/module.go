package customers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	var app application.App
	app = application.New(customers)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
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
