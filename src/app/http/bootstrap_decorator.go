package http

import (
	"laravelgo/foundation/bootstrap"
	decorators "laravelgo/foundation/decorator"
	"laravelgo/interface/application"
	contract "laravelgo/interface/decorator"
)

var bootstrappers = []contract.Bootstrap{
	bootstrap.LoadEnvironmentVariables(),
}

func Bootstrap(app application.App) application.App {
	decorator := decorators.BootstrapDecoratorStruct{}

	return decorator.BootstrapWith(app, bootstrappers)
}
