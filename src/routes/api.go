package routes

import (
	"lanvard/foundation/http/controllers"
	"lanvard/routing/router"
)

/*
|--------------------------------------------------------------------------
| API MapMethodRoutes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/
var Api = []router.RouteCollection{
	router.Get("/users/{ids?}", controllers.User.Index),
	router.Post("/users", controllers.User.Store),
	router.Delete("/users/{ids?}", controllers.User.Destroy),
}
