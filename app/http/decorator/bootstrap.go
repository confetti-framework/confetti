package decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/decorator"
)

var bootstraps = []inter.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(container *foundation.Container) inter.Container {
	bootstrapDecorator := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return bootstrapDecorator.BootstrapWith(container)
}
