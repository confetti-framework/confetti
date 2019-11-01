package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/foundation"
	httpFoundation "lanvard/foundation/http"
	contractsHttp "lanvard/interface/http"
	"testing"
)

type testInterface interface {}

func Test_one_binding(t *testing.T) {
	app := foundation.Container()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_multiple_binding(t *testing.T) {
	app := foundation.Container()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	app.Singleton(
		(*testInterface)(nil),
		httpFoundation.KernelStruct{},
	)

	assert.Len(t, app.GetBindings(), 2)
}

func Test_binding_two_the_same_interfaces(t *testing.T) {
	app := foundation.Container()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_container_make(t *testing.T) {
	container := foundation.Container()

	container.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.KernelStruct{},
	)

	kernel := container.Make((*contractsHttp.Kernel)(nil))

	assert.Equal(t, httpFoundation.KernelStruct{}, kernel)
}