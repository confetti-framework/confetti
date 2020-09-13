package providers

import (
	"github.com/lanvard/contract/inter"
)

type AppServiceProvider struct{}

// Register any container services.
func (a AppServiceProvider) Register(container inter.Container) inter.Container {

	//

	return container
}

// Bootstrap any container services.
func (a AppServiceProvider) Boot(container inter.Container) inter.Container {

	//

	return container
}
