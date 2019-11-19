package middleware

import (
	"lanvard/foundation"
	"lanvard/src/app/http/middleware"
)

type ValidatePostSizeStruct struct {
	App foundation.Application
}

func (v ValidatePostSizeStruct) Handle(data middleware.Passable, next middleware.Destination) middleware.Result {
	return next(data)
}
