package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"lanvard/src/adapter"
)

var User = struct {
	Index, Create, Store, Show, Edit, Update, Destroy inter.Controller
}{
	Index: func(request inter.Request) inter.Response {
		users, err := adapter.User{request}.AllE()
		if err != nil {
			return outcome.Json(err)
		}

		return outcome.Json(users)
	},

	Show: func(request inter.Request) inter.Response {
		user, err := adapter.User{request}.FindE()
		if err != nil {
			return outcome.Json(err)
		}

		return outcome.Json(user)
	},
}
