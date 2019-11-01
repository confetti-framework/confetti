package bootstrap

import (
	"laravelgo/foundation"
	interfaceConsole "laravelgo/interface/console"
	interfaceExceptions "laravelgo/interface/exception"
	interfaceHttp "laravelgo/interface/http"
	"laravelgo/src/app/console"
	"laravelgo/src/app/exception"
	"laravelgo/src/app/http"
)

func App() foundation.ApplicationStruct {

	/*
	|--------------------------------------------------------------------------
	| Create The ApplicationStruct
	|--------------------------------------------------------------------------
	|
	| The first thing we will do is create a new Laravel application instance
	| which serves as the "glue" for all the components of Laravel, and is
	| the IoC container for the system binding all of the various parts.
	|
	*/

	app := foundation.Application()

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
