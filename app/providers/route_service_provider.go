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
	collection.SetApp(app)

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	// collection.WhereMulti(map[string]string{
	// 	"id": "[0-9]+",
	// })

	routing.DecorateRoutes(collection)

	app.Singleton("routes", collection)

	return app
}