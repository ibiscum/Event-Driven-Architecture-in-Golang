package customers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/customers/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/customers/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/customers/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/customers/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/customers/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers, domainDispatcher),
		mono.Logger(),
	)

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
