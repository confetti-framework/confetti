package controllers

import (
	"lanvard/http"
	"lanvard/routing"
)

var User = struct {
	Index   routing.ControllerMethod
	Store   routing.ControllerMethod
	Destroy routing.ControllerMethod
	Show    routing.ControllerMethod
}{
	Index: func(request http.Request) http.Response {

		return http.Json("{\"test all users\"}")
	},

	Show: func(request http.Request) http.Response {
		return http.Json("{\"test show one user\"}")
	},

	Store: func(request http.Request) http.Response {
		return http.Json("{\"test store use\"}")
	},

	Destroy: func(request http.Request) http.Response {
		return http.Json("{\"test remove use\"}")
	},
}
