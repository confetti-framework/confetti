package http

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/foundation/http"
)

// NewKernel generates an HTTP Kernel. The kernel is responsible
// for building the middlewares and calling the controller.
func NewKernel(app inter.App) http.Kernel {
	return http.Kernel{App: &app}
}
