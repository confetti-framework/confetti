package decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator"
)

var bootstraps = []inter.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(app inter.App) inter.App {
	bootstrapDecorator := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return bootstrapDecorator.BootstrapWith(app)
}
