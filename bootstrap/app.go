package bootstrap

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"lanvard/app/console"
	"lanvard/app/exception"
	"lanvard/app/http"
	"lanvard/config"
)

var bootApp foundation.Application

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

	bootApp = foundation.Application{}
	bootApp.SetContainer(foundation.NewContainer())

	bootApp.BindPathsInContainer(config.App.BasePath)
	(*bootApp.Container()).Instance("env", config.App.Env)

	bootApp = *http.NewKernel(&bootApp).Bootstrap()
}

func NewAppFromBootApp() *foundation.Application {

	newContainer := bootApp.container.Copy()
	app := foundation.Application{
		container: newContainer,
	}

	app.container.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(&app),
	)

	app.container.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(&app),
	)

	app.container.Singleton(
		(*inter.ExceptionHandler)(nil),
		exception.NewHandler(&app),
	)

	return &app
}
