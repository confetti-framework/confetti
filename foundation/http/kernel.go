package http

import (
	"laravelgo/foundation"
	"net/http"
)

type KernelStruct struct {
	App foundation.ApplicationStruct
}

// Handle an incoming HTTP request.
func (k KernelStruct) Handle(request http.Request) string {
	return k.sendRequestThroughRouter(request)
}

// Send the given request through the middleware / router.
func (k KernelStruct) sendRequestThroughRouter(request http.Request) string {
	k.App.Container.Instance("request", request)

	k.Bootstrap()

	return "response"
}

func (k KernelStruct) Bootstrap() {
	if ! k.App.HasBeenBootstrapped {
		k.App.HasBeenBootstrapped = true

	}
}
