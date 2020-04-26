package providers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing"
	"lanvard/routes"
)

type RouteServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (p RouteServiceProvider) Boot(app inter.App) inter.App {
	collection := routing.NewRouteCollection()

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	app.Singleton("routes", collection)

	return app
}
