package routes

import (
	. "github.com/lanvard/routing"
	"lanvard/app/http/middleware"
	"lanvard/src/controller"
)

/*
|---------------------------------------------------------------------------
| API routes
|---------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. By default
| this is loaded in a group. The group is assigned to the "Api" middleware
| and is placed with "/api" prefix. Feel free to remove the prefix if you
| are using a subdomain for your API (which is recommended).
| Enjoy building your API!
|
*/

var Api = Group(
	Get("/users", controller.User.Index),
	Get("/user/{user}", controller.User.Show),
	Post("/users", controller.User.Store),
	Delete("/users/{users}", controller.User.Destroy),
).Prefix("/api").Middleware(middleware.Api...)
