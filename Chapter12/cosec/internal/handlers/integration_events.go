package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/errorsotel"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/sec"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/ordering/orderingpb"
)

type integrationHandlers[T ddd.Event] struct {
	orchestrator sec.Orchestrator[*models.CreateOrderData]
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, orchestrator sec.Orchestrator[*models.CreateOrderData], mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		orchestrator: orchestrator,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	_, err = subscriber.Subscribe(orderingpb.OrderAggregateChannel, handlers, am.MessageFilter{
		orderingpb.OrderCreatedEvent,
	}, am.GroupName("cosec-ordering"))
	return
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling integration event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled integration event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling integration event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case orderingpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderingpb.OrderCreated)

	var total float64
	items := make([]models.Item, len(payload.GetItems()))
	for i, item := range payload.GetItems() {
		items[i] = models.Item{
			ProductID: item.GetProductId(),
			StoreID:   item.GetStoreId(),
			Price:     item.GetPrice(),
			Quantity:  int(item.GetQuantity()),
		}
		total += float64(item.GetQuantity()) * item.GetPrice()
	}

	data := &models.CreateOrderData{
		OrderID:    payload.GetId(),
		CustomerID: payload.GetCustomerId(),
		PaymentID:  payload.GetPaymentId(),
		Items:      items,
		Total:      total,
	}

	// Start the CreateOrderSaga
	return h.orchestrator.Start(ctx, event.ID(), data)
}
