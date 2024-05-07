package ordering

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/es"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/orderingpb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("ordering.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("ordering.snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, shopping)
	app = logging.LogApplicationAccess(app, mono.Logger())
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
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
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err := serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}
	// order events
	if err := serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err := serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}
