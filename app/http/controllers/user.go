package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/routing/outcome"
	"github.com/confetti-framework/validation/rule"
	"github.com/confetti-framework/validation/val"
	"confetti-framework/resources/views"
)

func UserCreate(request inter.Request) inter.Response {
	return outcome.Html(views.UserCreate(request.App(), nil, "", ""))
}

func UserStore(request inter.Request) inter.Response {
	content := request.Content()
	failures := val.Validate(nil, content,
		val.Verify("name", rule.Required{} /*, rule.StringAble{}, rule.MaxCharacters{Max: 255}*/),
		val.Verify("email", rule.Required{} /*, rule.StringAble{}*/),
	)

	return outcome.Html(views.UserCreate(request.App(),
		failures,
		content.Get("name").String(),
		content.Get("email").String(),
	))
}
