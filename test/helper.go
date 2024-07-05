package test

import (
    "io"
    "net/http"
    "src/app/http/entities"
)

func GetBody(result *http.Response) string {
    defer result.Body.Close()
    body, err := io.ReadAll(result.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}

func GetRoute(routes []entities.Route, pattern string) entities.Route {
    for _, route := range routes {
        if route.Pattern == pattern {
            return route
        }
    }
    panic(`Route with pattern ` + pattern + ` not found`)
}
