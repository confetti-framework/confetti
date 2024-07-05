package routes

import "net/http"

type Controller func(w http.ResponseWriter, req *http.Request) error

type Route struct {
    Pattern    string
    Controller Controller
}

func newRoute(pattern string, controller Controller) Route {
    return Route{Pattern: pattern, Controller: controller}
}
