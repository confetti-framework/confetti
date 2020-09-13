package decorator

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type BootProviders struct{}

// Providers are located in config/providers/providers.go
func (r BootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range config.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
