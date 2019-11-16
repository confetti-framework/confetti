package middleware

import "lanvard/foundation"

type ValidatePostSizeStruct struct {
	App foundation.Application
}

func (v ValidatePostSizeStruct) Handle(data interface{}, next func(data interface{}) interface{}) interface{} {
	return next(data)
}
