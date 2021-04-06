package views

import (
	_ "embed"
	"github.com/confetti-framework/contract/inter"
)

//go:embed homepage.gohtml
var homepageTemplate string

// Homepage provide the homepage view
func Homepage(app inter.App, title string, description string) *HomepageView {
	return &HomepageView{
		Title:       title,
		Description: description,
		Locale:      Locale(app),
	}
}

// HomepageView contains the view options
type HomepageView struct {
	Title       string
	Description string
	Locale      string
}

// Template returns the content of the template
func (h HomepageView) Template() string {
	return homepageTemplate
}
