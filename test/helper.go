package test

import (
	"github.com/lanvard/contract/inter"
	foundation "github.com/lanvard/foundation/http"
	"lanvard/bootstrap"
	net "net/http"
	"net/http/httptest"
)

func ResponseByRequest(request inter.Request) inter.Response {

	app := request.App()
	if nil == request.App() {
		app = bootstrap.NewAppFromBoot()
	}

	response := httptest.ResponseRecorder{}

	app.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)
	request = request.SetApp(app)

	kernel := app.Make((*inter.HttpKernel)(nil)).(foundation.Kernel)

	return kernel.Handle(request)
}
