package console

import "laravelgo/foundation"

type KernelStruct struct {
	app foundation.ApplicationStruct
}

func Kernel(app foundation.ApplicationStruct) KernelStruct {
	return KernelStruct{app}
}
