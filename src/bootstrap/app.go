package bootstrap

import (
	consoleInterface "github.com/lanvard/contract/console"
	exceptionInterface "github.com/lanvard/contract/exception"
	interfaceHttp "github.com/lanvard/contract/http"
	"github.com/lanvard/foundation"
	"lanvard/src/app/console"
	"lanvard/src/app/exception"
	"lanvard/src/app/http"
)

var bootApp foundation.Application

func init() {
	/*
		|--------------------------------------------------------------------------
		| Create The Application
		|--------------------------------------------------------------------------
		|
		| The first thing we will do is create a new Laravel application instance
		| which serves as the "glue" for all the components of Laravel, and is
		| the IoC container for the system binding all of the various parts.
		|
	*/

	bootApp = foundation.Application{Container: foundation.NewContainer()}

	bootApp.BindPathsInContainer()

	bootApp = http.NewKernel(bootApp).Bootstrap()
}

func NewApp() foundation.Application {

	app := foundation.Application{
		Container: bootApp.Container.Copy(),
	}

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		http.NewKernel(app),
	)

	app.Container.Singleton(
		(*consoleInterface.Kernel)(nil),
		console.NewKernel(app),
	)

	app.Container.Singleton(
		(*exceptionInterface.Handler)(nil),
		exception.NewHandler(app),
	)

	return app
}
