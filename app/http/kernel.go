package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
)

func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{App: &app}
}
