package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"confetti-framework/app/providers"
)

type RegisterProviders struct{}

// Providers are located in config/providers/providers.go
func (r RegisterProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		container = bootstrapper.Register(container)
	}

	return container
}
