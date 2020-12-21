package views

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"golang.org/x/text/language"
	"strings"
)

func AppName(app inter.App) string {
	name, err := app.MakeE("config.App.Name")
	if err != nil {
		return ""
	}
	return name.(string)
}

func Locale(app inter.App) string {
	lang, err := app.MakeE("config.App.Locale")
	if err != nil {
		return ""
	}
	return lang.(language.Tag).String()
}

func StackTrace(app inter.App, err error) string {
	if !app.Make("config.App.Debug").(bool) {
		return ""
	}
	stackTrace, _ := errors.FindStack(err)
	result := fmt.Sprintf("%+v", stackTrace)
	result = strings.TrimLeft(result, "\n")
	return result
}
