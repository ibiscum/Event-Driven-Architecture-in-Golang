package handlers

import (
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/cosec/internal"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/cosec/internal/models"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/am"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/registry"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/sec"
)

func NewReplyHandlers(reg registry.Registry, orchestrator sec.Orchestrator[*models.CreateOrderData], mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewReplyHandler(reg, orchestrator, mws...)
}

func RegisterReplyHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(internal.CreateOrderReplyChannel, handlers, am.GroupName("cosec-replies"))
	return err
}
