package console

import "lanvard/foundation"

type KernelStruct struct {
	App foundation.Application
}

func Kernel(app foundation.Application) KernelStruct {
	return KernelStruct{app}
}
