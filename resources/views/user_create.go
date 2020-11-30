package views

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

func UserCreate(app inter.App) *UserCreateView {
	return &UserCreateView{
		Locale: Locale(app),
	}
}

type UserCreateView struct {
	Locale      string
	Title       string
	Description string
}

func (h UserCreateView) Template() string {
	return config.Path.Views + "/user_create.gohtml"
}
