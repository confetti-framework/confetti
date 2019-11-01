package decorator

import (
	"laravelgo/interface/application"
	"laravelgo/interface/decorator"
)

type BootstrapDecoratorStruct struct {}

func (d BootstrapDecoratorStruct) BootstrapWith(
	app application.App,
	bootstrappers []decorator.Bootstrap,
) application.App {

	return app
}
