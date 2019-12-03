package bootstrap

import (
	"github.com/jinzhu/copier"
	"lanvard/foundation"
	consoleInterface "lanvard/interface/console"
	exceptionInterface "lanvard/interface/exception"
	interfaceHttp "lanvard/interface/http"
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

	bootApp.SetBasePath()

	bootApp = http.NewKernel(bootApp).Bootstrap()
}

func NewApp() foundation.Application {
	var app foundation.Application

	// Copy booted app
	if copier.Copy(&app, &bootApp) != nil {
		panic("Can't copy application")
	}

	app.Container = foundation.CopyContainer(bootApp.Container)

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
