package providers

import (
	"github.com/lanvard/contract/inter"
)

type AppServiceProvider struct{}

// Register any application services.
func (a AppServiceProvider) Register(app inter.App) inter.App {

	//

	return app
}

// Bootstrap any application services.
func (a AppServiceProvider) Boot(app inter.App) inter.App {

	//

	return app
}
