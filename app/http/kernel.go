package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"lanvard/app/http/middleware"
)

func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{
		App:        &app,
		Middleware: pipes(),
	}
}

func pipes() []inter.Pipe {
	return []inter.Pipe{
		// todo remove or use ValidatePostSize
		middleware.ValidatePostSize{},
	}
}
