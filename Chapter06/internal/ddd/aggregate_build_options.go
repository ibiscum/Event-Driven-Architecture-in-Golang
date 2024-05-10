package ddd

import (
	"fmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter06/internal/registry"
)

type EventsSetter interface {
	setEvents([]Event)
}

func SetEvents(events ...Event) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(EventsSetter); ok {
			agg.setEvents(events)
			return nil
		}
		return fmt.Errorf("%T does not have the method setEvents([]ddd.Event)", v)
	}
}
