package decorator

import (
	"github.com/lanvard/foundation"
	"lanvard/config/providers"
)

type BootProviders struct{}

// Providers are located in config/providers/providers.go
func (r BootProviders) Bootstrap(app foundation.Application) foundation.Application {
	for _, bootstrapper := range providers.Providers.BootProviders {
		app = bootstrapper.Boot(app)
	}

	return app
}
