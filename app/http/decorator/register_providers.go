package decorator

import (
	"github.com/lanvard/foundation"
	"lanvard/config/providers"
)

type RegisterProviders struct{}

// Providers are located in config/providers/providers.go
func (r RegisterProviders) Bootstrap(app foundation.Application) foundation.Application {
	for _, bootstrapper := range providers.Providers.RegisterProviders {
		app = bootstrapper.Register(app)
	}

	return app
}
