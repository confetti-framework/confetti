package providers

import (
	"github.com/confetti-framework/framework/inter"
)

// AppServiceProvider contains all providers
type AppServiceProvider struct{}

// Register any container services.
func (a AppServiceProvider) Register(container inter.Container) inter.Container {

	//

	return container
}

// Boot any container services.
func (a AppServiceProvider) Boot(container inter.Container) inter.Container {

	//

	return container
}
