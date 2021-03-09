package views

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support/str"
	"src/config"
	"strings"
)

// Error provide an error view
func Error(app inter.App, err error) inter.View {
	status, _ := errors.FindStatus(err)

	return &ErrorView{
		Message:     str.UpperFirst(fmt.Sprintf("%v", err)),
		StackTraces: strings.Split(StackTrace(app, err), "\n"),
		Status:      status,
		AppName:     AppName(app),
		Locale:      Locale(app),
		template:    config.Path.Views + "/error.gohtml",
	}
}

// ErrorView contains the view options
type ErrorView struct {
	Message     string
	StackTraces []string
	Status      int
	AppName     string
	Locale      string
	template    string
}

// Template returns the template path
func (e ErrorView) Template() string {
	return e.template
}
