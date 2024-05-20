package commands

import (
	"context"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/stores/internal/domain"
)

type IncreaseProductPrice struct {
	ID    string
	Price float64
}

type IncreaseProductPriceHandler struct {
	products  domain.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewIncreaseProductPriceHandler(products domain.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) IncreaseProductPriceHandler {
	return IncreaseProductPriceHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h IncreaseProductPriceHandler) IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.IncreasePrice(cmd.Price)
	if err != nil {
		return err
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
