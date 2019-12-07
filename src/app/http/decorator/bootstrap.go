package decorator

import (
	contract "lanvard/contract/decorator"
	"lanvard/foundation"
	"lanvard/foundation/bootstrap"
	"lanvard/foundation/decorator"
)

var bootstraps = []contract.Bootstrap{
	bootstrap.LoadEnvironmentVariables{},
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(app foundation.Application) foundation.Application {
	dec := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
