package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
)

type ValidatePostSize struct {
	App foundation.Application
}

func (v ValidatePostSize) Handle(request inter.Request, next inter.MiddlewareDestination) inter.Response {
	// todo validate
	return next(request)
}
