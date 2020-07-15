package adapter

import (
	"github.com/lanvard/contract/inter"
	"lanvard/app/model"
	"lanvard/src/contract"
)

type User struct {
	Request inter.Request
}

func (adapter User) All() []contract.User {
	users, err := adapter.AllE()
	if err != nil {
		panic(err)
	}

	return users
}

func (adapter User) AllE() ([]contract.User, error) {
	var err error
	var users []contract.User

	users = append(users, model.NewUser(5435, "test@lanvard.com"))

	return users, err
}

func (adapter User) Find() contract.User {
	user, err := adapter.FindE()
	if err != nil {
		panic(err)
	}

	return user
}

func (adapter User) FindE() (contract.User, error) {
	userId, err := adapter.Request.Parameter("user").NumberE()
	if err != nil {
		return nil, err
	}

	return model.NewUser(userId, "test@lanvard.com"), nil
}
