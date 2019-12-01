package router

type Route struct {
	Uri        string
	Method     string
	Controller ControllerMethod
}

func NewRoute(uri string, method string, controller ControllerMethod) Route {
	return Route{Uri: uri, Method: method, Controller: controller}
}
