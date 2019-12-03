package controllers

import (
	"lanvard/http"
	. "lanvard/routing/router"
)

var User = struct {
	Index   ControllerMethod
	Store   ControllerMethod
	Destroy ControllerMethod
}{
	Index: func(request http.Request) http.Response {
		return http.Json("{\"test Index\": 123}")
	},

	Store: func(request http.Request) http.Response {
		return http.Json("{\"test Store\": 123}")
	},

	Destroy: func(request http.Request) http.Response {
		return http.Json("{\"test Destroy\": 123}")
	},
}
