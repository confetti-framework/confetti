package views

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

func PostCreate(app inter.App) *PostCreateView {
	return &PostCreateView{
		Locale: Locale(app),
	}
}

type PostCreateView struct {
	Locale string
}

func (h PostCreateView) Template() string {
	return config.Path.Views + "/post_create.gohtml"
}
