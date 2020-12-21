package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/decorator"
)

var bootstraps = []inter.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(container *foundation.Container) inter.Container {
	bootstrapDecorator := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return bootstrapDecorator.BootstrapWith(container)
}
