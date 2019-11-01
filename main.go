package lanvard

import (
	"fmt"
	. "laravelgo/interface/http"
	"laravelgo/src/bootstrap"
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

func handleKernel(w net.ResponseWriter, request *net.Request) {

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

	/*
		|--------------------------------------------------------------------------
		| Run The Application
		|--------------------------------------------------------------------------
		|
		| Once we have the application, we can handle the incoming request
		| through the kernel, and send the associated response back to
		| the client's browser allowing them to enjoy the creative
		| and wonderful application we have prepared for them.
		|
	*/

	fmt.Println(request.URL.Query())
	kernel := app.Container.Make((*Kernel)(nil)).(Kernel)

	response := kernel.Handle(request)

	w.WriteHeader(net.StatusOK)
	_, _ = fmt.Fprintf(w, response)
}
