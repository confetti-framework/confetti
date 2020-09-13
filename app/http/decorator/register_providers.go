package decorator

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type RegisterProviders struct{}

// Providers are located in config/providers/providers.go
func (r RegisterProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range config.Providers.RegisterProviders {
		container = bootstrapper.Register(container)
	}

	return container
}
