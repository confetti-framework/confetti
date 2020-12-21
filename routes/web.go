package routes

import (
	. "github.com/confetti-framework/routing"
	"confetti-framework/app/http/controllers"
	"confetti-framework/app/http/middleware"
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
	Get("user", controllers.UserCreate),
	Post("user", controllers.UserStore),
	Get("/", controllers.Homepage),
).Middleware(middleware.Web...)
