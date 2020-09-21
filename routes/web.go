package routes

import (
	. "github.com/lanvard/routing"
	"lanvard/app/http/middleware"
	"lanvard/src/controller"
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
	Get("/", controller.Homepage),
	Post("/", controller.Homepage),
).Middleware(middleware.Web...)
