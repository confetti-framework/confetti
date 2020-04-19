package main

import (
	"github.com/lanvard/contract/inter"
	foundation "github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/lanvard"
	"lanvard/bootstrap"
	"log"
	net "net/http"
)

func main() {
	net.HandleFunc("/", handleKernel)

	log.Println("Server is ready to handle requests")
	if err := net.ListenAndServe(":80", nil); err != nil && err != net.ErrServerClosed {
		log.Fatal("Could not listen", err)
	}

	log.Println("Server stopped")
}

func handleKernel(response net.ResponseWriter, request *net.Request) {

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
	app := bootstrap.NewAppFromBootApp()

	/*
	   |--------------------------------------------------------------------------
	   | Register the response writer
	   |--------------------------------------------------------------------------
	   |
	   | Lanvard only uses the response writer here in main.go. But we register
	   | the response writer if you need it anyway
	   |
	*/
	app.container.Singleton(
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
	   | and wonderful application we have prepared for them.
	   |
	*/
	kernel := app.Make((*inter.HttpKernel)(nil)).(foundation.Kernel)
	appResponse := kernel.Handle(lanvard.NewRequest(&app, *request))

	response.Write([]byte(appResponse.Content()))
	// todo convert custom 'buffer' response to default go response
}
