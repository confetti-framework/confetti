package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"lanvard/src/adapters"
)

var User = struct {
	Index, Create, Store, Show, Edit, Update, Destroy inter.Controller
}{
	Index: func(request inter.Request) inter.Response {
		users, err := adapters.User{request}.AllE()
		if err != nil {
			return outcome.Error(err)
		}

		return outcome.Json(users)
	},

	Show: func(request inter.Request) inter.Response {
		user := adapters.User{request}.Find()
		user, err := adapters.User{request}.FindE()
		if err != nil {
			return outcome.Error(err)
		}

		return outcome.Json(user)
	},
}
