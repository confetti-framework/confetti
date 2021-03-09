package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/outcome"
)

// Api contains all middlewares that only apply to the API endpoints.
var Api = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Json},
}
