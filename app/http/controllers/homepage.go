package controllers

import (
	"confetti-framework/resources/views"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Confetti", "Let's be creative!"))
}
