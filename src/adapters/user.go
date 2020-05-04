package adapters

import (
	"github.com/lanvard/contract/inter"
	"lanvard/app/model"
	"lanvard/src/contract"
)

var User = struct {
	AllByRequest func(inter.Request) ([]contract.User, error)
	OneByRequest func(inter.Request) (contract.User, error)
}{
	AllByRequest: func(request inter.Request) ([]contract.User, error) {
		var err error

		var users []contract.User

		users = append(users, model.NewUser(5435, "Reindert Vetter"))

		return users, err
	},

	OneByRequest: func(request inter.Request) (contract.User, error) {
		var err error
		userId := request.UrlValues().GetInt("user_id")

		user := model.NewUser(userId, "Reindert Vetter")
		return user, err
	},
}
