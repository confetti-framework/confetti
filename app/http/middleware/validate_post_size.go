package middleware

import "github.com/lanvard/contract/inter"

type ValidatePostSize struct{}

func (v ValidatePostSize) Handle(request inter.Request, next inter.MiddlewareDestination) inter.Response {
	// todo validate
	return next(request)
}
