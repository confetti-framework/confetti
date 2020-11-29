package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"lanvard/resources/views"
)

func PostCreate(request inter.Request) inter.Response {
	return outcome.Html(views.PostCreate(request.App()))
}

func PostStore(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Lanvard", "Let's be creative!"))
}
