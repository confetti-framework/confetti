package entity

import (
    "net/http"
)

type Controller func(w http.ResponseWriter, req *http.Request) error

type Route struct {
    Pattern     string
    Controller  Controller
    middlewares []Middleware
}

func (r Route) AppendMiddleware(middlewares ...Middleware) Route {
    r.middlewares = append(r.middlewares, middlewares...)
    return r
}

func (r Route) GetMiddlewares() []Middleware {
    return r.middlewares
}
