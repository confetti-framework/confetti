package routes

import (
	"confetti-framework/app/http/controllers"
	"confetti-framework/app/http/middleware"
	. "github.com/confetti-framework/foundation/http/routing"
)

/*
	|--------------------------------------------------------------------------
	| Web routes
	|--------------------------------------------------------------------------
	|
	| Here is where you can register Web routes for your website. By default
	| these routes are loaded within a group which is assigned the "Web"
	| middleware. Enjoy building your website!
	|
*/
var Web = Group(
	Get("/", controllers.Homepage),
).Middleware(middleware.Web...)
