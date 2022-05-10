package middleware

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/foundation/http/middleware"
	"github.com/confetti-framework/framework/foundation/http/outcome"
)

// Web contains all middlewares that only apply to the web pages.
var Web = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Html},
}
