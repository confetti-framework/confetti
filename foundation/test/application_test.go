package test

import (
	"github.com/stretchr/testify/assert"
	"laravelgo/foundation"
	httpFoundation "laravelgo/foundation/http"
	interfaceHttp "laravelgo/interface/http"
	"laravelgo/src/app/http"
	"testing"
)

func Test_binding(t *testing.T) {
	app := foundation.Application()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		http.Kernel(app),
	)
	assert.Len(t, app.Container.GetBindings(), 1)
}

func Test_application_make(t *testing.T) {
	app := foundation.Application()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	kernel := app.Container.Make((*interfaceHttp.Kernel)(nil))

	assert.Equal(t, httpFoundation.KernelStruct{}, kernel)
}