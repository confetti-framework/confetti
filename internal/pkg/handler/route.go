package handler

import (
	"net/http"
)

type Controller func(w http.ResponseWriter, req *http.Request) error

type Route struct {
	Pattern     string
	Controller  Controller
	Middlewares []Middleware
}

func (r Route) AppendMiddleware(middlewares ...Middleware) Route {
	r.Middlewares = append(r.Middlewares, middlewares...)
	return r
}

func (r Route) GetMiddlewares() []Middleware {
	return r.Middlewares
}
