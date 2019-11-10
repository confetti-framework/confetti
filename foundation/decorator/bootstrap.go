package decorator

import (
	"lanvard/foundation"
	"lanvard/interface/decorator"
)

type BootstrapDecoratorStruct struct {
	Bootstraps []decorator.Bootstrap
}

func (d BootstrapDecoratorStruct) BootstrapWith(app foundation.Application) foundation.Application {
	for _, bootstrapper := range d.Bootstraps {
		app = bootstrapper.Bootstrap(app)
	}

	return app
}
