package providers

import (
	"github.com/lanvard/contract/decorator"
	"lanvard/src/app/providers"
)

var Providers = struct {
	RegisterProviders []decorator.RegisterServiceProvider
	BootProviders     []decorator.BootServiceProvider
}{
	/*
	   |--------------------------------------------------------------------------
	   | Autoloaded Register Service Providers
	   |--------------------------------------------------------------------------
	   |
	   | The service providers listed here will be automatically loaded on the
	   | request to your application. Within the register method, you should only
	   | bind things into the service container. You should never attempt to
	   | register any event listeners, routes, or any other piece of functionality
	   | within the register method. Otherwise, you may accidentally use a service
	   | that is provided by a service provider which has not loaded yet.
	   |
	*/
	RegisterProviders: []decorator.RegisterServiceProvider{},

	/*
	   |--------------------------------------------------------------------------
	   | Autoloaded Boot Service Providers
	   |--------------------------------------------------------------------------
	   |
	   | This method is called after all other service providers have been
	   | registered, meaning you have access to all other services that have been
	   | registered by the framework. Feel free to add your own services to this
	   | slice to grant expanded functionality to your applications.
	   |
	   | You can have a service provider with a register and a boot method. Than
	   | you have to add this service to RegisterProviders and BootProviders
	   |
	*/
	BootProviders: []decorator.BootServiceProvider{
		providers.RouteServiceProvider{},
	},
}
