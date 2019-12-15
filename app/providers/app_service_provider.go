package providers

import (
	"github.com/lanvard/foundation"
)

type AppServiceProvider struct{}

// Register any application services.
func (a AppServiceProvider) Register(app foundation.Application) foundation.Application {

	return app
}

// Bootstrap any application services.
func (a AppServiceProvider) Boot(app foundation.Application) foundation.Application {

	return app
}
