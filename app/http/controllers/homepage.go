package controllers

import (
	"github.com/confetti-framework/framework/foundation/http/outcome"
	"github.com/confetti-framework/framework/inter"
	"src/resources/views"
)

// Homepage contains the code responsible for generating the home page.
func Homepage(request inter.Request) inter.Response {
	return outcome.Html(views.Homepage(request.App(), "Confetti", "Let's be creative!"))
}
