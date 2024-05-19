package cosec

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/cosec/internal"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/cosec/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/cosec/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/depot/depotpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/di"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/sec"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/tm"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/payments/paymentspb"
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
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := paymentspb.Registrations(reg); err != nil {
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
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			pg.NewOutboxStore("cosec.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := pg.NewOutboxStore("cosec.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("commandStream", func(c di.Container) (any, error) {
		return am.NewCommandStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		inboxStore := pg.NewInboxStore("cosec.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("sagaRepo", func(c di.Container) (any, error) {
		reg := c.Get("registry").(registry.Registry)
		return sec.NewSagaRepository[*models.CreateOrderData](
			reg,
			pg.NewSagaStore(
				"cosec.sagas",
				c.Get("tx").(*sql.Tx),
				reg,
			),
		), nil
	})
	container.AddSingleton("saga", func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	container.AddScoped("orchestrator", func(c di.Container) (any, error) {
		return logging.LogReplyHandlerAccess[*models.CreateOrderData](
			sec.NewOrchestrator[*models.CreateOrderData](
				c.Get("saga").(sec.Saga[*models.CreateOrderData]),
				c.Get("sagaRepo").(sec.SagaRepository[*models.CreateOrderData]),
				c.Get("commandStream").(am.CommandStream),
			),
			"CreateOrderSaga", mono.Logger(),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.Event](
			handlers.NewIntegrationEventHandlers(
				c.Get("orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
			),
			"IntegrationEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, container)

	return
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
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
			logger.Error().Err(err).Msg("cosec outbox processor encountered an error")
		}
	}()
}
