package views

import (
	"github.com/confetti-framework/contract/inter"
	"confetti-framework/config"
)

func Homepage(app inter.App, title string, description string) *HomepageView {
	return &HomepageView{
		Title:       title,
		Description: description,
		Locale:      Locale(app),
	}
}

type HomepageView struct {
	Title       string
	Description string
	Locale      string
}

func (h HomepageView) Template() string {
	return config.Path.Views + "/homepage.gohtml"
}
