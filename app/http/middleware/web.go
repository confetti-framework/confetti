package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
)

var Web = []inter.HttpMiddleware{
	middleware.ResponseEncoder{Encoder: outcome.Html},
}
