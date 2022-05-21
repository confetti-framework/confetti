package middleware

import (
	"github.com/confetti-framework/framework/foundation/http/middleware"
	"github.com/confetti-framework/framework/foundation/http/outcome"
	"github.com/confetti-framework/framework/inter"
)

// Web contains all middlewares that only apply to the web pages.
var Web = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Html},
}
