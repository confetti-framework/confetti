package middleware

import (
	"lanvard/foundation"
	"lanvard/src/app/http/pipeline"
)

type ValidatePostSizeStruct struct {
	App foundation.Application
}

func (v ValidatePostSizeStruct) Handle(data pipeline.Passable, next pipeline.Destination) pipeline.Result {
	return next(data)
}
