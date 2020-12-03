package views

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

func UserCreate(app inter.App, failures []error, name, email string) *UserCreateView {
	return &UserCreateView{
		Locale:   Locale(app),
		Failures: failures,
		Name:     name,
		Email:    email,
	}
}

type UserCreateView struct {
	Locale   string
	Name     string
	Email    string
	Failures []error
}

func (h UserCreateView) Template() string {
	return config.Path.Views + "/user_create.gohtml"
}
