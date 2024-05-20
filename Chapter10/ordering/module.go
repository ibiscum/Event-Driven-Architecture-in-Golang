package ordering

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/baskets/basketspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/depot/depotpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/di"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/es"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/tm"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/orderingpb"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := basketspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotpb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), c.Get("logger").(zerolog.Logger)), nil
	})
	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return grpc.Dial(ctx, mono.Config().Rpc.Address())
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			pg.NewOutboxStore("ordering.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("ordering.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		inboxStore := pg.NewInboxStore("ordering.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			pg.NewEventStore("ordering.events", tx, reg),
			pg.NewSnapshotStore("ordering.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Order](
			domain.OrderAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.LogApplicationAccess(
			application.New(
				c.Get("orders").(domain.OrderRepository),
				c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.Event]),
			),
			c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.Event](
			handlers.NewDomainEventHandlers(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.Event](
			handlers.NewIntegrationEventHandlers(
				c.Get("app").(application.App),
			),
			"IntegrationEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.LogCommandHandlerAccess[ddd.Command](
			handlers.NewCommandHandlers(c.Get("app").(application.App)),
			"Commands", c.Get("logger").(zerolog.Logger),
		), nil
	})

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(container)
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err = serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}
	// order events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderRejected{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderApproved{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err = serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("ordering outbox processor encountered an error")
		}
	}()
}
