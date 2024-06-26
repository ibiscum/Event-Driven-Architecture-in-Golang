package handlers

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/baskets/basketspb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/depot/depotpb"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/ordering/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/ordering/internal/application/commands"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/ordering/internal/domain"
)

type integrationHandlers[T ddd.Event] struct {
	app application.App
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(app application.App) ddd.EventHandler[ddd.Event] {
	return integrationHandlers[ddd.Event]{
		app: app,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) (err error) {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, eventMsg)
	})

	err = subscriber.Subscribe(basketspb.BasketAggregateChannel, evtMsgHandler, am.MessageFilter{
		basketspb.BasketCheckedOutEvent,
	}, am.GroupName("ordering-baskets"))
	if err != nil {
		return err
	}

	err = subscriber.Subscribe(depotpb.ShoppingListAggregateChannel, evtMsgHandler, am.MessageFilter{
		depotpb.ShoppingListCompletedEvent,
	}, am.GroupName("ordering-depot"))

	return
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case basketspb.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	case depotpb.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*basketspb.BasketCheckedOut)

	items := make([]domain.Item, len(payload.GetItems()))
	for i, item := range payload.GetItems() {
		items[i] = domain.Item{
			ProductID:   item.GetProductId(),
			StoreID:     item.GetStoreId(),
			StoreName:   item.GetStoreName(),
			ProductName: item.GetProductName(),
			Price:       item.GetPrice(),
			Quantity:    int(item.GetQuantity()),
		}
	}

	return h.app.CreateOrder(ctx, commands.CreateOrder{
		ID:         payload.GetId(),
		CustomerID: payload.GetCustomerId(),
		PaymentID:  payload.GetPaymentId(),
		Items:      items,
	})
}

func (h integrationHandlers[T]) onShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*depotpb.ShoppingListCompleted)

	return h.app.ReadyOrder(ctx, commands.ReadyOrder{ID: payload.GetOrderId()})
}
