package es

import (
	"fmt"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter09/internal/registry"
)

type VersionSetter interface {
	setVersion(int)
}

func SetVersion(version int) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(VersionSetter); ok {
			agg.setVersion(version)
			return nil
		}
		return fmt.Errorf("%T does not have the method setVersion(int)", v)
	}
}
