package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
)

func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{App: &app}
}
