package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/routing/outcome"
)

var Api = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Json},
}
