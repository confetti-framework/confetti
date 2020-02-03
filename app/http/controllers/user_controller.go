package controllers

import (
	"github.com/lanvard/http"
	"github.com/lanvard/routing"
)

var User = struct {
	Index   routing.Controller
	Store   routing.Controller
	Destroy routing.Controller
	Show    routing.Controller
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
