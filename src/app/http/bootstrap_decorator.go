package http

import (
	"lanvard/foundation"
	"lanvard/foundation/bootstrap"
	decorators "lanvard/foundation/decorator"
	contract "lanvard/interface/decorator"
)

var bootstrappers = []contract.Bootstrap{
	bootstrap.LoadEnvironmentVariables(),
}

func Bootstrap(app foundation.Application) foundation.Application {

	decorator := decorators.BootstrapDecoratorStruct{}

	return decorator.BootstrapWith(app, bootstrappers)
}
