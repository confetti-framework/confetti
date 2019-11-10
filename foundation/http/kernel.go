package http

import (
	"lanvard/foundation"
	"lanvard/src/app/http/decorator"
	"net/http"
)

type KernelStruct struct {
	App foundation.Application
}

// Handle an incoming HTTP request.
func (k KernelStruct) Handle(request http.Request) string {
	return k.sendRequestThroughRouter(request)
}

// Send the given request through the middleware / router.
func (k KernelStruct) sendRequestThroughRouter(request http.Request) string {
	k.App.Container.Instance("request", request)

	return "response"
}

func (k KernelStruct) Bootstrap() KernelStruct {
	decorator.Bootstrap(k.App)

	if !k.App.HasBeenBootstrapped {
		k.App.HasBeenBootstrapped = true
	}

	return k
}
