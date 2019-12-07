package test

import (
	"github.com/stretchr/testify/assert"
	contractsHttp "lanvard/contract/http"
	"lanvard/foundation"
	httpFoundation "lanvard/foundation/http"
	"testing"
)

type testInterface interface{}

func Test_one_binding(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_multiple_binding(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	app.Singleton(
		(*testInterface)(nil),
		httpFoundation.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 2)
}

func Test_binding_two_the_same_interfaces(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_container_make(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*contractsHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	kernel := container.Make((*contractsHttp.Kernel)(nil))

	assert.Equal(t, httpFoundation.Kernel{}, kernel)
}
