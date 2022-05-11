package decorator

import (
	"github.com/confetti-framework/framework/inter"
	"src/app/providers"
)

type bootProviders struct{}

// Bootstrap handles all bootProviders. Providers are configured in app/providers/provider_index.go
func (r bootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
