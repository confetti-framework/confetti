package middleware

import (
	"github.com/confetti-framework/framework/foundation/http/middleware"
	"github.com/confetti-framework/framework/foundation/http/outcome"
	"github.com/confetti-framework/framework/inter"
)

// Api contains all middlewares that only apply to the API endpoints.
var Api = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Json},
}
