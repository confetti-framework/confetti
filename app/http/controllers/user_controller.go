package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/outcome"
)

var User = struct {
	Index, Store, Destroy, Show inter.ControllerMethod
}{

	Index: func(request inter.Request) inter.Response {
		return outcome.Json("{\"test all users\"}")
	},

	Show: func(request inter.Request) inter.Response {
		return outcome.Json("{\"test show one user\"}")
	},

	Store: func(request inter.Request) inter.Response {
		return outcome.Json("{\"test store use\"}")
	},

	Destroy: func(request inter.Request) inter.Response {
		return outcome.Json("{\"test remove use\"}")
	},
}
