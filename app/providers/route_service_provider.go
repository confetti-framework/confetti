package providers

import (
	"github.com/lanvard/foundation"
	"lanvard/routes"
	"lanvard/routing"
)

type RouteServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (p RouteServiceProvider) Boot(app foundation.Application) foundation.Application {
	collection := routing.NewRouteCollection()

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	app.Container.Singleton("routes", collection)

	return app
}
