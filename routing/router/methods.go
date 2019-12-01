package router

import (
	"lanvard/http"
	net "net/http"
)

type ControllerMethod func(interface{}, http.Request) http.Response

// All of the methods supported by the router.
var allMethods = []string{
	net.MethodGet,
	net.MethodHead,
	net.MethodPost,
	net.MethodPut,
	net.MethodPatch,
	net.MethodDelete,
	net.MethodOptions,
}

// Register new GET routes
func Get(uri string, controller ControllerMethod) RouteCollection {
	methods := []string{net.MethodGet, net.MethodHead}
	return createRoutes(methods, uri, controller)
}

// Register new POST routes
func Post(uri string, controller ControllerMethod) RouteCollection {
	return createRoute(net.MethodPost, uri, controller)
}

// Register new PUT routes
func Put(uri string, controller ControllerMethod) RouteCollection {
	return createRoute(net.MethodPut, uri, controller)
}

// Register new PATCH routes
func Patch(uri string, controller ControllerMethod) RouteCollection {
	return createRoute(net.MethodPatch, uri, controller)
}

// Register new DELETE routes
func Delete(uri string, controller ControllerMethod) RouteCollection {
	return createRoute(net.MethodDelete, uri, controller)
}

// Register new OPTIONS routes
func Options(uri string, controller ControllerMethod) RouteCollection {
	return createRoute(net.MethodOptions, uri, controller)
}

// Register a new route responding to all methods
func Any(uri string, controller ControllerMethod) RouteCollection {
	return createRoutes(allMethods, uri, controller)
}
