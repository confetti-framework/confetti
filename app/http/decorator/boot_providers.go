package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"src/app/providers"
)

type bootProviders struct{}

// Bootstrap handles all bootProviders. Providers are located in config/providers/providers.go
func (r bootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
