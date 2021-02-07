package main

import (
	"confetti-framework/bootstrap"
	"confetti-framework/config"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"log"
	net "net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("Start " + config.App.Name + " to handle requests")
	server := &net.Server{
		Addr:         ":" + strconv.Itoa(config.App.Port),
		Handler:      net.HandlerFunc(HandleKernel),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != net.ErrServerClosed {
		log.Fatal("Could not ", err)
	}

	log.Println("Server stopped")
}

func HandleKernel(response net.ResponseWriter, request *net.Request) {
	/*
	   |--------------------------------------------------------------------------
	   | Turn On The Lights
	   |--------------------------------------------------------------------------
	   |
	   | We need to illuminate Go development, so let us turn on the lights.
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

	appRequest := http.NewRequest(http.Options{App: app, Source: *request})

	defer func() {
		if rec := recover(); rec != nil {
			appResponse := kernel.RecoverFromMiddlewarePanic(rec)
			exposeResponse(response, appResponse)
		}
	}()

	appResponse := kernel.Handle(appRequest)

	exposeResponse(response, appResponse)
}

func exposeResponse(response net.ResponseWriter, appResponse inter.Response) {
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
