package views

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
)

type errorView struct {
	inter.View
	Message    string
	StackTrace string
	Status     int
	AppName    string
	Locale     string
	app        inter.App
	template   string
}

func Error(app inter.App, err error) inter.View {
	errorMessage := str.UpperFirst(fmt.Sprintf("%v", err))
	status, _ := errors.FindStatus(err)

	return &errorView{
		Message:    errorMessage,
		StackTrace: StackTrace(app, err),
		Status:     status,
		AppName:    AppName(app),
		Locale:     Locale(app),
		app:        app,
		template:   "error.gohtml",
	}
}

func (e errorView) App() inter.App {
	return e.app
}

func (e errorView) Template() string {
	return e.template
}
