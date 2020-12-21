package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/routing/outcome"
	"confetti-framework/resources/views"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Confetti", "Let's be creative!"))
}
