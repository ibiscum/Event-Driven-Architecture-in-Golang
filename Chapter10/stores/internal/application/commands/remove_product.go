package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	products  domain.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewRemoveProductHandler(products domain.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) RemoveProductHandler {
	return RemoveProductHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.Remove()
	if err != nil {
		return err
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
