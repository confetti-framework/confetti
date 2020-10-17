package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	foundation "github.com/lanvard/foundation/http/middleware"
	"lanvard/app/http/middleware"
)

var middlewares = []inter.HttpMiddleware{
	foundation.RequestBodyDecoder{},
	middleware.ValidatePostSize{}, // todo remove or use ValidatePostSize
	middleware.RouteModelBinding{},
}

func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{
		App:        &app,
		Middleware: middlewares,
	}
}
