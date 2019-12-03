package http

import (
	"lanvard/foundation"
	"lanvard/foundation/http"
	foundationMiddleware "lanvard/foundation/http/middleware"
	"lanvard/src/app/http/middleware"
)

func NewKernel(app foundation.Application) http.Kernel {
	return http.Kernel{
		App:        app,
		Middleware: pipes(),
	}
}

func pipes() []middleware.PipeInterface {
	return []middleware.PipeInterface{
		// todo push app
		foundationMiddleware.ValidatePostSize{},
	}
}
