package views

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"lanvard/config"
)

type errorView struct {
	Message    string
	StackTrace string
	Status     int
	AppName    string
	Locale     string
	template   string
}

func Error(app inter.App, err error) inter.View {
	status, _ := errors.FindStatus(err)

	return &errorView{
		Message:    str.UpperFirst(fmt.Sprintf("%v", err)),
		StackTrace: StackTrace(app, err),
		Status:     status,
		AppName:    AppName(app),
		Locale:     Locale(app),
		template:   config.Path.Views + "/error.gohtml",
	}
}

func (e errorView) Template() string {
	return e.template
}
