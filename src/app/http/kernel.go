package http

import (
	"lanvard/foundation"
	"lanvard/foundation/http"
	"lanvard/foundation/http/middleware"
	pipelineContract "lanvard/interface/pipeline"
)

func Kernel(app foundation.Application) http.KernelStruct {
	return http.KernelStruct{App: app, Middleware: pipes()}
}

func pipes() []pipelineContract.Pipe {
	return []pipelineContract.Pipe{
		middleware.ValidatePostSizeStruct{},
	}
}
