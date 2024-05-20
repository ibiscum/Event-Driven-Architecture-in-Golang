package cosec

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/cosec/internal"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/cosec/internal/handlers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/cosec/internal/logging"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/customers/customerspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/depot/depotpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/jetstream"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/monolith"
	pg "github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/postgres"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/internal/sec"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/ordering/orderingpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter08/payments/paymentspb"
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
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = depotpb.Registrations(reg); err != nil {
		return err
	}
	if err = paymentspb.Registrations(reg); err != nil {
		return err
	}
	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	sagaStore := pg.NewSagaStore("cosec.sagas", mono.DB(), reg)
	sagaRepo := sec.NewSagaRepository[*models.CreateOrderData](reg, sagaStore)

	// setup application
	orchestrator := logging.LogReplyHandlerAccess[*models.CreateOrderData](
		sec.NewOrchestrator[*models.CreateOrderData](internal.NewCreateOrderSaga(), sagaRepo, commandStream),
		"CreateOrderSaga", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(orchestrator),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlers(replyStream, orchestrator); err != nil {
		return err
	}

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
