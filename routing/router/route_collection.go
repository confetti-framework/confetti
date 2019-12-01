package router

import (
	"github.com/gorilla/mux"
	"lanvard/http"
)

type MapUrlRoutes = map[string]Route
type MapMethodRoutes map[string]MapUrlRoutes

type RouteCollection struct {
	// An array of the routes keyed by method.
	routesMapRoutes MapMethodRoutes
	routes          []Route
}

func NewRouteCollection() RouteCollection {
	return RouteCollection{}
}

func (c *RouteCollection) Push(route Route) {
	// Todo add domain
	domainAndUri := route.Uri
	c.routesMapRoutes[route.Method][domainAndUri] = route
	c.routes = append(c.routes, route)
}

func (c *RouteCollection) Merge(routeCollections []RouteCollection) {
	for _, routeCollection := range routeCollections {
		for _, route := range routeCollection.All() {
			c.Push(route)
		}
	}
}

func (c RouteCollection) All() []Route {
	return c.routes
}

func (c RouteCollection) Match(request http.Request) Route {
	routesWithUrl := c.getByMethod(request.GetMethod())

	// First, we will see if we can find a matching route for this current request
	// method. If we can, great, we can just return it so that it can be called
	// by the consumer. Otherwise we will check for routes with another verb.
	route, found := c.matchAgainstRoutes(routesWithUrl, request)

	if found {
		return route
	}

	// If no route was found we will now check if a matching route is specified by
	// another HTTP verb. If it is we will need to throw a MethodNotAllowed and
	// inform the user agent of which HTTP verb it should use for this route.
	ok := c.hasAlternateMethod(request)

	if ok {
		// todo implement error handling with MethodNotAllowedHttp error
		panic("The %s method is not supported for this route. Supported methods: %s.")
	}

	// todo implement error handling
	panic("throw NotFoundHttp error")
}

func (c RouteCollection) getByMethod(method string) MapUrlRoutes {
	return c.routesMapRoutes[method]
}

func (c *RouteCollection) matchAgainstRoutes(routes MapUrlRoutes, request http.Request) (Route, bool) {

	for uri, route := range routes {
		var match mux.RouteMatch
		ok := new(mux.Route).Path(uri).Match(&request.Original, &match)
		if ok {
			return route, true
		}
	}

	var nil Route
	return nil, false
}

func (c RouteCollection) hasAlternateMethod(request http.Request) bool {

	for _, route := range c.routes {
		var match mux.RouteMatch
		ok := new(mux.Route).Path(route.Uri).Match(&request.Original, &match)
		if ok {
			return true
		}
	}

	return false
}
