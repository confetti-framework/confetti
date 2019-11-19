package main

import (
	"fmt"
	"lanvard/foundation/http"
	httpInterface "lanvard/interface/http"
	"lanvard/src/bootstrap"
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
	app := bootstrap.App()

	fmt.Println("app.Container.Singleton")
	fmt.Println(len(app.Container.GetBindings()))

	/*
	   |--------------------------------------------------------------------------
	   | Register the response writer
	   |--------------------------------------------------------------------------
	   |
	   | Register the response writer so that we can fill it at another time. For
	   | example, the response writer is provided by the middleware.
	   |
	*/
	app.Container.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the incoming request
	   | through the kernel, and send the associated res back to
	   | the client's browser allowing them to enjoy the creative
	   | and wonderful application we have prepared for them.
	   |
	*/

	kernel := app.Make((*httpInterface.Kernel)(nil)).(http.KernelStruct)

	fmt.Println("In kernel:")
	fmt.Println(len(kernel.App.Container.GetBindings()))
	response = kernel.Handle(*request)

	response.WriteHeader(net.StatusOK)
	_, _ = fmt.Fprint(response)
}
