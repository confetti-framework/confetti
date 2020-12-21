package providers

import (
	"github.com/confetti-framework/contract/inter"
	"confetti-framework/config"
)

type EnvServiceProvider struct{}

// Register any container services.
func (a EnvServiceProvider) Register(container inter.Container) inter.Container {
	container.Bind("env", config.App.Env)

	return container
}
