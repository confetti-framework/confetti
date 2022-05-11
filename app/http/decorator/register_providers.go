package decorator

import (
	"github.com/confetti-framework/framework/inter"
	"src/app/providers"
)

type registerProviders struct{}

// Bootstrap handles all registerProviders. Providers are configured in app/providers/provider_index.go
func (r registerProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		container = bootstrapper.Register(container)
	}

	return container
}
