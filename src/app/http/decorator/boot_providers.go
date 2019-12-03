package decorator

import (
	"lanvard/config"
	"lanvard/foundation"
)

type BootProviders struct{}

func (r BootProviders) Bootstrap(app foundation.Application) foundation.Application {
	for _, bootstrapper := range config.App.BootProviders {
		app = bootstrapper.Boot(app)
	}

	return app
}
