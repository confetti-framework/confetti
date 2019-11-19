package http

import (
	"lanvard/foundation"
	"lanvard/foundation/http"
	foundationMiddleware "lanvard/foundation/http/middleware"
	"lanvard/src/app/http/middleware"
)

func Kernel(app foundation.Application) http.KernelStruct {
	return http.KernelStruct{App: app, Middleware: pipes()}
}

func pipes() []middleware.PipeInterface {
	return []middleware.PipeInterface{
		foundationMiddleware.ValidatePostSizeStruct{},
	}
}
