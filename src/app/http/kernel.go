package http

import (
	"laravelgo/foundation"
	"laravelgo/foundation/http"
)

func Kernel(app foundation.ApplicationStruct) http.KernelStruct {
	return http.KernelStruct{app}
}
