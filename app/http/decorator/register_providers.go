package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"src/app/providers"
)

type RegisterProviders struct{}

func (r RegisterProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		container = bootstrapper.Register(container)
	}

	return container
}
