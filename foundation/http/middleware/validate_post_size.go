package middleware

import (
	"lanvard/foundation"
	"lanvard/src/app/http/middleware"
)

type ValidatePostSize struct {
	App foundation.Application
}

func (v ValidatePostSize) Handle(data middleware.Passable, next middleware.Destination) middleware.Result {
	// todo validate
	return next(data)
}
