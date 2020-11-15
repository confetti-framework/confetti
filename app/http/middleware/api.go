package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
)

var Api = []inter.HttpMiddleware{
	middleware.DefaultResponseOutcome{Outcome: outcome.Json},
}
