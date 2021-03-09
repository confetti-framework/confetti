package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
	"src/resources/views"
)

// Homepage contains the code responsible for generating the home page.
func Homepage(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Confetti", "Let's be creative!"))
}
