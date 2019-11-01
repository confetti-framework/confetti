package test

import (
	"github.com/stretchr/testify/assert"
	httpFoundation "lanvard/foundation/http"
	interfaceHttp "lanvard/interface/http"
	"lanvard/src/app/http"
	"lanvard/src/bootstrap"
	"testing"
)

func Test_binding(t *testing.T) {
	app := bootstrap.App()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		http.Kernel(app),
	)
	assert.Len(t, app.Container.GetBindings(), 1)
}

func Test_application_make(t *testing.T) {
	app := bootstrap.App()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	kernel := app.Container.Make((*interfaceHttp.Kernel)(nil))

	assert.Equal(t, httpFoundation.KernelStruct{}, kernel)
}
