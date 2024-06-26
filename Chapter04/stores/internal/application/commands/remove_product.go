package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter04/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

func NewRemoveProductHandler(products domain.ProductRepository, domainPublisher ddd.EventPublisher) RemoveProductHandler {
	return RemoveProductHandler{
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Remove(); err != nil {
		return err
	}

	if err = h.products.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return err
	}

	return nil
}
