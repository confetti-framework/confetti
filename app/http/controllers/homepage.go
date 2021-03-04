package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
	"src/resources/views"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Confetti", "Let's be creative!"))
}
