package views

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

func BlogCreate(app inter.App) *PostCreateView {
	return &PostCreateView{
		Locale: Locale(app),
	}
}

type PostCreateView struct {
	Locale      string
	Title       string
	Description string
}

func (h PostCreateView) Template() string {
	return config.Path.Views + "/blog_create.gohtml"
}
