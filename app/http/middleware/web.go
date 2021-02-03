package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/outcome"
)

var Web = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Html},
}
