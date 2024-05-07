package baskets

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/baskets/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/es"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/stores/storespb"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("baskets.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("baskets.snapshots", mono.DB(), reg),
	)
	baskets := es.NewAggregateRepository[*domain.Basket](domain.BasketAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, orders),
		mono.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewOrderHandlers(orders),
		"Order", mono.Logger(),
	)
	storeHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewStoreHandlers(mono.Logger()),
		"Store", mono.Logger(),
	)
	productHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewProductHandlers(mono.Logger()),
		"Product", mono.Logger(),
	)

	// setup Driver adapters
	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)
	if err = handlers.RegisterStoreHandlers(storeHandlers, eventStream); err != nil {
		return err
	}
	if err = handlers.RegisterProductHandlers(productHandlers, eventStream); err != nil {
		return err
	}

	return
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// Basket
	if err := serde.Register(domain.Basket{}, func(v interface{}) error {
		basket := v.(*domain.Basket)
		basket.Items = make(map[string]domain.Item)
		return nil
	}); err != nil {
		return err
	}
	// basket events
	if err := serde.Register(domain.BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketCheckedOut{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketItemAdded{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketItemRemoved{}); err != nil {
		return err
	}
	// basket snapshots
	if err := serde.RegisterKey(domain.BasketV1{}.SnapshotName(), domain.BasketV1{}); err != nil {
		return err
	}

	return nil
}
