package main

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"confetti-framework/bootstrap"
	"log"
	net "net/http"
	"strings"
	"time"
)

func main() {
	log.Println("Server is ready to handle requests")
	server := &net.Server{
		Addr:         ":80",
		Handler:      net.HandlerFunc(HandleKernel),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != net.ErrServerClosed {
		log.Fatal("Could not listen. You may want to kill build process by `sudo -S pkill -SIGINT ___go_bui`", err)
	}

	log.Println("Server stopped")
}

func HandleKernel(response net.ResponseWriter, request *net.Request) {

	/*
	   |--------------------------------------------------------------------------
	   | Turn On The Lights
	   |--------------------------------------------------------------------------
	   |
	   | We need to illuminate PHP development, so let us turn on the lights.
	   | This bootstraps the framework and gets it ready for use, then it
	   | will load up this application so that we can run it and send
	   | the responses back to the browser and delight our users.
	   |
	*/
	app := bootstrap.NewAppFromBoot()

	/*
	   |--------------------------------------------------------------------------
	   | Register the response writer
	   |--------------------------------------------------------------------------
	   |
	   | Confetti only uses the response writer here in main.go. But we register
	   | the response writer if you need it anyway
	   |
	*/
	app.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the incoming request
	   | through the kernel, and send the associated response back to
	   | the client allowing them to enjoy the creative
	   | and awesome application we have prepared for them.
	   |
	*/
	kernel := app.Make((*inter.HttpKernel)(nil)).(http.Kernel)
	appResponse := kernel.Handle(http.NewRequest(http.Options{App: app, Source: *request}))

	// Add HTTP headers
	for key, values := range appResponse.GetHeaders() {
		response.Header().Add(key, strings.Join(values, "; "))
	}

	// Add HTTP status
	response.WriteHeader(appResponse.GetStatus())

	// Add HTTP body
	_, err := response.Write([]byte(appResponse.GetBody()))
	if err != nil {
		panic(err)
	}
}
