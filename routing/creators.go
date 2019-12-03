package routing

func createRoutes(methods []string, uri string, controller ControllerMethod) RouteCollection {
	routes := NewRouteCollection()
	for _, method := range methods {
		route := NewRoute(uri, method, controller)
		routes.Push(route)
	}

	return routes
}

func createRoute(method string, uri string, controller ControllerMethod) RouteCollection {
	methods := []string{method}
	return createRoutes(methods, uri, controller)
}
