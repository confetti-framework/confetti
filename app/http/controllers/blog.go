package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"lanvard/resources/views"
)

func PostCreate(request inter.Request) inter.Response {
	failures := val.Validate(request.Content(),
		val.Verify("title", rule.Required{}, rule.Max{Max: 255}),
		val.Verify("body", rule.Required{}),
	)
	if len(failures) > 0 {
		return outcome.Html("failures")
	}

	return outcome.Html(views.PostCreate(request.App()))
}

func PostStore(request inter.Request) inter.Response {
	failures := val.Validate(request.Content(),
		val.Verify("title", rule.Required{}, rule.Max{Max: 255}),
		val.Verify("body", rule.Required{}),
	)
	if len(failures) > 0 {
		return outcome.Html("failures")
	}

	return outcome.Html(views.PostCreate(request.App()))
}
