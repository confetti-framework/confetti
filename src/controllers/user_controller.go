package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/outcome"
	"lanvard/src/adapters"
)

var User = struct {
	Index, Create, Store, Show, Edit, Update, Destroy inter.ControllerMethod
}{
	Index: func(request inter.Request) inter.Response {
		users, err := adapters.User.AllByRequest(request)
		if err != nil {
			return outcome.Error(err)
		}

		return outcome.Json(users)
	},

	Show: func(request inter.Request) inter.Response {
		user, err := adapters.User.OneByRequest(request)
		if err != nil {
			return outcome.Error(err)
		}

		return outcome.Json(user)
	},
}
