package views

import (
	_ "embed"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support/str"
	"strings"
)

//go:embed error.gohtml
var errorTemplate string

// Error provide an error view
func Error(app inter.App, err error) inter.View {
	status, _ := errors.FindStatus(err)

	return &ErrorView{
		Message:     str.UpperFirst(fmt.Sprintf("%v", err)),
		StackTraces: strings.Split(StackTrace(app, err), "\n"),
		Status:      status,
		AppName:     AppName(app),
		Locale:      Locale(app),
	}
}

// ErrorView contains the view options
type ErrorView struct {
	Message     string
	StackTraces []string
	Status      int
	AppName     string
	Locale      string
}

// Template returns the template path
func (e ErrorView) Template() string {
	return errorTemplate
}
