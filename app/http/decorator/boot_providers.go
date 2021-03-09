package decorator

import (
	"github.com/confetti-framework/contract/inter"
	"src/app/providers"
)

type BootProviders struct{}

// Bootstrap handles all BootProviders. Providers are located in config/providers/providers.go
func (r BootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
