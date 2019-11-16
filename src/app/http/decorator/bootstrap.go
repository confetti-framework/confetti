package decorator

import (
	"lanvard/foundation"
	"lanvard/foundation/bootstrap"
	"lanvard/foundation/decorator"
	contract "lanvard/interface/decorator"
)

var bootstraps = []contract.Bootstrap{
	bootstrap.ProfessorGenerator(),
	bootstrap.LoadEnvironmentVariables(),
}

func Bootstrap(app foundation.Application) foundation.Application {
	dec := decorator.BootstrapDecoratorStruct{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
