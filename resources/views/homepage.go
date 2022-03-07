package views

import (
	_ "embed"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
)

//go:embed homepage.gohtml
var homepageTemplate string

// Homepage provide the homepage view
func Homepage(app inter.App, title string, description string) *HomepageView {
	return &HomepageView{
		Title:       title,
		Description: description,
		Locale:      Locale(app),
		ApiUrl:      outcome.UrlByName(app, "ping"),
	}
}

// HomepageView contains the view options
type HomepageView struct {
	Title       string
	Description string
	Locale      string
	ApiUrl      string
}

// Template returns the content of the template
func (h HomepageView) Template() string {
	return homepageTemplate
}
