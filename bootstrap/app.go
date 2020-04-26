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

	bootApp = decorator.Bootstrap(bootApp)
}

func NewAppFromBoot() *foundation.Application {

	container := *bootApp.Container()
	newContainer := container.Copy()
	app := foundation.Application{}

	app.SetContainer(newContainer)

	app.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(&app),
	)

	app.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(&app),
	)

	app.Singleton(
		(*inter.ExceptionHandler)(nil),
		exception.NewHandler(&app),
	)

	return &app
}
