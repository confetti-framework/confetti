package bootstrap

import (
	"fmt"
	"github.com/jinzhu/copier"
	"lanvard/foundation"
	interfaceApp "lanvard/interface/application"
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

	bootApp = foundation.Application{Container: foundation.Container()}

	bootApp.SetBasePath()

	bootApp.Container.Instance((*interfaceApp.Container)(nil), bootApp)

	/*
		|--------------------------------------------------------------------------
		| Bind Important Interfaces
		|--------------------------------------------------------------------------
		|
		| Next, we need to bind some important interfaces into the container so
		| we will be able to resolve them when needed. The kernels serve the
		| incoming requests to this application from both the web and CLI.
		|
	*/

	bootApp.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		http.Kernel(bootApp).Bootstrap(),
	)

	bootApp.Container.Singleton(
		(*consoleInterface.Kernel)(nil),
		console.Kernel(bootApp),
	)

	bootApp.Container.Singleton(
		(*exceptionInterface.Handler)(nil),
		exception.Handler(bootApp),
	)
}

func App() foundation.Application {
	var app foundation.Application

	err := copier.Copy(&app, &bootApp)
	if err != nil {
		fmt.Println(err)
		panic("Can't copy application")
	}

	return app
}
