package console

import (
    "errors"
    "fmt"
    "net/http"
    "src/app/config"
    "src/app/http/entities"
    "src/app/http/routes"
)

type AppServe struct {
}

func (s AppServe) Name() string {
    return "app:serve"
}

func (s AppServe) Description() string {
    return "Start the http server to handle requests"
}

func (s AppServe) Handle() error {
    fmt.Printf("\u001B[32mStarting server:\u001B[0m %s\n", s.getListenAddr())

    // Register the routes
    mux := http.NewServeMux()
    registrar := registerRoutes(routes.Api, routes.HandleApiRoute)
    registrar(mux)

    // Create the server
    server := http.Server{
        Addr:         s.getListenAddr(),
        WriteTimeout: config.AppServe.Timeout,
        ReadTimeout:  config.AppServe.Timeout,
        IdleTimeout:  config.AppServe.Timeout,
        Handler:      mux,
    }

    // Listen to all the routes
    err := server.ListenAndServe()

    // Do not report when user closed the server
    if errors.Is(err, http.ErrServerClosed) {
        return nil
    }
    return fmt.Errorf("could not %v", err)
}

func (s AppServe) getListenAddr() string {
    return fmt.Sprintf("%s:%d", config.AppServe.Host, config.AppServe.Port)
}

func registerRoutes(routes []entities.Route, routeHandler func(response http.ResponseWriter, request *http.Request, route entities.Route)) func(mux *http.ServeMux) {
    return func(mux *http.ServeMux) {
        for _, route := range routes {
            mux.HandleFunc(route.Pattern, func(response http.ResponseWriter, request *http.Request) {
                routeHandler(response, request, route)
            })
        }
    }
}
