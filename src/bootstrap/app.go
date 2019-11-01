package bootstrap

import (
	"lanvard/foundation"
	interfaceApp "lanvard/interface/application"
	interfaceConsole "lanvard/interface/console"
	interfaceExceptions "lanvard/interface/exception"
	interfaceHttp "lanvard/interface/http"
	"lanvard/src/app/console"
	"lanvard/src/app/exception"
	"lanvard/src/app/http"
	"path/filepath"
	"runtime"
)

func App() foundation.Application {

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

	app := foundation.Application{Container: foundation.Container()}

	_, filename, _, _ := runtime.Caller(1)
	app.SetBasePath(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(filename)))))

	app.Container.Instance((*interfaceApp.Container)(nil), app)

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

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		http.Kernel(app),
	)

	app.Container.Singleton(
		(*interfaceConsole.Kernel)(nil),
		console.Kernel(app),
	)

	app.Container.Singleton(
		(*interfaceExceptions.Handler)(nil),
		exception.Handler(app),
	)

	return app
}
