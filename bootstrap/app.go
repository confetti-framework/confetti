package bootstrap

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"src/app/console"
	"src/app/http"
	"src/app/http/decorator"
)

var bootContainer inter.Container

func init() {

	/*
		|--------------------------------------------------------------------------
		| Create boot container
		|--------------------------------------------------------------------------
		|
		| The first thing we will do is create a new boot container instance
		| which is the IoC container for the system binding all binding that is
		| equal for each request. This container is readonly after init.
		|
	*/

	bootContainer = decorator.Bootstrap(foundation.NewContainer())
}

// NewAppFromBoot We create a new inter.App for every new process. We use the
// bootContainer where one-time instances are registered and booted.
func NewAppFromBoot() inter.App {

	/*
		|--------------------------------------------------------------------------
		| Create The Application
		|--------------------------------------------------------------------------
		|
		| The second thing we will do is create a new Confetti application instance
		| which serves as the "glue" for all the components of Confetti, and is
		| the IoC container for the request binding all of the various parts.
		|
	*/

	container := foundation.NewContainerByBoot(bootContainer)
	app := foundation.NewApp()
	app.SetContainer(container)

	app.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(app),
	)

	app.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(app),
	)

	// Bind this function so that a new application can be created later. This is
	// necessary because an application must be made for the command. Later
	// applications must be made for each request.
	app.Bind(inter.AppProvider, NewAppFromBoot)

	return app
}
