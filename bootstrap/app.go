package bootstrap

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"lanvard/app/console"
	"lanvard/app/exception"
	"lanvard/app/http"
	"lanvard/app/http/decorator"
	"lanvard/config"
)

var bootApp inter.App

func init() {
	// @todo find a way to boot app in init. App has to be the same (in each container instance (before boot and after))
}

func NewAppFromBoot() inter.App {

	// START BOOT
	/*
		|--------------------------------------------------------------------------
		| Create The Application
		|--------------------------------------------------------------------------
		|
		| The first thing we will do is create a new Lanvard application instance
		| which serves as the "glue" for all the components of Lanvard, and is
		| the IoC container for the system binding all of the various parts.
		|
	*/

	bootApp = &foundation.Application{}
	bootApp.SetContainer(foundation.NewContainer())

	bootApp.BindPathsInContainer(config.App.BasePath)
	bootApp.Instance("env", config.App.Env)
	// END BOOT


	bootApp = decorator.Bootstrap(bootApp)
	app := bootApp

	// @todo when we boot in init, we want to copy booted app
	// container := *bootApp.Container()
	// newContainer := container.Copy()
	// app := foundation.Application{}
	// app.SetContainer(newContainer)

	app.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(app),
	)

	app.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(app),
	)

	app.Singleton(
		(*inter.ExceptionHandler)(nil),
		exception.NewHandler(app),
	)

	return app
}
