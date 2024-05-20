package stores

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/es"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/stores/storespb"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("stores.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("stores.snapshots", mono.DB(), reg),
	)
	stores := es.NewAggregateRepository[*domain.Store](domain.StoreAggregate, reg, aggregateStore)
	products := es.NewAggregateRepository[*domain.Product](domain.ProductAggregate, reg, aggregateStore)
	catalog := postgres.NewCatalogRepository("stores.products", mono.DB())
	mall := postgres.NewMallRepository("stores.stores", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(stores, products, catalog, mall),
		mono.Logger(),
	)
	catalogHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewCatalogHandlers(catalog),
		"Catalog", mono.Logger(),
	)
	mallHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewMallHandlers(mall),
		"Mall", mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
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
	handlers.RegisterCatalogHandlers(catalogHandlers, domainDispatcher)
	handlers.RegisterMallHandlers(mallHandlers, domainDispatcher)
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = storespb.RegisterAsyncAPI(mono.Mux()); err != nil {
		return err
	}

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(domain.Store{}, func(v any) error {
		store := v.(*domain.Store)
		store.Aggregate = es.NewAggregate("", domain.StoreAggregate)
		return nil
	}); err != nil {
		return
	}
	// store events
	if err = serde.Register(domain.StoreCreated{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.StoreParticipationEnabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.StoreParticipationDisabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.Register(domain.StoreRebranded{}); err != nil {
		return
	}
	// store snapshots
	if err = serde.RegisterKey(domain.StoreV1{}.SnapshotName(), domain.StoreV1{}); err != nil {
		return
	}

	// Product
	if err = serde.Register(domain.Product{}, func(v any) error {
		store := v.(*domain.Product)
		store.Aggregate = es.NewAggregate("", domain.ProductAggregate)
		return nil
	}); err != nil {
		return
	}
	// product events
	if err = serde.Register(domain.ProductAdded{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRebranded{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.ProductPriceIncreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.ProductPriceDecreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRemoved{}); err != nil {
		return
	}
	// product snapshots
	if err = serde.RegisterKey(domain.ProductV1{}.SnapshotName(), domain.ProductV1{}); err != nil {
		return
	}

	return
}
