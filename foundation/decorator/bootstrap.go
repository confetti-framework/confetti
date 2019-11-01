package decorator

import (
	"lanvard/foundation"
	"lanvard/interface/decorator"
)

type BootstrapDecoratorStruct struct{}

func (d BootstrapDecoratorStruct) BootstrapWith(
	app foundation.Application,
	bootstrappers []decorator.Bootstrap,
) foundation.Application {

	for _, bootstrapper := range bootstrappers {
		app = bootstrapper.Bootstrap(app)
	}

	return app
}
