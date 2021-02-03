package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/decorator/container_decorator"
)

var bootstraps = []inter.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(container *foundation.Container) inter.Container {
	handler := container_decorator.Handler{Bootstraps: bootstraps}

	return handler.BootstrapWith(container)
}
