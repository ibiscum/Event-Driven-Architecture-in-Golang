package handlers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/ddd"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/registry/serdes"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/internal/application"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/internal/application/commands"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/internal/domain"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores/storespb"
)

func TestStoresProducer(t *testing.T) {
	var err error

	stores := domain.NewFakeStoreRepository()
	products := domain.NewFakeProductRepository()
	mall := domain.NewFakeMallRepository()
	catalog := domain.NewFakeCatalogRepository()

	type rawEvent struct {
		Name    string
		Payload json.RawMessage
	}

	reg := registry.New()
	err = storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	verifier := message.Verifier{}
	err = verifier.Verify(t, message.VerifyMessageRequest{
		VerifyRequest: provider.VerifyRequest{
			Provider:                   "stores-pub",
			ProviderVersion:            "1.0.0",
			BrokerURL:                  "http://127.0.0.1:9292",
			BrokerUsername:             "pactuser",
			BrokerPassword:             "pactpass",
			PublishVerificationResults: true,
			AfterEach: func() error {
				stores.Reset()
				products.Reset()
				return nil
			},
		},
		MessageHandlers: map[string]message.Handler{
			"a StoreCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				// Assign
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.New(stores, products, catalog, mall, dispatcher)
				publisher := am.NewFakeMessagePublisher[ddd.Event]()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				// Act
				err := app.CreateStore(context.Background(), commands.CreateStore{
					ID:       "store-id",
					Name:     "NewStore",
					Location: "NewLocation",
				})
				if err != nil {
					return nil, nil, err
				}

				// Assert
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
			"a StoreRebranded message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.New(stores, products, catalog, mall, dispatcher)
				publisher := am.NewFakeMessagePublisher[ddd.Event]()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				store := domain.NewStore("store-id")
				store.Name = "NewStore"
				store.Location = "NewLocation"
				stores.Reset(store)

				err := app.RebrandStore(context.Background(), commands.RebrandStore{
					ID:   "store-id",
					Name: "RebrandedStore",
				})
				if err != nil {
					return nil, nil, err
				}

				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
