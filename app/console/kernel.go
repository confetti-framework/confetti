package console

import (
	"github.com/lanvard/contract/inter"
)

type Kernel struct {
	App inter.App
}

func NewKernel(app inter.App) Kernel {
	return Kernel{app}
}
