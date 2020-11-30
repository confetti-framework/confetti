package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"lanvard/resources/views"
)

func UserCreate(request inter.Request) inter.Response {
	return outcome.Html(views.UserCreate(request.App(), "", ""))
}

func UserStore(request inter.Request) inter.Response {
	content := request.Content()
	failures := val.Validate(content,
		val.Verify("name", rule.Required{}, rule.StringAble{}, rule.Max{Max: 255}),
		val.Verify("email", rule.Required{}, rule.StringAble{}),
	)
	if len(failures) > 0 {
		return outcome.Html(failures)
	}

	return outcome.Html(views.UserCreate(request.App(),
		content.Get("name").String(),
		content.Get("email").String(),
	))
}
