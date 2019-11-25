package http

import (
	"lanvard/foundation"
	"lanvard/http"
	"lanvard/src/app/http/decorator"
	"lanvard/src/app/http/middleware"
)

type Kernel struct {
	App        foundation.Application
	Middleware []middleware.PipeInterface
}

// Handle an incoming HTTP request.
func (k Kernel) Handle(request http.Request) http.Response {
	return k.sendRequestThroughRouter(request)
	// @todo event RequestHandled
}

// Send the given request through the middleware / router.
func (k Kernel) sendRequestThroughRouter(request http.Request) http.Response {
	k.App.Container.Instance("request", request)

	return middleware.NewPipeline(k.App).
		Send(request).
		Through(k.Middleware).
		Then(k.dispatchToRouter())
}

func (k Kernel) Bootstrap() foundation.Application {
	k.App = decorator.Bootstrap(k.App)
	k.App.HasBeenBootstrapped = true

	return k.App
}

func (k Kernel) dispatchToRouter() func(request http.Request) http.Response {

	return func(r http.Request) http.Response {
		return http.NewResponse().SetContent("ResponseTest")
	}
}
