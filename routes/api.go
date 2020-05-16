package routes

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/outcome"
	. "github.com/lanvard/routing"
	"lanvard/src/controllers"
)

/*
|--------------------------------------------------------------------------
| API MapMethodRoutes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. By
| default this is loaded in a group and is assigned to the "api" middleware
| group. By default, API group is placed with "/api" prefix. Feel free to
| remove the prefix if you are using a subdomain for your api. Enjoy
| building your API!
|
*/
var Api = Group(
	Get("/users", controllers.User.Index),
	Get("/users/{users?}", controllers.User.Index),
	Get("/user/{user}", controllers.User.Show),
	Post("/users", controllers.User.Store),
	Delete("/users/{users}", controllers.User.Destroy),
).Prefix("/api")

