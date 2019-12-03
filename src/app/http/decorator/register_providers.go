package decorator

import (
	"lanvard/config"
	"lanvard/foundation"
)

type RegisterProviders struct{}

func (r RegisterProviders) Bootstrap(app foundation.Application) foundation.Application {
	for _, bootstrapper := range config.App.RegisterProviders {
		app = bootstrapper.Register(app)
	}

	return app
}
