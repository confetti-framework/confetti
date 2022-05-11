package providers

import (
	foundationMiddleware "github.com/confetti-framework/framework/foundation/http/middleware"
	"github.com/confetti-framework/framework/foundation/http/routing"
	"github.com/confetti-framework/framework/inter"
	"src/app/http/middleware"
	"src/routes"
)

var globalMiddlewares = []inter.HttpMiddleware{
	foundationMiddleware.RequestID{},
	foundationMiddleware.RequestBodyDecoder{},
	middleware.RouteModelBinding{},
}

// RouteServiceProvider is responsible for generate the routes
type RouteServiceProvider struct{}

// Boot defined your router model bindings, pattern filters, etc.
func (r RouteServiceProvider) Boot(container inter.Container) inter.Container {
	collection := routing.NewRouteCollection()

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	// Here you can make your adjustments that apply to all routes:
	// collection.WhereMulti(map[string]string{
	// 	"id": "[0-9]+",
	// })

	// These middlewares should be executed on all routes
	collection.Middleware(globalMiddlewares...)

	routing.DecorateRoutes(collection)
	container.Singleton("routes", collection)

	return container
}
