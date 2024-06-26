package customers

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/grpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/customers/internal/rest"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/di"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/tm"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()), nil
	})
	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			pg.NewOutboxStore("customers.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return postgres.NewCustomerRepository("customers.customers", c.Get("tx").(*sql.Tx)), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("customers.outbox", tx)
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
		inboxStore := pg.NewInboxStore("customers.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.LogApplicationAccess(
			application.New(
				c.Get("customers").(domain.CustomerRepository),
				c.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent]),
			),
			c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.AggregateEvent](
			handlers.NewDomainEventHandlers(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(zerolog.Logger),
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
	if err = handlers.RegisterCommandHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return nil
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("customers outbox processor encountered an error")
		}
	}()
}
