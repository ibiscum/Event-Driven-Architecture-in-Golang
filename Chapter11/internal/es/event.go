package es

import (
	"fmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter11/internal/ddd"
)

type EventApplier interface {
	ApplyEvent(ddd.Event) error
}

type EventCommitter interface {
	CommitEvents()
}

func LoadEvent(v interface{}, event ddd.AggregateEvent) error {
	type loader interface {
		EventApplier
		VersionSetter
	}

	agg, ok := v.(loader)
	if !ok {
		return fmt.Errorf("%T does not have the methods implemented to load events", v)
	}

	if err := agg.ApplyEvent(event); err != nil {
		return err
	}
	agg.SetVersion(event.AggregateVersion())

	return nil
}
