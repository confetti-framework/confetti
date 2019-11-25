package decorator

import (
	"lanvard/foundation"
	"lanvard/foundation/bootstrap"
	"lanvard/foundation/decorator"
	contract "lanvard/interface/decorator"
)

var bootstraps = []contract.Bootstrap{
	bootstrap.BootLoadEnvironmentVariables(),
}

func Bootstrap(app foundation.Application) foundation.Application {
	dec := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
