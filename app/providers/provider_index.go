package providers

import (
	"github.com/confetti-framework/framework/foundation/providers"
	"github.com/confetti-framework/framework/inter"
	"src/config"
)

// Providers contains all service providers
var Providers = struct {
	RegisterProviders []inter.RegisterServiceProvider
	BootProviders     []inter.BootServiceProvider
}{
	/*
	   |--------------------------------------------------------------------------
	   | Register Service Providers
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
	RegisterProviders: []inter.RegisterServiceProvider{
		EnvServiceProvider{},
		providers.ConfigServiceProvider{Index: config.Index},
		AppServiceProvider{},
		ViewServiceProvider{},
		ResponseServiceProvider{},
	},

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
	   | you have to add this provider to RegisterProviders and BootProviders
	   |
	*/
	BootProviders: []inter.BootServiceProvider{
		AppServiceProvider{},
		RouteServiceProvider{},
		providers.DatabaseServiceProvider{Connections: config.Database.Connections},
	},
}
