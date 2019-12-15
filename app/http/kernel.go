package http

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	foundationMiddleware "github.com/lanvard/foundation/http/middleware"
	middleware2 "lanvard/app/http/middleware"
)

func NewKernel(app foundation.Application) http.Kernel {
	return http.Kernel{
		App:        app,
		Middleware: pipes(),
	}
}

func pipes() []middleware2.PipeInterface {
	return []middleware2.PipeInterface{
		// todo push app
		foundationMiddleware.ValidatePostSize{},
	}
}
