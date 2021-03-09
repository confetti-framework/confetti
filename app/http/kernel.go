package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
)

// NewKernel generates a HTTP Kernel. The kernel is responsible
// for building the middlewares and calling the controller.
func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{App: &app}
}
