package http

import (
	"lanvard/foundation"
	"lanvard/foundation/http"
)

func Kernel(app foundation.Application) http.KernelStruct {
	return http.KernelStruct{app}
}
