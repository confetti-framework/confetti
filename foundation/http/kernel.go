package http

import (
	"lanvard/foundation"
	"lanvard/src/app/http/decorator"
	"lanvard/src/app/http/pipeline"
	"net/http"
	"net/http/httptest"
)

type KernelStruct struct {
	App        foundation.Application
	Middleware []pipeline.PipeInterface
}

// Handle an incoming HTTP request.
func (k KernelStruct) Handle(request http.Request) http.ResponseWriter {
	return k.sendRequestThroughRouter(request)
	// @todo event RequestHandled
}

// Send the given request through the middleware / router.
func (k KernelStruct) sendRequestThroughRouter(request http.Request) http.ResponseWriter {
	k.App.Container.Instance("request", request)

	return httptest.NewRecorder()
	// todo handle response callback type
	// return Pipeline(k.App).
	// 	Send(request).
	// 	Through(k.Middleware).
	// 	Then(k.dispatchToRouter())
}

func (k KernelStruct) Bootstrap() KernelStruct {
	decorator.Bootstrap(k.App)
	k.App.HasBeenBootstrapped = true

	return k
}

func (k KernelStruct) dispatchToRouter() func(data http.Request) http.ResponseWriter {
	return func(data http.Request) http.ResponseWriter {
		response := k.App.Make("response").(http.ResponseWriter)

		_, _ = response.Write([]byte("ResponseTest"))

		// todo handle route
		return response
	}
}
