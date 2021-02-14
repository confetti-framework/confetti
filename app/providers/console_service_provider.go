package providers

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/flag_type"
	"reflect"
)

type ConsoleServiceProvider struct{}

func (c ConsoleServiceProvider) Register(container inter.Container) inter.Container {
	// Flag types.
	// These types are used to convert the flags from
	// the commands to a specific value in a command field.
	container.Singleton((*map[interface{}]flag.Value)(nil), map[interface{}]flag.Value{
		reflect.String: new(flag_type.String),
	})

	return container
}
