package views

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

func UserCreate(app inter.App, name, email string) *UserCreateView {
	return &UserCreateView{
		Locale: Locale(app),
		Name:   name,
		Email:  email,
	}
}

type UserCreateView struct {
	Locale string
	Name   string
	Email  string
}

func (h UserCreateView) Template() string {
	return config.Path.Views + "/user_create.gohtml"
}
