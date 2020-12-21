package views

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support/str"
	"confetti-framework/config"
	"strings"
)

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

type ErrorView struct {
	Message     string
	StackTraces []string
	Status      int
	AppName     string
	Locale      string
	template    string
}

func (e ErrorView) Template() string {
	return e.template
}
