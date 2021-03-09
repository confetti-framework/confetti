package providers

import (
	"github.com/confetti-framework/contract/inter"
	"src/config"
)

// EnvServiceProvider is responsible for binding all environments, in case you would need them later.
type EnvServiceProvider struct{}

// Register is responsible for binding all environments, in case you would need them later.
func (a EnvServiceProvider) Register(container inter.Container) inter.Container {
	container.Bind("env", config.App.Env)

	return container
}
