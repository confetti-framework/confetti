package controllers

import (
	"github.com/lanvard/foundation/http/lanvard"
	"github.com/lanvard/routing"
)

var User = struct {
	Index   routing.ControllerMethod
	Store   routing.ControllerMethod
	Destroy routing.ControllerMethod
	Show    routing.ControllerMethod
}{
	Index: func(request lanvard.Request) lanvard.Response {

		return lanvard.Json("{\"test all users\"}")
	},

	Show: func(request lanvard.Request) lanvard.Response {
		return lanvard.Json("{\"test show one user\"}")
	},

	Store: func(request lanvard.Request) lanvard.Response {
		return lanvard.Json("{\"test store use\"}")
	},

	Destroy: func(request lanvard.Request) lanvard.Response {
		return lanvard.Json("{\"test remove use\"}")
	},
}
