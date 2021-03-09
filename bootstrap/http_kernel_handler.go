package bootstrap

import (
	"github.com/confetti-framework/foundation/http"
	net "net/http"
)

// HandleHttpKernel turns On The Lights. We need to illuminate Go development,
// so let us turn on the lights. This bootstraps the framework and gets it ready
// for use, then it will load up this application so that we can run it and send
// the responses back to the browser and delight our users.
func HandleHttpKernel(response net.ResponseWriter, request *net.Request) {
	app := NewAppFromBoot()
	http.HandleHttpKernel(app, response, request)
}
