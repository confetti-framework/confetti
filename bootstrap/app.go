package bootstrap

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"lanvard/app/console"
	"lanvard/app/exception"
	"lanvard/app/http"
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

	bootApp = foundation.Application{Container: foundation.NewContainer()}

	bootApp.BindPathsInContainer()

	bootApp = *http.NewKernel(&bootApp).Bootstrap()
}

func NewApp() *foundation.Application {

	newContainer := bootApp.Container.Copy()
	app := foundation.Application{
		Container: newContainer,
	}

	app.Container.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(&app),
	)

	app.Container.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(&app),
	)

	app.Container.Singleton(
		(*inter.ExceptionHandler)(nil),
		exception.NewHandler(&app),
	)

	return &app
}
