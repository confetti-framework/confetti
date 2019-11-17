package http

import (
	"lanvard/foundation"
	"lanvard/foundation/http"
	"lanvard/foundation/http/middleware"
	"lanvard/src/app/http/pipeline"
)

func Kernel(app foundation.Application) http.KernelStruct {
	return http.KernelStruct{App: app, Middleware: pipes()}
}

func pipes() []pipeline.PipeInterface {
	return []pipeline.PipeInterface{
		middleware.ValidatePostSizeStruct{},
	}
}
