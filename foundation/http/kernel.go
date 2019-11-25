package http

import (
	"lanvard/foundation"
	"lanvard/http"
	"lanvard/src/app/http/decorator"
	"lanvard/src/app/http/middleware"
)

type KernelStruct struct {
	App        foundation.Application
	Middleware []middleware.PipeInterface
}

// Handle an incoming HTTP request.
func (k KernelStruct) Handle(request http.RequestStruct) http.ResponseStruct {
	return k.sendRequestThroughRouter(request)
	// @todo event RequestHandled
}

// Send the given request through the middleware / router.
func (k KernelStruct) sendRequestThroughRouter(request http.RequestStruct) http.ResponseStruct {
	k.App.Container.Instance("request", request)

	return middleware.Pipeline(k.App).
		Send(request).
		Through(k.Middleware).
		Then(k.dispatchToRouter())
}

func (k KernelStruct) Bootstrap() foundation.Application {
	k.App = decorator.Bootstrap(k.App)
	k.App.HasBeenBootstrapped = true

	return k.App
}

func (k KernelStruct) dispatchToRouter() func(request http.RequestStruct) http.ResponseStruct {

	return func(r http.RequestStruct) http.ResponseStruct {
		return http.Response().SetContent("ResponseTest")
	}
}
