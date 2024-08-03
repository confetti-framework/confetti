package command

import (
	"errors"
	"fmt"
	"net/http"
	"src/config"
	"src/internal/ping"
	"src/internal/pkg/handler"
)

type AppServe struct {
}

func (s AppServe) Name() string {
	return "api:serve"
}

func (s AppServe) Description() string {
	return "Start the http server to handle requests"
}

var ApiRoutes = []handler.Route{
	handler.New("GET /ping", ping.Index),
}

func (s AppServe) Handle() error {
	fmt.Printf("\u001B[32mStarting server:\u001B[0m %s\n", s.getListenAddr())

	// Register the routes
	mux := http.NewServeMux()

	registrar := handler.RegisterRoutes(handler.AppendApiByPath(ApiRoutes), handler.HandleApiRoute)
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
