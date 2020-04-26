package decorator

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config/providers"
)

type RegisterProviders struct{}

// Providers are located in config/providers/providers.go
func (r RegisterProviders) Bootstrap(app inter.App) inter.App {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		app = bootstrapper.Register(app)
	}

	return app
}
