package controllers

import (
	"lanvard/http"
	"lanvard/routing"
)

var User = struct {
	Index   routing.ControllerMethod
	Store   routing.ControllerMethod
	Destroy routing.ControllerMethod
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
