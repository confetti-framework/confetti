package router

import (
	"lanvard/foundation"
	"lanvard/http"
)

type Router struct {
	app            foundation.Application
	routes         RouteCollection
	currentRequest http.Request
}

func NewRouter(app foundation.Application) Router {
	routes := app.Make("routes").(RouteCollection)
	return Router{app: app, routes: routes}
}

func (r Router) DispatchToRoute(request http.Request) http.Response {
	r.currentRequest = request

	route := r.routes.Match(request)

	// todo implement event Events\RouteMatched
	// todo implement RouteMiddleware
	// todo remove nil
	return route.Controller(nil, request)
}
