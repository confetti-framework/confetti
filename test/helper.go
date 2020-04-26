package test

import (
	"github.com/lanvard/contract/inter"
	foundation "github.com/lanvard/foundation/http"
	"lanvard/bootstrap"
	net "net/http"
	"net/http/httptest"
)

func ResponseByRequest(request inter.Request) inter.Response {
	app := bootstrap.NewAppFromBoot()

	response := httptest.ResponseRecorder{}

	app.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)

	kernel := app.Make((*inter.HttpKernel)(nil)).(foundation.Kernel)

	request.SetApp(app)

	return kernel.Handle(request)
}
