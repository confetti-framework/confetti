package providers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing"
	"lanvard/routes"
)

type RouteServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (p RouteServiceProvider) Boot(container inter.Container) inter.Container {
	collection := routing.NewRouteCollection()
	collection.SetContainer(container)

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	// Here you can make your adjustments that apply to all routes:
	// collection.WhereMulti(map[string]string{
	// 	"id": "[0-9]+",
	// })

	routing.DecorateRoutes(collection)

	container.Singleton("routes", collection)

	return container
}
