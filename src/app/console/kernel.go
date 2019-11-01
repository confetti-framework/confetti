package console

import "lanvard/foundation"

type KernelStruct struct {
	app foundation.Application
}

func Kernel(app foundation.Application) KernelStruct {
	return KernelStruct{app}
}
