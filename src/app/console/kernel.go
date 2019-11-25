package console

import "lanvard/foundation"

type Kernel struct {
	App foundation.Application
}

func NewKernel(app foundation.Application) Kernel {
	return Kernel{app}
}
