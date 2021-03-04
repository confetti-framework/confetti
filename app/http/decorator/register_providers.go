package decorator

import (
	"confetti-framework/app/providers"
	"github.com/confetti-framework/contract/inter"
)

type RegisterProviders struct{}

func (r RegisterProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		container = bootstrapper.Register(container)
	}

	return container
}
