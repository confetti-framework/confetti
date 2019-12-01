package providers

import (
	"lanvard/foundation"
	"lanvard/routing/router"
	"lanvard/src/routes"
)

type RouteServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (p RouteServiceProvider) Boot(app foundation.Application) foundation.Application {
	collection := router.NewRouteCollection()

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	app.Container.Instance("routes", collection)

	return app
}
